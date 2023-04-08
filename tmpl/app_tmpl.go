package tmpl

import (
	. "github.com/kallaurru/termui/v3"
	w "github.com/kallaurru/termui/v3/widgets"
	"image"
	"sync"
)

type AppTmpl struct {
	Mx        sync.RWMutex
	ModeEdit  bool // true is edit, false read
	Size      image.Rectangle
	Theme     *WidgetTheme
	Grid      *Grid
	ChanLog   chan *w.LogRecord
	ChanDraw  chan Drawable
	ChanEvent chan *Event
	focus     string
	storage   map[string]interface{}
}

func NewAppTmpl(isRealBuf bool, chanSizes ...uint8) *AppTmpl {
	var size image.Rectangle
	var chLSize, chDSize, chESize = 8, 8, 8

	if isRealBuf {
		// будет работать только после Init()
		xMax, yMax := TerminalDimensions()
		size = image.Rect(0, 0, xMax, yMax)
	} else {
		size = image.Rect(0, 0, 120, 80)
	}
	if len(chanSizes) > 0 {
		switch len(chanSizes) {
		case 3:
			chLSize, chDSize, chESize = int(chanSizes[0]), int(chanSizes[1]), int(chanSizes[2])
		case 2:
			chLSize, chDSize = int(chanSizes[0]), int(chanSizes[1])
		case 1:
			chLSize, chDSize, chESize = int(chanSizes[0]), int(chanSizes[0]), int(chanSizes[0])
		}
	}
	return &AppTmpl{
		Theme:     NewMyDefaultWidgetTheme(),
		ModeEdit:  false,
		Size:      size,
		ChanLog:   make(chan *w.LogRecord, chLSize),
		ChanDraw:  make(chan Drawable, chDSize),
		ChanEvent: make(chan *Event, chESize),

		focus:   "",
		storage: nil,
	}
}

func (app *AppTmpl) ClearFocus() {
	app.focus = ""
}

func (app *AppTmpl) SetFocus(id string) {
	app.Mx.Lock()
	defer app.Mx.Unlock()

	app.focus = id
}

func (app *AppTmpl) Focus() string {
	app.Mx.RLock()
	defer app.Mx.RUnlock()

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
	app.Mx.Lock()
	defer app.Mx.Unlock()

	app.storage[id] = param
}

func (app *AppTmpl) GetParam(id string) interface{} {
	app.Mx.RLock()
	defer app.Mx.RUnlock()

	val, ok := app.storage[id]
	if !ok {
		return nil
	}

	return val
}

func (app *AppTmpl) Render() {
	app.ChanDraw <- app.Grid
}
