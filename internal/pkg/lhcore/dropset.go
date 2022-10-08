package lhcore

import (
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var (
	ErrStageNotFound = errors.New("stage not found")
	ErrItemNotFound  = errors.New("item not found")
)

type IDSet struct {
	ServerID uint8
	StageID  uint32
	ItemID   uint32
}

func (s IDSet) ID() uint64 {
	// give ServerID 4 bits (16 servers supported), StageID 28 bits (268435456 stages supported), ItemID 28 bits (268435456 items supported) so that ServerID with any one of StageID or ItemID can be combined into uint32 without collision. The most significant 4 bits are reserved for future use.
	// putting serverID in the highest bits to make it easier for sorting and more friendly for hashmap
	return uint64(s.ServerID)<<56 | uint64(s.StageID)<<28 | uint64(s.ItemID)
}

func (s IDSet) StagePair() IDPair {
	return IDPair{
		ServerID: s.ServerID,
		EntityID: s.StageID,
	}
}

func (s IDSet) ItemPair() IDPair {
	return IDPair{
		ServerID: s.ServerID,
		EntityID: s.ItemID,
	}
}

func (s IDSet) GoString() string {
	return "IDSet{StageID: " + strconv.FormatUint(uint64(s.StageID), 10) + ", ItemID: " + strconv.FormatUint(uint64(s.ItemID), 10) + ", ServerID: " + strconv.FormatUint(uint64(s.ServerID), 10) + "}"
}

type IDPair struct {
	ServerID uint8
	EntityID uint32
}

func (p IDPair) ID() uint32 {
	// give ServerID 4 bits (16 servers supported), EntityID 28 bits (268435456 entities supported) so that ServerID with EntityID can be combined into uint32 without collision
	return uint32(p.ServerID)<<28 | p.EntityID
}

type DropElementValue struct {
	// secs are the subsections in which denotes the quantity of drops for
	// different generations
	secs sync.Map

	sync.RWMutex
	// sum is the sum of all the subsections; it is here for
	// optimizing the calculation of subsections.
	sum uint64
	// abs is the currently-relied on absolute matrix calculation result
	// received from the last matrix batch.
	abs uint64
	// gen is the current generation of the absolute matrix calculation.
	gen uint64
}

func (d *DropElementValue) Sum() uint64 {
	return atomic.LoadUint64(&d.sum)
}

func (d *DropElementValue) Incr(delta, generation uint64) uint64 {
	currentGen := atomic.LoadUint64(&d.gen)
	if generation < currentGen {
		return atomic.LoadUint64(&d.sum)
	}
	atomic.AddUint64(&d.sum, delta)
	initval := uint64(0)
	subsection, _ := d.secs.LoadOrStore(generation, &initval)
	return atomic.AddUint64(subsection.(*uint64), delta)
}

func (d *DropElementValue) CutOut(value, generation uint64) {
	d.Lock()
	defer d.Unlock()
	d.secs.Range(func(key, value any) bool {
		if key.(uint64) < generation {
			d.secs.Delete(key)
		}
		return true
	})
	d.abs = value
	d.gen = generation
	d.recalcSum()
}

func (d *DropElementValue) recalcSum() {
	d.sum = 0
	d.secs.Range(func(key, value any) bool {
		d.sum += *value.(*uint64)
		return true
	})
	d.sum += d.abs
}

type DropElement struct {
	IDSet

	// Values
	Times    DropElementValue
	Quantity DropElementValue

	// Subscriptions
	Subscriptions sync.Map
}

func (e *DropElement) Incr(times, quantity, generation uint64) {
	log.Trace().
		Interface("idset", e.IDSet).
		Uint64("times", times).
		Uint64("quantity", quantity).
		Msg("DropElement.Incr")
	incrT := e.Times.Incr(times, generation)
	incrQ := e.Quantity.Incr(quantity, generation)

	liteval := &LiteValue{
		Times:    incrT,
		Quantity: incrQ,
	}
	e.Subscriptions.Range(func(key, value any) bool {
		log.Trace().
			Str("clientId", value.(*Sub).ClientID).
			Interface("idset", e.IDSet).
			Interface("liteValue", liteval).
			Msgf("sending litevalue")
		value.(*Sub).Set(e.IDSet, liteval)
		return true
	})
}

func (e *DropElement) CutOut(times, quantity, generation uint64) {
	e.Times.CutOut(times, generation)
	e.Quantity.CutOut(quantity, generation)
}

func (e *DropElement) AddSub(sub *Sub) {
	e.Subscriptions.Store(sub.ClientID, sub)
}

func (e *DropElement) RemoveSub(sub *Sub) {
	e.Subscriptions.Delete(sub.ClientID)
}

type DropSet struct {
	mu sync.Mutex

	CombineElements map[uint64]*DropElement
	StageElements   map[uint32]map[uint32]*DropElement
	ItemElements    map[uint32]map[uint32]*DropElement
}

func NewDropSet() *DropSet {
	return &DropSet{
		CombineElements: make(map[uint64]*DropElement),
		StageElements:   make(map[uint32]map[uint32]*DropElement),
		ItemElements:    make(map[uint32]map[uint32]*DropElement),
	}
}

func (d *DropSet) GetOrCreateElement(idset IDSet) *DropElement {
	d.mu.Lock()
	defer d.mu.Unlock()

	actual, ok := d.CombineElements[idset.ID()]
	if ok {
		return actual
	}

	element := &DropElement{
		IDSet: idset,
	}
	d.CombineElements[idset.ID()] = element
	actual = element

	spairId := idset.StagePair().ID()
	ipairId := idset.ItemPair().ID()

	_, ok = d.StageElements[spairId]
	if !ok {
		made := map[uint32]*DropElement{idset.ItemID: element}
		d.StageElements[spairId] = made
	} else {
		d.StageElements[spairId][idset.ItemID] = element
	}

	_, ok = d.ItemElements[ipairId]
	if !ok {
		made := map[uint32]*DropElement{idset.StageID: element}
		d.ItemElements[ipairId] = made
	} else {
		d.ItemElements[ipairId][idset.StageID] = element
	}

	return actual
}

func (d *DropSet) ReplaceSubToStageElements(stageID uint32, server uint8, sub *Sub) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.removeSub(sub)

	pair := IDPair{ServerID: server, EntityID: stageID}
	elements, ok := d.StageElements[pair.ID()]
	if !ok {
		return ErrStageNotFound
	}

	log.Debug().Interface("elements", elements).Msg("ReplaceSubToStageElements")

	for _, element := range elements {
		element.AddSub(sub)
	}

	return nil
}

func (d *DropSet) ReplaceSubToItemElements(itemID uint32, server uint8, sub *Sub) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.removeSub(sub)

	pair := IDPair{ServerID: server, EntityID: itemID}
	elements, ok := d.ItemElements[pair.ID()]
	if !ok {
		return ErrItemNotFound
	}

	log.Debug().Interface("elements", elements).Msg("ReplaceSubToItemElements")

	for _, element := range elements {
		log.Debug().Interface("element", element).Msg("ReplaceSubToItemElements")
		element.AddSub(sub)
	}

	return nil
}

func (d *DropSet) removeSub(sub *Sub) {
	for _, element := range d.CombineElements {
		element.RemoveSub(sub)
	}
}

func (d *DropSet) RemoveSub(sub *Sub) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.removeSub(sub)
}

func (d *DropSet) IncrTimes(stageId uint32, server uint8, generation uint64) {
	pairid := IDPair{ServerID: server, EntityID: stageId}.ID()
	m := d.StageElements[pairid]
	for _, element := range m {
		element.Incr(1, 0, generation)
	}
}
