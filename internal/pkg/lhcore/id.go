package lhcore

import "strconv"

type IDSet struct {
	ServerID uint8
	StageID  uint32
	ItemID   uint32
}

func (s IDSet) Hash() uint64 {
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

func (p IDPair) Hash() uint32 {
	// give ServerID 4 bits (16 servers supported), EntityID 28 bits (268435456 entities supported) so that ServerID with EntityID can be combined into uint32 without collision
	return uint32(p.ServerID)<<28 | p.EntityID
}
