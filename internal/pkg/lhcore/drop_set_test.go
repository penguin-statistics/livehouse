package lhcore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropSet(t *testing.T) {
	ds := NewDropSet()
	idset := IDSet{
		ServerID: 0, // CN
		StageID:  131,
		ItemID:   121,
	}
	el1 := ds.GetOrCreateElement(idset)
	el2 := ds.GetOrCreateElement(idset)

	assert.Equal(t, el1, el2)
}
