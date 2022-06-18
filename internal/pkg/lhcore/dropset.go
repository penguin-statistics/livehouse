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
	StageID uint32
	ItemID  uint32
}

func (s IDSet) ID() uint64 {
	return uint64(s.StageID)<<32 | uint64(s.ItemID)
}

func (s IDSet) GoString() string {
	return "IDSet{StageID: " + strconv.FormatUint(uint64(s.StageID), 10) + ", ItemID: " + strconv.FormatUint(uint64(s.ItemID), 10) + "}"
}

type DropElementValue struct {
	// secs are the subsections in which denotes the quantity of drops for
	// different geenrations
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

func (d *DropElementValue) Incr(delta, generation uint64) {
	currentGen := atomic.LoadUint64(&d.gen)
	if generation < currentGen {
		return
	}
	atomic.AddUint64(&d.sum, delta)
	initval := uint64(0)
	subsection, _ := d.secs.LoadOrStore(generation, &initval)
	atomic.AddUint64(subsection.(*uint64), delta)
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
	e.Times.Incr(times, generation)
	e.Quantity.Incr(quantity, generation)

	liteval := &LiteValue{
		Times:    times,
		Quantity: quantity,
	}
	log.Debug().Interface("liteval", liteval).Msg("Incr")
	e.Subscriptions.Range(func(key, value any) bool {
		log.Debug().Msgf("Sending litevalue to sub %s", value.(*Sub).ClientID)
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
	StageElements   map[uint32][]*DropElement
	ItemElements    map[uint32][]*DropElement
}

func NewDropSet() *DropSet {
	return &DropSet{
		CombineElements: make(map[uint64]*DropElement),
		StageElements:   make(map[uint32][]*DropElement),
		ItemElements:    make(map[uint32][]*DropElement),
	}
}

func (d *DropSet) GetOrCreateElement(idset IDSet) *DropElement {
	d.mu.Lock()
	defer d.mu.Unlock()

	element := &DropElement{
		IDSet: idset,
	}

	actual, ok := d.CombineElements[idset.ID()]
	if !ok {
		d.CombineElements[idset.ID()] = element
		actual = element
	}

	stageEl, ok := d.StageElements[idset.StageID]
	if !ok {
		made := []*DropElement{element}
		d.StageElements[idset.StageID] = made
		stageEl = made
	} else {
		d.StageElements[idset.StageID] = append(stageEl, element)
	}

	itemEl, ok := d.ItemElements[idset.ItemID]
	if !ok {
		made := []*DropElement{element}
		d.ItemElements[idset.ItemID] = made
		itemEl = made
	} else {
		d.ItemElements[idset.ItemID] = append(itemEl, element)
	}

	return actual
}

func (d *DropSet) ReplaceSubToStageElements(stageID uint32, sub *Sub) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.removeSub(sub)

	elements, ok := d.StageElements[stageID]
	if !ok {
		return ErrStageNotFound
	}

	log.Debug().Interface("elements", elements).Msg("ReplaceSubToStageElements")

	for _, element := range elements {
		log.Debug().Interface("element", element).Msg("ReplaceSubToStageElements")
		element.AddSub(sub)
	}

	return nil
}

func (d *DropSet) ReplaceSubToItemElements(itemID uint32, sub *Sub) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.removeSub(sub)

	elements, ok := d.ItemElements[itemID]
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
