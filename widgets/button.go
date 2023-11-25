package widgets

import (
	"fmt"
	. "github.com/kallaurru/termui/v3"
	"image"
	"strings"
)

const limitLabel = 6

type Button struct {
	Block
	wt    *WidgetTheme
	label string
}

func NewButton(center image.Point) *Button {
	def := "OK"
	wt := NewButtonDefaultTheme()
	b := &Button{
		Block: *NewBlock(),
		label: def,
		wt:    wt,
	}

	b.Block.Border = true
	b.makeSizes(center, limitLabel)
	b.Active()
	// центрируем надпись
	b.ChangeLabel(def)

	return b
}

func (b *Button) Active() {
	b.Block.BorderStyle = b.wt.GetBorderStyle(true)
}

func (b *Button) NoActive() {
	b.Block.BorderStyle = b.wt.GetBorderStyle(false)
}

func (b *Button) ChangeLabel(label string) {
	labelLen := len(label)

	if len(label) > limitLabel {
		ru := []rune(label)
		ru = ru[:6]
		label = string(ru)
	}

	diff := limitLabel - labelLen

	if diff >= 2 {
		add := diff / 2
		label = fmt.Sprintf(
			"%s%s%s",
			strings.Repeat(" ", add),
			label, strings.Repeat(" ", add))
	}

	b.label = label
}

func (b *Button) Draw(buf *Buffer) {
	b.Block.Draw(buf)

	cells := ParseStyles(b.label, b.wt.GetTitleStyle(false))

	rows := SplitCells(cells, '\n')

	for y, r := range rows {
		if y+b.Inner.Min.Y >= b.Inner.Max.Y {
			break
		}
		r = TrimCells(r, b.Inner.Dx())
		for _, cx := range BuildCellWithXArray(r) {
			x, cell := cx.X, cx.Cell
			buf.SetCell(cell, image.Pt(x, y).Add(b.Inner.Min))
		}
	}
}

func (b *Button) makeSizes(center image.Point, limitLabel int) {
	half := limitLabel / 2
	ymin := center.Y - 2
	ymax := center.Y + 1
	xmin := center.X - half - 1
	xmax := center.X + half + 1

	b.Block.SetRect(xmin, ymin, xmax, ymax)
}
