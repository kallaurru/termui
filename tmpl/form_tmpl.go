package tmpl

import (
	. "github.com/kallaur/termui/v3"
	. "github.com/kallaur/termui/v3/widgets"
)

// FormTmpl - шаблоны для создания ui-приложений
type FormTmpl struct {
	Block
	grid     *Grid
	gridTmpl *GridTmpl
	theme    *WidgetTheme
}

func NewFormTmpl(isRealBuf bool, rows uint8, rSizes ...uint8) *FormTmpl {
	var maxX, maxY int
	fTmpl := &FormTmpl{
		Block:    *NewBlock(),
		grid:     NewGrid(),
		gridTmpl: NewGridTmpl(rows),
	}
	if isRealBuf {
		maxX, maxY = TerminalDimensions()
	} else {
		maxX, maxY = 120, 80
	}

	fTmpl.SetRect(0, 0, maxX, maxY)

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

func (ft *FormTmpl) AddGrid(grid *Grid) *FormTmpl {
	ft.grid = grid

	return ft
}

func (ft *FormTmpl) AddGridTmpl(grTmpl *GridTmpl) *FormTmpl {
	ft.gridTmpl = grTmpl

	return ft
}
