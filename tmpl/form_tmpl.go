package tmpl

import (
	"container/list"
	. "github.com/kallaurru/termui/v3"
	. "github.com/kallaurru/termui/v3/widgets"
	"sync"
)

// FormTmpl - шаблоны для создания ui-приложений
type FormTmpl struct {
	Block
	mtx      *sync.RWMutex
	loggerId string
	hlpId    string
	editMode bool
	focus    *list.Element
	grid     *Grid
	theme    *WidgetTheme
	// просто список виджетов, которые присутствуют в форме, включая всплывающие окна
	// далее везде будут ссылки на узлы данного списка
	widgets  *list.List
	handlers *list.List // EventHandler

	idxWidget  map[string]*list.Element
	idxHandler map[string]*list.Element
}

func NewFormTmpl(isRealBuf bool, schema *GridSchema) *FormTmpl {
	fSize := func(isReal bool) (int, int) {
		if isRealBuf {
			return TerminalDimensions()
		} else {
			// потом можно заменить на значения из конфигурации
			return 120, 80
		}
	}

	fTmpl := &FormTmpl{
		Block:      *NewBlock(),
		mtx:        &sync.RWMutex{},
		loggerId:   LoggerId,
		hlpId:      HelperId,
		editMode:   false,
		grid:       nil,
		focus:      nil,
		widgets:    list.New(),
		handlers:   list.New(),
		idxWidget:  make(map[string]*list.Element),
		idxHandler: make(map[string]*list.Element),
	}

	maxX, maxY := fSize(isRealBuf)
	fTmpl.SetRect(0, 0, maxX, maxY)
	formInner := fTmpl.Inner

	fTmpl.grid = schema.BuildGrid(formInner.Max.X, formInner.Max.Y)
	// добавляем виджеты из схемы
	widgetMap := schema.GetWidgets()
	for name, widget := range widgetMap {
		element := fTmpl.widgets.PushBack(widget)
		fTmpl.idxWidget[name] = element
	}

	return fTmpl
}

func (ft *FormTmpl) AddTitle(title string) *FormTmpl {
	ft.Block.MakeGlamourTitle(title)

	return ft
}

func (ft *FormTmpl) AddTheme(theme *WidgetTheme) *FormTmpl {
	ft.theme = theme

	return ft
}

func (ft *FormTmpl) SetHelperId(id string) {
	ft.hlpId = id
}

func (ft *FormTmpl) SetLoggerId(id string) {
	ft.loggerId = id
}

func (ft *FormTmpl) EditModeOn() {
	ft.editMode = true
}

func (ft *FormTmpl) EditModeOff() {
	ft.editMode = false
}

func (ft *FormTmpl) AddHandler(name string, handler EventHandler) {
	ft.mtx.Lock()
	defer ft.mtx.Unlock()

	_, ok := ft.idxHandler[name]
	if ok {
		return
	}
	elem := ft.handlers.PushBack(handler)
	ft.idxHandler[name] = elem
}

func (ft *FormTmpl) DelHandler(name string) {
	elem, ok := ft.idxHandler[name]
	if !ok {
		return
	}

	ft.handlers.Remove(elem)
}
