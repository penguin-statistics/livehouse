package lhcore

import (
	"sync"
	"sync/atomic"
)

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
	StageID uint32
	ItemID  uint32

	// Values
	Times    DropElementValue
	Quantity DropElementValue
}

func (e *DropElement) ID() uint64 {
	return uint64(e.StageID)<<32 | uint64(e.ItemID)
}

func (e *DropElement) Incr(times, quantity, generation uint64) {
	e.Times.Incr(times, generation)
	e.Quantity.Incr(quantity, generation)
}

func (e *DropElement) CutOut(times, quantity, generation uint64) {
	e.Times.CutOut(times, generation)
	e.Quantity.CutOut(quantity, generation)
}

type DropSet struct {
	CombineElements sync.Map
	StageElements   sync.Map
	ItemElements    sync.Map
}

func NewDropSet() *DropSet {
	return &DropSet{}
}

func (d *DropSet) GetOrCreateElement(stageID, itemID uint32) *DropElement {
	element := &DropElement{
		StageID: stageID,
		ItemID:  itemID,
	}
	actual, _ := d.CombineElements.LoadOrStore(element.ID(), element)
	d.StageElements.Store(stageID, actual)
	d.ItemElements.Store(itemID, actual)

	return actual.(*DropElement)
}
