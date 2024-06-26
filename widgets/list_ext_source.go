package widgets

import (
	. "github.com/kallaurru/termui/v3"
	rw "github.com/mattn/go-runewidth"
	"image"
)

// ListES - список со ссылкой на источник
type ListES struct {
	Block
	Rows             *[]string
	WrapText         bool
	TextStyle        Style
	SelectedRow      int
	topRow           int
	SelectedRowStyle Style
}

func NewListES() *ListES {
	return &ListES{
		Block:            *NewBlock(),
		Rows:             nil,
		TextStyle:        Theme.List.Text,
		SelectedRowStyle: Theme.List.Text,
	}
}

func (self *ListES) Draw(buf *Buffer) {
	var inRows []string
	if self.Rows == nil {
		return
	}
	inRows = *self.Rows
	self.Block.Draw(buf)

	point := self.Inner.Min
	// adjusts view into widget
	if self.SelectedRow >= self.Inner.Dy()+self.topRow {
		self.topRow = self.SelectedRow - self.Inner.Dy() + 1
	} else if self.SelectedRow < self.topRow {
		self.topRow = self.SelectedRow
	}

	// draw rows
	for row := self.topRow; row < len(inRows) && point.Y < self.Inner.Max.Y; row++ {
		cells := ParseStyles(inRows[row], self.TextStyle)
		if self.WrapText {
			cells = WrapCells(cells, uint(self.Inner.Dx()))
		}
		for j := 0; j < len(cells) && point.Y < self.Inner.Max.Y; j++ {
			style := cells[j].Style
			if row == self.SelectedRow {
				style = self.SelectedRowStyle
			}
			if cells[j].Rune == '\n' {
				point = image.Pt(self.Inner.Min.X, point.Y+1)
			} else {
				if point.X+1 == self.Inner.Max.X+1 && len(cells) > self.Inner.Dx() {
					buf.SetCell(NewCell(ELLIPSES, style), point.Add(image.Pt(-1, 0)))
					break
				} else {
					buf.SetCell(NewCell(cells[j].Rune, style), point)
					point = point.Add(image.Pt(rw.RuneWidth(cells[j].Rune), 0))
				}
			}
		}
		point = image.Pt(self.Inner.Min.X, point.Y+1)
	}

	// draw UP_ARROW if needed
	if self.topRow > 0 {
		buf.SetCell(
			NewCell(UP_ARROW, NewStyle(ColorWhite)),
			image.Pt(self.Inner.Max.X-1, self.Inner.Min.Y),
		)
	}

	// draw DOWN_ARROW if needed
	if len(inRows) > int(self.topRow)+self.Inner.Dy() {
		buf.SetCell(
			NewCell(DOWN_ARROW, NewStyle(ColorWhite)),
			image.Pt(self.Inner.Max.X-1, self.Inner.Max.Y-1),
		)
	}
}

// ScrollAmount scrolls by amount given. If amount is < 0, then scroll up.
// There is no need to set self.topRow, as this will be set automatically when drawn,
// since if the selected item is off screen then the topRow variable will change accordingly.
func (self *ListES) ScrollAmount(amount int) {
	if self.Rows == nil {
		return
	}
	inRows := *self.Rows

	if len(inRows)-int(self.SelectedRow) <= amount {
		self.SelectedRow = len(inRows) - 1
	} else if int(self.SelectedRow)+amount < 0 {
		self.SelectedRow = 0
	} else {
		self.SelectedRow += amount
	}
}

func (self *ListES) ScrollUp() {
	self.ScrollAmount(-1)
}

func (self *ListES) ScrollDown() {
	self.ScrollAmount(1)
}

func (self *ListES) ScrollPageUp() {
	// If an item is selected below top row, then go to the top row.
	if self.SelectedRow > self.topRow {
		self.SelectedRow = self.topRow
	} else {
		self.ScrollAmount(-self.Inner.Dy())
	}
}

func (self *ListES) ScrollPageDown() {
	self.ScrollAmount(self.Inner.Dy())
}

func (self *ListES) ScrollHalfPageUp() {
	self.ScrollAmount(-int(FloorFloat64(float64(self.Inner.Dy()) / 2)))
}

func (self *ListES) ScrollHalfPageDown() {
	self.ScrollAmount(int(FloorFloat64(float64(self.Inner.Dy()) / 2)))
}

func (self *ListES) ScrollTop() {
	self.SelectedRow = 0
}

func (self *ListES) ScrollBottom() {
	if self.Rows == nil {
		return
	}
	inRows := *self.Rows
	self.SelectedRow = len(inRows) - 1
}
