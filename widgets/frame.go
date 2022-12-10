package widgets

import (
	. "github.com/kallaur/termui/v3"
)

// Frame Окно с возможностью редактирования текста
type Frame struct {
	Block
	drawBorder bool
	isEditable bool
}
