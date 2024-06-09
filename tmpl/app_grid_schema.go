package tmpl

import tui "github.com/kallaurru/termui/v3"

// AppGridSchema - компонент позволяющий сформировать сетку виджетов (tui.Grid) для терминала
// Каждая схема создает или несколько колонок в одной строке или несколько строк в колонке
type AppGridSchema struct {
	rows    int
	cols    int
	deep    int // 0 значит корень // -1 = адаптивный уровень от первого до N
	sizes   []tui.AdaptiveSize
	cells   []interface{}
	isValid bool // признак что схема прошла валидацию
}

func NewAppGridSchema(isRoot, asRows bool, sizes ...tui.AdaptiveSize) (AppGridSchema, bool) {
	var (
		rows, cols, deep int
		full             int
	)

	ags := AppGridSchema{}
	if len(sizes) == 0 {
		return ags, false
	}
	for _, s := range sizes {
		full += s.ToInt()
	}
	if full > tui.AdaptiveSizeMaxInt {
		return ags, false
	}

	tmpSizes := make([]tui.AdaptiveSize, 0, len(sizes))
	tmpSizes = append(tmpSizes, sizes...)
	cells := make([]interface{}, 0, len(sizes))

	if asRows {
		cols = 1
		rows = len(sizes)
	} else {
		cols = len(sizes)
		rows = 1
	}

	if isRoot {
		deep = 0
	} else {
		deep = -1
	}

	ags.cols = cols
	ags.rows = rows
	ags.deep = deep
	ags.sizes = tmpSizes
	ags.cells = cells

	return ags, true
}

func (ags *AppGridSchema) AddItem(val interface{}) (ok bool) {
	defer func() {
		if ok {
			ags.isValid = false
		}
	}()
	switch val.(type) {
	case tui.GridItem:
		ags.cells = append(ags.cells, val)
		return true
	case AppGridSchema:
		ags.cells = append(ags.cells, val)
		return true
	default:
		return false
	}
}

func (ags *AppGridSchema) Validate() bool {
	ags.isValid = true
	return false
}

func (ags *AppGridSchema) Grid(xMin, yMin, xMax, yMax int) (*tui.Grid, bool) {
	grid := tui.NewGrid()
	if ags.deep != 0 {
		return grid, false
	}
	grid.SetRect(xMin, yMin, xMax, yMax)
	items := ags.buildCell()
	grid.Set(items)

	return grid, false
}

func (ags *AppGridSchema) buildCell() []tui.GridItem {
	var items []tui.GridItem

	if ags.cols > 1 {
		items = ags.buildCols()
	}

	if ags.rows > 1 {
		items = ags.buildRows()
	}

	return items
}

func (ags *AppGridSchema) buildCols() []tui.GridItem {
	items := make([]tui.GridItem, 0, len(ags.sizes))
	for i := 0; i < len(ags.sizes); i++ {
		cell := ags.cells[i]

		switch t := cell.(type) {
		case tui.GridItem:
			items = append(items, t)
		case AppGridSchema:
			itemsAGS := t.buildCell()
			items = append(items, itemsAGS...)
		}
	}
	return items
}

func (ags *AppGridSchema) buildRows() []tui.GridItem {

}
