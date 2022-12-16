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

	return nil
}

// AddCell - индексация строк и столбцов начинается с 0
func (gt *GridTmpl) AddCell(row, col, size AdaptiveSize, widget Drawable) uint8 {
	return CellTmplAdded
}
