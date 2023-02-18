package tmpl

import (
	"errors"
	. "github.com/kallaurru/termui/v3"
	"github.com/satori/go.uuid"
)

const errorMsg = "unknown type"

type CellDetail struct {
	idx  uint16       // индекс колонки
	size AdaptiveSize // ширина колонки
	name string       // имя для идентификации виджета на форме
	data interface{}  // или Drawable или *GridSchema
}

func MakeCellDetailWithWidget(row, col uint8, w AdaptiveSize, widget interface{}) *CellDetail {
	cd := NewCellDetail(row, col, w)
	cd.SetWidget(widget)

	return cd
}

func MakeCellDetailWithSchema(row, col uint8, w AdaptiveSize, schema *GridSchema) *CellDetail {
	cd := NewCellDetail(row, col, w)
	cd.SetSchema(schema)

	return cd
}

func NewCellDetail(row, col uint8, size AdaptiveSize) *CellDetail {
	cd := &CellDetail{
		idx:  0,
		size: size, // не более 100
		name: "",
	}

	cd.makeCellAddr(row, col)
	return cd
}

func (cd *CellDetail) IsSchema() bool {
	_, ok := cd.data.(*GridSchema)

	return ok
}

func (cd *CellDetail) GetName() string {
	if cd.name == "" {
		cd.name = uuid.NewV4().String()
	}

	return cd.name
}

func (cd *CellDetail) SetName(name string) {
	cd.name = name
}

func (cd *CellDetail) SetWidget(widget interface{}) {
	cd.setData(widget)
}

func (cd *CellDetail) GetWidget() (Drawable, error) {
	if cd.IsSchema() {
		return nil, errors.New(errorMsg)
	}

	return cd.data.(Drawable), nil
}

func (cd *CellDetail) SetSchema(schema *GridSchema) {
	cd.setData(schema)
}

func (cd *CellDetail) GetSchema() (*GridSchema, error) {
	if cd.IsSchema() {
		return cd.data.(*GridSchema), nil
	}

	return nil, errors.New(errorMsg)
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
