package tmpl

import (
	. "github.com/kallaurru/termui/v3"
	. "github.com/kallaurru/termui/v3/widgets"
)

// FormTmpl - шаблоны для создания ui-приложений
type FormTmpl struct {
	Block
	grid   *Grid
	schema *GridSchema
	theme  *WidgetTheme
}

func NewFormTmpl(isRealBuf bool, schema *GridSchema) *FormTmpl {
	var maxX, maxY int
	fTmpl := &FormTmpl{
		Block:  *NewBlock(),
		grid:   nil,
		schema: schema,
	}
	if isRealBuf {
		maxX, maxY = TerminalDimensions()
	} else {
		maxX, maxY = 120, 80
	}

	fTmpl.SetRect(0, 0, maxX, maxY)

	grid := fTmpl.schema.BuildGrid()
	formInner := fTmpl.Inner
	grid.SetRect(formInner.Min.X, formInner.Min.Y, formInner.Max.X, formInner.Max.Y)

	return fTmpl
}

func (ft *FormTmpl) AddTitle(title string) *FormTmpl {
	ft.Title = title

	return ft
}

func (ft *FormTmpl) AddTheme(theme *WidgetTheme) *FormTmpl {
	ft.theme = theme

	return ft
}
