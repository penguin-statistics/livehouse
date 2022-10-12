package lhcore

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubMaps(t *testing.T) {
	sub := &Sub{
		ClientID: "abc",
	}

	sub.buf.Store(&sync.Map{})
	m, ok := sub.buf.Load().(*sync.Map)
	assert.True(t, ok, "expect map to be existed once stored in atomic.Value")

	m.Store("abc", sub)
	v, ok := m.Load("abc")
	assert.True(t, ok, "expect value to be existed once stored in map")
	assert.Equal(t, v, sub, "expect pointer stored in map equals to the pointer passed to (*sync.Map)#Store()")

	sub.buf.Store(m)

	// start over
	m, ok = sub.buf.Load().(*sync.Map)
	assert.True(t, ok, "expect map to be existed once stored in atomic.Value")

	v, ok = m.Load("abc")
	assert.True(t, ok, "expect value to be existed once stored in map")
	assert.Equal(t, v, sub, "expect pointer stored in map equals to the pointer passed to (*sync.Map)#Store()")
}

func TestSubMethods(t *testing.T) {
	sub := NewSub("123")
	idset := IDSet{
		ServerID: 0,
		StageID:  18,
		ItemID:   7,
	}
	lv := &LiteValue{}
	sub.Set(idset, lv)
	assert.Equal(t, sub.Get(idset), lv, "expect pointer got from sub.Get() equals to pointer of original value")
}
