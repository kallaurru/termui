package tmpl

import (
	"container/list"
	. "github.com/kallaurru/termui/v3"
)

const (
	LoggerId = "logger:id"
	HelperId = "helper:id"
)

/**
Весь макет делим на строки, далее в ячейках могут быть вложенные схемы
*/

// GridSchema - схема для создания для адаптивной сетки
// Deprecated: нужно использовать AppGridSchema
type GridSchema struct {
	rowsSizes map[uint8]AdaptiveSize
	cells     map[uint8]*list.List // в списке или схема или детали (GridSchema, CellDetail)
	idx       map[string]Drawable
}

// NewGridSchema - указываем ширину строк. В сумме должно быть не больше 100%
func NewGridSchema(sizes ...AdaptiveSize) *GridSchema {
	gs := &GridSchema{
		rowsSizes: make(map[uint8]AdaptiveSize),
		cells:     make(map[uint8]*list.List),
		idx:       make(map[string]Drawable),
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

func (gs *GridSchema) AddCell(rowInSchema uint8, cd *CellDetail) {
	name := cd.GetName()
	_, ok := gs.idx[name]
	if !ok {
		widget, err := cd.GetWidget()
		if err != nil {
			if cd.IsSchema() {
				schema, _ := cd.GetSchema()
				gs.mergeWidgets(schema)
				gs.addCell(rowInSchema, cd)
			}
			return
		}
		gs.idx[name] = widget
	}
	gs.addCell(rowInSchema, cd)
}

func (gs *GridSchema) AddSchema(rowInSchema uint8, schema *GridSchema) {
	gs.mergeWidgets(schema)
	gs.addCell(rowInSchema, schema)
}

// Build - по количеству строк в схеме
func (gs *GridSchema) Build() []interface{} {
	lenCells := len(gs.cells)
	out := make([]interface{}, 0, lenCells)
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

func (gs *GridSchema) BuildGrid(maxX, maxY int) *Grid {
	items := gs.Build()
	grid := NewGrid()
	grid.SetRect(0, 0, maxX, maxY)
	grid.Set(items...)

	return grid
}

func (gs *GridSchema) GetWidgets() map[string]Drawable {
	return gs.idx
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
	gs.cells[row] = l
}

func (gs *GridSchema) compile(row uint8) GridItem {
	cells := gs.cells[row]
	items := make([]interface{}, 0, cells.Len())
	for e := cells.Front(); e != nil; e = e.Next() {
		value := e.Value
		cd, _ := value.(*CellDetail)
		if cd.IsSchema() {
			schema, _ := cd.data.(*GridSchema)
			gi := NewCol(cd.getSize(), schema.Build()...)
			items = append(items, gi)
		} else {
			widget, _ := cd.data.(Drawable)
			gi := NewCol(cd.getSize(), widget)
			items = append(items, gi)
		}
	}

	size := gs.rowsSizes[row]
	return NewRow(size.FloatSize(), items...)
}

func (gs *GridSchema) mergeWidgets(schema *GridSchema) {
	widgets := schema.GetWidgets()
	for name, drw := range widgets {
		_, ok := gs.idx[name]
		if !ok {
			gs.idx[name] = drw
		}
	}
}
