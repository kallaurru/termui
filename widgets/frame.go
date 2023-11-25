package widgets

import (
	. "github.com/kallaurru/termui/v3"
	"image"
)

// Frame Окно с возможностью редактирования текста
type Frame struct {
	Block
	input   []rune
	lastIdx int
	isEdit  bool
}

func NewFrame(center image.Point, title string) *Frame {
	f := &Frame{
		Block:   *NewBlock(),
		input:   make([]rune, 32, 32),
		isEdit:  false,
		lastIdx: 0, // где должен мигать курсор
	}

	f.setBlockRect(center, len(f.input))
	f.Block.Border = true
	f.Block.MakeGlamourTitle(title)

	return f
}

func (f *Frame) Input() string {
	return string(f.input)
}

func (f *Frame) CaptureModeOn() {
	f.isEdit = true
}

func (f *Frame) setBlockRect(center image.Point, limitX int) {
	half := limitX / 2

	yMin := center.Y - 2
	yMax := center.Y + 2
	xMin := center.X - half - 1
	xMax := center.X + half + 1

	f.Block.SetRect(xMin, yMin, xMax, yMax)
}
