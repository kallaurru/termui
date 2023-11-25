package widgets

import (
	. "github.com/kallaurru/termui/v3"
)

// Frame Окно с возможностью редактирования текста
type Frame struct {
	Block
	title      string
	drawBorder bool
	isEditable bool
}
