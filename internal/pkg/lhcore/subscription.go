package lhcore

import (
	"sync"
	"sync/atomic"
)

type Element struct {
	StageID uint32
	ItemID  uint32

	LiteValue
}

type LiteValue struct {
	Quantity uint64
	Times    uint64
}

type Sub struct {
	ClientID string

	// buf is an atomic.Value protected sync.Map which in turn
	// stores map[IDSet]LiteValue
	buf atomic.Value
}

func NewSub(clientId string) *Sub {
	s := &Sub{
		ClientID: clientId,
		buf:      atomic.Value{},
	}
	s.buf.Store(&sync.Map{})
	return s
}

func (s *Sub) Get(id IDSet) *LiteValue {
	buf := s.buf.Load().(*sync.Map)
	if v, ok := buf.Load(id); ok {
		return v.(*LiteValue)
	}
	return nil
}

func (s *Sub) Set(id IDSet, value *LiteValue) {
	buf := s.buf.Load().(*sync.Map)
	buf.Store(id, value)
}

func (s *Sub) Swap() *sync.Map {
	return s.buf.Swap(&sync.Map{}).(*sync.Map)
}

func (s *Sub) Flush() []Element {
	elements := make([]Element, 0)
	buf := s.Swap()
	buf.Range(func(key, value any) bool {
		elements = append(elements, Element{
			StageID: key.(IDSet).StageID,
			ItemID:  key.(IDSet).ItemID,
			LiteValue: LiteValue{
				Quantity: value.(*LiteValue).Quantity,
				Times:    value.(*LiteValue).Times,
			},
		})
		return true
	})
	return elements
}
