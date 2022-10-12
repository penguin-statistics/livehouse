package lhcore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDSetPairConv(t *testing.T) {
	idset := IDSet{
		ServerID: 1,  // US
		StageID:  18, // main_01-07
		ItemID:   7,  // 30012
	}

	assert.Equal(t, idset.ItemPair(), IDPair{
		ServerID: idset.ServerID,
		EntityID: idset.ItemID,
	})

	assert.Equal(t, idset.StagePair(), IDPair{
		ServerID: idset.ServerID,
		EntityID: idset.StageID,
	})
}

func TestIDSetHash(t *testing.T) {
	idset := IDSet{
		ServerID: 1,  // US
		StageID:  18, // main_01-07
		ItemID:   7,  // 30012
	}

	// 0b00000001_0000000000000000000000010010_0000000000000000000000000111
	//   |   |    |                            |
	//   |Reserved|StageID: dec(18)=bin(10010) | ItemID: dec(7)=bin(111)
	//       |ServerID: US ServerID is dec(1)=bin(1)
	assert.Equal(t, idset.Hash(), uint64(0b00000001_0000000000000000000000010010_0000000000000000000000000111))
}
