package tmpl

import (
	"container/list"
	. "github.com/kallaur/termui/v3"
)

type GridSchema struct {
	rowsSizes map[uint8]AdaptiveSize
	cells     map[uint8]*list.List
}

func NewGridSchema(sizes ...AdaptiveSize) *GridSchema {
	gs := &GridSchema{
		rowsSizes: make(map[uint8]AdaptiveSize),
		cells:     make(map[uint8]*list.List),
	}
	if len(sizes) == 0 {
		gs.rowsSizes[0] = NewAdaptiveSizeMax()
		return gs
	}

	for idx, s := range sizes {
		gs.rowsSizes[uint8(idx)] = s
	}

	return gs
}

func (gs *GridSchema) AddCell() {

}

func (gs *GridSchema) Build() GridItem {

}
