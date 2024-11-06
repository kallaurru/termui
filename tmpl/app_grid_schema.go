package tmpl

import (
	"errors"
	tui "github.com/kallaurru/termui/v3"
)

// AppGridSchema - компонент позволяющий сформировать сетку виджетов (tui.Grid) для терминала
// Каждая схема компилируется или в строку или в колонку
type AppGridSchema struct {
	items int
	deep  int // 0 значит корень // -1 = адаптивный уровень от первого до N
	asRow bool
	sizes []tui.AdaptiveSize
	cells []interface{} // AppGridSchema or GridItem
}

func NewAppGridSchema(isRoot, asRows bool, sizes ...tui.AdaptiveSize) (*AppGridSchema, bool) {
	var (
		full, deep int
	)

	ags := &AppGridSchema{}
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

	if isRoot {
		deep = 0
	} else {
		deep = -1
	}

	ags.asRow = asRows
	ags.items = len(sizes)
	ags.deep = deep
	ags.sizes = tmpSizes
	ags.cells = cells

	return ags, true
}

func (ags *AppGridSchema) AddItem(val interface{}) bool {
	if !ags.hasFreePlace() {
		return false
	}
	switch t := val.(type) {
	case tui.Drawable:
		ags.addGridItem(t)
		return true
	case *AppGridSchema:
		ok := t.SetDeep(ags.deep)
		if ok {
			ags.cells = append(ags.cells, t)
			return true
		}
		return false
	default:
		return false
	}

}

func (ags *AppGridSchema) Grid(xMin, yMin, xMax, yMax int) (*tui.Grid, bool) {
	grid := tui.NewGrid()
	if ags.deep != 0 {
		return grid, false
	}

	var items []interface{}

	grid.SetRect(xMin, yMin, xMax, yMax)

	for i := 0; i < len(ags.cells); i++ {
		value := ags.cells[i]
		if itemOfType, ok := value.(tui.GridItem); ok {
			if ags.asRow {
				items = append(items, tui.NewRow(ags.sizes[i].FloatSize(), itemOfType))
			} else {
				items = append(items, tui.NewCol(ags.sizes[i].FloatSize(), itemOfType))
			}
			continue
		}

		if itemOfType, ok := value.(*AppGridSchema); ok {
			inSchemaItems := itemOfType.compileSchema()
			if ags.asRow {
				items = append(items, tui.NewRow(ags.sizes[i].FloatSize(), inSchemaItems...))
			} else {
				items = append(items, tui.NewCol(ags.sizes[i].FloatSize(), inSchemaItems...))
			}
		}
	}

	grid.Set(items...)

	return grid, true
}

func (ags *AppGridSchema) SetDeep(ownerDeep int) bool {
	var maxDeep = 2 // максимум три уровня вложенности
	deep := ownerDeep + 1
	if deep <= maxDeep {
		ags.deep = deep
		return true
	}
	return false
}

// вариант когда нужно все вложенные в схему схемы и/или компоненты собрать в один массив
// в указанном порядке
func (ags *AppGridSchema) compileSchema() []interface{} {
	var localItems []interface{}

	for i := 0; i < len(ags.cells); i++ {
		value := ags.cells[i]
		if itemType, ok := value.(tui.GridItem); ok {
			localItems = append(localItems, itemType)
			continue
		}

		if itemType, ok := value.(*AppGridSchema); ok {
			schemaItems := itemType.compileSchema()
			if ags.asRow {
				localItems = append(localItems, tui.NewRow(ags.sizes[i].FloatSize(), schemaItems...))
				continue
			}
			localItems = append(localItems, tui.NewCol(ags.sizes[i].FloatSize(), schemaItems...))
		}
	}

	return localItems
}

func (ags *AppGridSchema) hasFreePlace() bool {
	return ags.items > 0 && ags.items > len(ags.cells)
}

func (ags *AppGridSchema) addGridItem(w tui.Drawable) {
	var item tui.GridItem

	nextSize, err := ags.nextSize()
	if err != nil {
		return
	}
	if ags.asRow {
		item = tui.NewCol(nextSize.FloatSize(), w)
	} else {
		item = tui.NewRow(nextSize.FloatSize(), w)
	}
	ags.cells = append(ags.cells, item)
}

func (ags *AppGridSchema) nextSize() (tui.AdaptiveSize, error) {
	if !ags.hasFreePlace() {
		return tui.NewAdaptiveSizeZero(), errors.New("free space is over")
	}
	return ags.sizes[len(ags.cells)], nil
}

func (ags *AppGridSchema) hasNextLevel() bool {
	var maxDeep int = 2 // максимум три уровня вложенности
	return ags.deep < maxDeep
}
