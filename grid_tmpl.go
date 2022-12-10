package termui

const (
	CellTmplAdded     uint8 = 0x00
	CellTmplBadRowIdx uint8 = 0x01
	CellIsExists      uint8 = 0x02
	CellOtherErr      uint8 = 0x04
)

type GridTmpl struct {
	rows  uint8 // количество строк в шаблоне
	ccr   uint8 // количество столбцов в строке по умолчанию
	cells map[uint16]*CellDetail
}

type CellDetail struct {
	idx  uint16   // индекс колонки
	size uint8    // число не более 10 включительно
	draw Drawable // виджет который там должен находится
}

func NewGridTmpl(rows uint8) *GridTmpl {
	return &GridTmpl{
		rows:  rows,
		ccr:   1,
		cells: make(map[uint16]*CellDetail),
	}
}

// AddCell - индексация строк и столбцов начинается с 0
func (gt *GridTmpl) AddCell(row, col uint8, widget Drawable) uint8 {
	if row > gt.rows {
		return CellTmplBadRowIdx
	}

	idx := gt.makeCellAddr(row, col)
	_, ok := gt.cells[idx]
	if ok {
		return CellIsExists
	}
	gt.cells[idx] = widget
}

func (gt *GridTmpl) makeCellAddr(row uint8, col uint8) uint16 {
	var idx uint16 = 0

	idx |= uint16(row) << 8
	idx |= uint16(col)

	return idx
}
