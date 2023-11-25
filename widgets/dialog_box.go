package widgets

import (
	. "github.com/kallaurru/termui/v3"
	"image"
)

const limitMsgLine = 32

type DialogBox struct {
	Block
	msg    *Paragraph
	button *Button
}

func NewDialogBox(center image.Point, title, msg string) *DialogBox {
	db := &DialogBox{
		Block:  *NewBlock(),
		msg:    NewParagraph(),
		button: NewButton(image.Pt(center.X, center.Y+5)),
	}

	db.msg.Text = msg
	db.msg.Block.Border = false

	db.makeBlockRect(center)
	db.Block.MakeGlamourTitle(title)

	return db
}

func (db *DialogBox) SetButtonLabel(label string) {
	db.button.ChangeLabel(label)
}

func (db *DialogBox) Draw(buf *Buffer) {
	db.Block.Draw(buf)
	db.msg.Draw(buf)
	db.button.Draw(buf)
}

func (db *DialogBox) makeBlockRect(center image.Point) {
	half := limitMsgLine / 2
	xMin := center.X - half - 2
	xMax := center.X + half + 2
	yMax := db.button.Block.Max.Y + 2
	yMin := center.Y - 8
	db.Block.SetRect(xMin, yMin, xMax, yMax)
	db.msg.Block.SetRect(xMin+1, yMin+2, xMax-1, center.Y-2)
}
