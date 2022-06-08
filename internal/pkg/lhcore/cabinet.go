package lhcore

import "sync"

type Registers struct {
	Sum uint32

	Subsections *sync.Map
}

type Cabinet struct {
	StageID uint32
	ItemID  uint32

	// Values
	Times    Registers
	Quantity Registers
}

type Rack struct {
	Cabinets sync.Map
}
