package lhcore

import (
	"sync"
	"sync/atomic"

	"github.com/rs/zerolog/log"
)

// DropElementValue
type DropElementValue struct {
	// secs are the subsections in which denotes the quantity of drops for
	// different generations.
	// The key of the map is a uint64 generation value, and the value of the map
	// is a *uint64 actual value corresponding to suhc generation. The pointer
	// is for atomic.AddUint64 to update the value atomically so that DropElementValue
	// could be Incr-ed with low overhead.
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

// DropElement contains all metrics for a drop element identified by
// an IDSet, as well all subscriptions to such DropElement.
type DropElement struct {
	IDSet

	// Times describes the amount of occurances of reports such drop
	// may possibly drop.
	Times DropElementValue
	// Quantity describes the amount of occurances that such item has dropped
	Quantity DropElementValue

	// Subscriptions, map[ClientID]*Sub
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
