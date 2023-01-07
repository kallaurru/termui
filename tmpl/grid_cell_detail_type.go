package tmpl

import . "github.com/kallaurru/termui/v3"

type CellDetail struct {
	idx  uint16 // индекс колонки
	size AdaptiveSize
	data interface{} // или Drawable или *GridSchema
}

func NewCellDetail(row, col uint8, size AdaptiveSize) *CellDetail {
	cd := &CellDetail{
		idx:  0,
		size: size, // не более 100
	}
	cd.makeCellAddr(row, col)

	return cd
}

func (cd *CellDetail) IsSchema() bool {
	_, ok := cd.data.(*GridSchema)

	return ok
}

func (cd *CellDetail) SetWidget(widget Drawable) {
	cd.setData(widget)
}

func (cd *CellDetail) SetSchema(schema *GridSchema) {
	cd.setData(schema)
}

func (cd *CellDetail) GetRowIdx() uint8 {
	return cd.getRowIdx()
}

func (cd *CellDetail) getIdx() uint16 {
	return cd.idx
}

func (cd *CellDetail) getSize() float64 {
	return cd.size.FloatSize()
}

func (cd *CellDetail) makeCellAddr(row uint8, col uint8) {
	var idx uint16 = 0

	idx |= uint16(row) << 8
	idx |= uint16(col)

	cd.idx = idx
}

func (cd *CellDetail) getRowIdx() uint8 {
	tmp := cd.idx >> 8
	tmp &= 0xff

	return uint8(tmp)
}

func (cd *CellDetail) setData(data interface{}) {
	cd.data = data
}
