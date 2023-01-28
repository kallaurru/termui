package widgets

import . "github.com/kallaurru/termui/v3"

type WidgetTheme struct {
	titleStyle        Style
	borderStyle       Style
	activeTitleStyle  Style
	activeBorderStyle Style
}

//GetTitleStyle - получить стиль заголовка
func (wt *WidgetTheme) GetTitleStyle(isActive bool) Style {
	if isActive {
		return wt.activeTitleStyle
	}

	return wt.titleStyle
}

//GetBorderStyle - получить стиль линии обводки
func (wt *WidgetTheme) GetBorderStyle(isActive bool) Style {
	if isActive {
		return wt.activeBorderStyle
	}

	return wt.borderStyle
}

func NewDefWidgetTheme() *WidgetTheme {
	return &WidgetTheme{
		titleStyle:        NewStyle(ColorBlack, ColorWhite),
		borderStyle:       NewStyle(ColorWhite),
		activeTitleStyle:  NewStyle(ColorBlack, ColorWhite, ModifierBold),
		activeBorderStyle: NewStyle(ColorGreenBlue),
	}
}

func NewMyDefaultWidgetTheme() *WidgetTheme {
	return &WidgetTheme{
		titleStyle:        NewStyle(ColorWhite),
		borderStyle:       NewStyle(ColorWhite),
		activeTitleStyle:  NewStyle(ColorGreenBlue, ModifierBold),
		activeBorderStyle: NewStyle(ColorGreenBlue),
	}
}
