package widgets

import (
	"fmt"
	. "github.com/kallaurru/termui/v3"
	"image"
)

var availableHexKeyMap = map[string][]rune{
	"0": []rune("0"),
	"1": []rune("1"),
	"2": []rune("2"),
	"3": []rune("3"),
	"4": []rune("4"),
	"5": []rune("5"),
	"6": []rune("6"),
	"7": []rune("7"),
	"8": []rune("8"),
	"9": []rune("9"),
	"a": []rune("a"),
	"A": []rune("A"),
	"b": []rune("b"),
	"B": []rune("B"),
	"c": []rune("c"),
	"C": []rune("C"),
	"d": []rune("d"),
	"D": []rune("D"),
	"e": []rune("e"),
	"E": []rune("E"),
	"f": []rune("f"),
	"F": []rune("F"),
	"x": []rune("x"),
	"X": []rune("X"),
}

// Frame Окно с возможностью редактирования текста
type Frame struct {
	Block
	input     []rune
	maxLength int
	isEdit    bool
	title     string
}

func NewFrame(center image.Point, title string) *Frame {
	f := &Frame{
		Block:     *NewBlock(),
		isEdit:    false,
		maxLength: 32,
		title:     title,
	}

	f.Block.Border = true
	f.Block.PaddingTop = 1
	f.Block.PaddingLeft = 1
	f.setBlockRect(center, f.maxLength)
	f.Block.MakeGlamourTitle(title)
	f.clearBuffer()

	return f
}

func (f *Frame) Input() string {
	return string(f.input)
}

func (f *Frame) CaptureModeOn() {
	f.isEdit = true
	f.Block.MakeGlamourTitle(fmt.Sprintf("%s *", f.title))
}

func (f *Frame) CaptureModeOff() {
	f.isEdit = false
	f.Block.MakeGlamourTitle(f.title)
}

func (f *Frame) AddSymbol(code string, ch rune) {
	if !f.isEdit {
		return
	}
	if code == "<Backspace>" {
		if len(f.input) == 0 {
			return
		} else {
			inputLen := len(f.input)
			f.input = f.input[:inputLen-1]
		}
		f.Block.MakeGlamourTitle("Backspace press")
		return
	}
	// проверка наличия свободного места
	if len(f.input) >= f.maxLength-1 {
		f.CaptureModeOff()
		return
	}

	// пробел
	if code == "<Space>" {
		f.input = append(f.input, 32)
		return
	}

	if ch == 0 {
		return
	}

	f.input = append(f.input, ch)
}

func (f *Frame) Draw(buf *Buffer) {
	var str string
	if len(f.input) < f.maxLength-1 && f.isEdit {
		str = fmt.Sprintf("%s%s", string(f.input), string(BARS[1]))
	} else {
		str = string(f.input)
	}

	cells := ParseStyles(str, NewStyle(ColorWhite))

	rows := SplitCells(cells, '\n')

	for y, r := range rows {
		if y+f.Inner.Min.Y >= f.Inner.Max.Y {
			break
		}
		r = TrimCells(r, f.Inner.Dx())
		for _, cx := range BuildCellWithXArray(r) {
			x, cell := cx.X, cx.Cell
			buf.SetCell(cell, image.Pt(x, y).Add(f.Inner.Min))
		}
	}
	f.Block.Draw(buf)
}

func (f *Frame) setBlockRect(center image.Point, limitX int) {
	half := limitX / 2

	yMin := center.Y - 2
	yMax := center.Y + 2
	xMin := center.X - half - 3
	xMax := center.X + half + 3

	f.Block.SetRect(xMin, yMin, xMax, yMax)
}

func (f *Frame) clearBuffer() {
	f.input = make([]rune, 0, f.maxLength)
}
