package tmpl

import (
	"container/list"
	. "github.com/kallaur/termui/v3"
)

/** Весь макет делим на строки, далее в ячейках могут быть вложенные схемы */

type GridSchema struct {
	rowsSizes map[uint8]AdaptiveSize
	cells     map[uint8]*list.List // в списке или схема или детали (GridSchema, CellDetail)
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
	sum := NewAdaptiveSize(0)
	for idx, s := range sizes {
		gs.rowsSizes[uint8(idx)] = s
		sum = sum.Sum(s)
	}
	if sum.IsMax() {
		return gs
	}
	return nil
}

func (gs *GridSchema) AddCell(row uint8, cd *CellDetail) {
	gs.addCell(row, cd)
}

func (gs *GridSchema) AddCellSchema(row uint8, grdS *GridSchema) {
	gs.addCell(row, grdS)
}

// Build - по количеству строк в схеме
func (gs *GridSchema) Build() []GridItem {
	lenCells := len(gs.cells)
	out := make([]GridItem, 0, lenCells)
	for i := 0; i < lenCells; i++ {
		r := uint8(i)
		cells, ok := gs.cells[r]
		if !ok || cells == nil {
			continue
		}
		item := gs.compile(r)
		out = append(out, item)
	}
	return out
}

func (gs *GridSchema) BuildGrid() *Grid {
	items := gs.Build()
	grid := NewGrid()
	grid.Set(items)

	return grid
}

func (gs *GridSchema) Sizes() []float64 {
	length := len(gs.rowsSizes)
	out := make([]float64, 0, length)
	for i := 0; i < length; i++ {
		idx := uint8(i)
		size, ok := gs.rowsSizes[idx]
		if ok {
			out = append(out, size.FloatSize())
		}
	}

	return out
}

func (gs *GridSchema) addCell(row uint8, value interface{}) {
	_, has := gs.rowsSizes[row]
	if !has {
		// такого индекса строки нет
		return
	}
	l, has := gs.cells[row]
	if !has {
		l = list.New()
	}
	l.PushBack(value)
}

func (gs *GridSchema) compile(row uint8) GridItem {
	cells := gs.cells[row]

	for e := cells.Front(); e != nil; e = e.Next() {
	}
}
