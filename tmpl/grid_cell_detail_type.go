package tmpl

import . "github.com/kallaur/termui/v3"

const (
	CellSizePercent30       uint8 = 30
	CellSizePercentile      uint8 = 25
	CellSizeTwoPercentile   uint8 = 50
	CellSizeThreePercentile uint8 = 75
	CellSizeFull            uint8 = 100
)

type CellDetail struct {
	idx  uint16   // индекс колонки
	size uint8    // число не более 10 включительно
	draw Drawable // виджет который там должен находится
}

func NewCellDetail(row, col, size uint8, widget Drawable) *CellDetail {
	cd := &CellDetail{
		idx:  0,
		size: 0, // не более 100
		draw: widget,
	}

	if size > 100 {
		return nil
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
