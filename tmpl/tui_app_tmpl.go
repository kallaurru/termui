package tmpl

import (
	. "github.com/kallaurru/termui/v3"
	"image"
)

type kbListener func(params ...interface{}) error

type TUIAppTmpl struct {
	grid  *Grid           // сетка для расположения виджетов перерисовываем все
	theme *WidgetTheme    // тема для оформления виджетов
	size  image.Rectangle // размеры основного окна приложения. Для позиционирования модальных
	// режим редактирования, когда у нас в фокусе виджет в который можно вводить данные с клавиатуры
	// пока все что больше 0 - режим редактирования включен
	modeEdit uint8
	focus    string // идентификатор виджета на котором установлен фокус
	// нам нужно вставить хранилище обработчиков нажатий клавиш
	listener map[string]kbListener
}

// New - создать новый шаблон приложения
func New(isRealBuf bool, theme *WidgetTheme) *TUIAppTmpl {
	var (
		maxX = 120
		maxY = 80
		size image.Rectangle
	)
	if isRealBuf {
		maxX, maxY = TerminalDimensions()
	}
	size = image.Rect(0, 0, maxX, maxY)

	return &TUIAppTmpl{
		grid:     nil,
		theme:    theme,
		size:     size,
		modeEdit: 0,
		focus:    "",
		listener: make(map[string]kbListener),
	}
}

// SetGrid - если сетка была сформирована во вне
func (app *TUIAppTmpl) SetGrid(grid *Grid) {
	app.grid = grid
}

func (app *TUIAppTmpl) SetFocus(id string) {
	app.focus = id
}

func (app *TUIAppTmpl) ClearFocus() {
	app.focus = ""
}

func (app *TUIAppTmpl) GetSize() image.Rectangle {
	return app.size
}

func (app *TUIAppTmpl) Draw(wgd Drawable) {
	Render(wgd)
}

func (app *TUIAppTmpl) Listen() {
	kbEvents := PollEvents()

	for {
		select {
		case kb := <-kbEvents:
			// обрабатываем событие клавиатуры
			handler, ok := app.listener[kb.ID]
			if !ok {
				continue
			}
			err := handler()
			if err != nil {
				// обработка ошибок
			}
		default:

		}
	}
}

func (app *TUIAppTmpl) Close() {
	Clear()
	Close()
}

func (app *TUIAppTmpl) Render() {
	Render(app.grid)
}
