package tmpl

import . "github.com/kallaur/termui/v3"

const (
	CellTmplAdded     uint8 = 0x00
	CellTmplBadRowIdx uint8 = 0x01
	CellTmplIsExists  uint8 = 0x02
	CellTmplOtherErr  uint8 = 0x04
)

type GridTmpl struct {
	rows  uint8          // количество строк в шаблоне
	sizes []AdaptiveSize // высоты строк
	cells map[uint16]*CellDetail
}

func NewGridTmpl(rCount uint8, rSize ...AdaptiveSize) *GridTmpl {
	var s []AdaptiveSize
	if len(rSize) == 0 {
		s = make([]AdaptiveSize, 1, 1)
		s[0] = NewAdaptiveSizeMax()
	} else {
		s = make([]AdaptiveSize, 0, len(rSize))
		s = append(s, rSize...)
	}

	return &GridTmpl{
		rows:  rCount,
		sizes: s,
		cells: make(map[uint16]*CellDetail),
	}
}

// AddCell - индексация строк и столбцов начинается с 0
func (gt *GridTmpl) AddCell(row, col uint8, size AdaptiveSize, widget Drawable) uint8 {
	detail := NewCellDetail(row, col, size, widget)
	gt.cells[detail.idx] = detail
	return CellTmplAdded
}
