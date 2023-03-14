package tmpl

import (
	. "github.com/kallaurru/termui/v3"
	w "github.com/kallaurru/termui/v3/widgets"
	"image"
	"sync"
)

type AppTmpl struct {
	Mx        sync.RWMutex
	Mode      bool // true is edit, false read
	Size      image.Rectangle
	Theme     *WidgetTheme
	Grid      *Grid
	ChanLog   chan *w.LogRecord
	ChanDraw  chan Drawable
	ChanEvent chan Event
	focus     string
	storage   map[string]interface{}
}

func NewAppTmpl(isRealBuf bool) *AppTmpl {
	var size image.Rectangle
	if isRealBuf {
		// будет работать только после Init()
		xMax, yMax := TerminalDimensions()
		size = image.Rect(0, 0, xMax, yMax)
	} else {
		size = image.Rect(0, 0, 120, 80)
	}

	return &AppTmpl{
		Theme:     NewMyDefaultWidgetTheme(),
		Mode:      false,
		Size:      size,
		ChanLog:   make(chan *w.LogRecord, 8),
		ChanDraw:  make(chan Drawable),
		ChanEvent: make(chan Event),

		focus:   "",
		storage: nil,
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

func (app *AppTmpl) Close() {
	close(app.ChanLog)
	close(app.ChanDraw)
}

func (app *AppTmpl) Info(msg string) {
	app.ChanLog <- w.NewLogRecordPtr(msg, w.LogRecTypeInfo)
}

func (app *AppTmpl) Warn(msg string) {
	app.ChanLog <- w.NewLogRecordPtr(msg, w.LogRecTypeWarn)
}

func (app *AppTmpl) Err(msg string) {
	app.ChanLog <- w.NewLogRecordPtr(msg, w.LogRecTypeErr)
}

func (app *AppTmpl) AddToStorage(id string, param interface{}) {
	if app.storage == nil {
		app.storage = make(map[string]interface{})
	}

	app.storage[id] = param
}

func (app *AppTmpl) GetParam(id string) interface{} {
	val, ok := app.storage[id]
	if !ok {
		return nil
	}

	return val
}
