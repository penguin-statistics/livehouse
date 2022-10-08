package lhcore

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var (
	ErrStageNotFound = errors.New("stage not found")
	ErrItemNotFound  = errors.New("item not found")
)

// DropSet consists a completed collection of all `DropElement`s existing on such instance.
// It stores a singular `DropElement` into three different maps, to improve index efficiencies under
// different workloads. See docuemntation on those particular maps for their intended usage.
type DropSet struct {
	mu sync.Mutex

	// CombineElements is used to get a singular element (if exists)
	// by indexing its ServerID, StageID, and ItemID.
	// Such indexing uint64 is calculated using IDSet#ID().
	CombineElements map[uint64]*DropElement

	// StageElements is used to get a list of elements for a particular (ServerID, StageID) pair (if exists)
	// It is actually map[EntityID]map[EntityID]*DropElement, in which
	// the first EntityID consists of a (ServerID, StageID) pair, and
	// the second EntityID consists of a (ServerID, ItemID) pair.
	StageElements map[uint32]map[uint32]*DropElement

	// ItemElements is used to get a list of elements for a particular (ServerID, ItemID) pair (if exists)
	// It is actually map[EntityID]map[EntityID]*DropElement, in which
	// the first EntityID consists of a (ServerID, ItemID) pair, and
	// the second EntityID consists of a (ServerID, StageID) pair.
	ItemElements map[uint32]map[uint32]*DropElement
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

	actual, ok := d.CombineElements[idset.Hash()]
	if ok {
		return actual
	}

	element := &DropElement{
		IDSet: idset,
	}
	d.CombineElements[idset.Hash()] = element
	actual = element

	spairId := idset.StagePair().Hash()
	ipairId := idset.ItemPair().Hash()

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
	elements, ok := d.StageElements[pair.Hash()]
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
	elements, ok := d.ItemElements[pair.Hash()]
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
	pairid := IDPair{ServerID: server, EntityID: stageId}.Hash()
	m := d.StageElements[pairid]
	for _, element := range m {
		element.Incr(1, 0, generation)
	}
}
