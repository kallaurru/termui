package tmpl

import . "github.com/kallaur/termui/v3"

type CellDetail struct {
	idx  uint16 // индекс колонки
	size AdaptiveSize
	draw Drawable // виджет который там должен находится
}

func NewCellDetail(row, col uint8, size AdaptiveSize, widget Drawable) *CellDetail {
	cd := &CellDetail{
		idx:  0,
		size: size, // не более 100
		draw: widget,
	}
	cd.makeCellAddr(row, col)

	return cd
}

func (cd *CellDetail) getIdx() uint16 {
	return cd.idx
}

func (cd *CellDetail) getSize() float64 {
	base := 100

	return float64(cd.size) / float64(base)
}

func (cd *CellDetail) makeCellAddr(row uint8, col uint8) {
	var idx uint16 = 0

	idx |= uint16(row) << 8
	idx |= uint16(col)

	cd.idx = idx
}
