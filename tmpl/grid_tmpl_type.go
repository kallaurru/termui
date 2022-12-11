package tmpl

import . "github.com/kallaur/termui/v3"

const (
	CellTmplAdded     uint8 = 0x00
	CellTmplBadRowIdx uint8 = 0x01
	CellTmplIsExists  uint8 = 0x02
	CellTmplOtherErr  uint8 = 0x04
)

type GridTmpl struct {
	rows  uint8   // количество строк в шаблоне
	sizes []uint8 // высоты строк
	ccr   uint8   // количество столбцов в строке по умолчанию
	cells map[uint16]*CellDetail
}

func NewGridTmpl(rows uint8, sizes ...uint8) *GridTmpl {
	const max uint8 = 100
	var (
		counter       = 0
		defSize uint8 = 0
		allS    uint8 = 0
		s             = make([]uint8, 0, int(rows))
	)

	l := len(sizes)
	defSize = max / rows

	for counter = 0; counter < int(rows); counter++ {
		if l == 0 {
			allS += defSize
			s[counter] = defSize
			continue
		}
		if l-1 > counter {
			size := sizes[counter]
			allS += size
			s[counter] = size
			continue
		}
		sz := max - allS

	}
	return &GridTmpl{
		rows:  rows,
		ccr:   1,
		cells: make(map[uint16]*CellDetail),
		sizes: make([]uint8, 0, int(rows)),
	}
}

// AddCell - индексация строк и столбцов начинается с 0
func (gt *GridTmpl) AddCell(row, col, size uint8, widget Drawable) uint8 {
	if row > gt.rows {
		return CellTmplBadRowIdx
	}
	cd := NewCellDetail(row, col, size, widget)
	idx := cd.idx

	_, ok := gt.cells[idx]
	if ok {
		return CellTmplIsExists
	}
	gt.cells[idx] = cd

	return CellTmplAdded
}
