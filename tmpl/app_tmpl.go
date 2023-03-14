package tmpl

import (
	. "github.com/kallaurru/termui/v3"
	"image"
	"sync"
)

type AppTmpl struct {
	Mx    sync.RWMutex
	Mode  bool // true is edit, false read
	Size  image.Rectangle
	Theme *WidgetTheme
	Grid  *Grid

	focus string
}

func NewAppTmpl(isRealBuf bool) AppTmpl {
	var size image.Rectangle
	if isRealBuf {
		// будет работать только после Init()
		xMax, yMax := TerminalDimensions()
		size = image.Rect(0, 0, xMax, yMax)
	} else {
		size = image.Rect(0, 0, 120, 80)
	}

	return AppTmpl{
		Theme: NewMyDefaultWidgetTheme(),
		focus: "",
		Mode:  false,
		Size:  size,
	}
}

func (app *AppTmpl) ClearFocus() {
	app.focus = ""
}

func (app *AppTmpl) SetFocus(id string) {
	app.focus = id
}
func (app *AppTmpl) Focus() string {
	return app.focus
}
