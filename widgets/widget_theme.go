package widgets

import . "github.com/kallaur/termui/v3"

type WidgetTheme struct {
	titleStyle        Style
	borderStyle       Style
	activeTitleStyle  Style
	activeBorderStyle Style
}

//GetTitleStyle - получить стиль заголовка
func (ft *WidgetTheme) GetTitleStyle(isActive bool) Style {
	if isActive {
		return ft.activeTitleStyle
	}

	return ft.titleStyle
}

//GetBorderStyle - получить стиль линии обводки
func (ft *WidgetTheme) GetBorderStyle(isActive bool) Style {
	if isActive {
		return ft.activeBorderStyle
	}

	return ft.borderStyle
}

func NewDefWidgetTheme() *WidgetTheme {
	return &WidgetTheme{
		titleStyle:        NewStyle(ColorBlack, ColorWhite),
		borderStyle:       NewStyle(ColorWhite),
		activeTitleStyle:  NewStyle(ColorBlack, ColorWhite, ModifierBold),
		activeBorderStyle: NewStyle(Color(23)),
	}
}
