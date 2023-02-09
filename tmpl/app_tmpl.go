package tmpl

import (
	. "github.com/kallaurru/termui/v3"
	"image"
	"sync"
)

type AppTmpl struct {
	Mx    sync.RWMutex
	Focus string
	Mode  bool // true is edit, false read
	Size  image.Rectangle
	Theme *WidgetTheme
	Grid  *Grid
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
		Focus: "",
		Mode:  false,
		Size:  size,
	}
}
