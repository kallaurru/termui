package widgets

import (
	. "github.com/kallaurru/termui/v3"
	"image"
)

// ATable - таблица с расширенным функционалом форматирования
type ATable struct {
	Table
	WidgetTheme      *WidgetTheme
	ColSeparator     bool
	ColTextAlignment map[int]Alignment
	maxColIdx        int
	useCellAlignment bool
	callerUniqKey    func(idxRow, idxCol, idxMaxCol int) int
}

func NewATable() *ATable {
	return &ATable{
		Table: Table{
			Block:         *NewBlock(),
			TextStyle:     Theme.Table.Text,
			RowSeparator:  true,
			RowStyles:     make(map[int]Style),
			ColumnResizer: func() {},
		},
		WidgetTheme:      NewDefWidgetTheme(),
		maxColIdx:        0,
		useCellAlignment: false,
		callerUniqKey: func(idxRow, idxCol, idxMaxCol int) int {
			return idxRow*idxMaxCol + idxCol + 1
		},
	}
}

func (at *ATable) UseCellAlignment() {
	at.useCellAlignment = true
}

func (at *ATable) AddColAlignment(colIdx int, alignment Alignment) {
	if !at.useCellAlignment {
		at.ColTextAlignment[colIdx] = alignment
	}
}

func (at *ATable) AddCellAlignment(idxRow, idxCol int, alignment Alignment) {
	if at.useCellAlignment {
		uniq := at.callerUniqKey(idxRow, idxCol, at.maxColIdx)
		at.ColTextAlignment[uniq] = alignment
	}
}

// AddColAlignmentHelperTheme - желательно вызвать после того как заполнено свойство Rows
func (at *ATable) AddColAlignmentHelperTheme() {
	if !at.useCellAlignment {
		at.ColTextAlignment[0] = AlignRight
		at.ColTextAlignment[1] = AlignLeft
		return
	}
	for row, _ := range at.Rows {
		// таблица 2 колонки
		at.ColTextAlignment[at.callerUniqKey(row, 0, at.maxColIdx)] = AlignRight
		at.ColTextAlignment[at.callerUniqKey(row, 1, at.maxColIdx)] = AlignLeft
	}
}

func (at *ATable) GetCellAlignment(idxRow, idxCol int) Alignment {
	if at.useCellAlignment {
		uniq := at.callerUniqKey(idxRow, idxCol, at.maxColIdx)
		alignment, ok := at.ColTextAlignment[uniq]
		if ok {
			return alignment
		}
	}

	alignment, ok := at.ColTextAlignment[idxCol]
	if ok {
		return alignment
	}

	return at.TextAlignment
}

func (at *ATable) Draw(buf *Buffer) {
	at.Table.Block.Draw(buf)

	at.ColumnResizer()
	// определяем максимальный индекс колонки, для работы с расположением текста в ячейках
	if at.Rows != nil {
		columnCount := len(at.Rows[0])
		if columnCount > 0 {
			at.maxColIdx = columnCount - 1
		} else {
			at.maxColIdx = 0
		}
	}

	columnWidths := at.ColumnWidths
	if len(columnWidths) == 0 {
		columnCount := len(at.Rows[0])
		columnWidth := at.Inner.Dx() / columnCount
		for i := 0; i < columnCount; i++ {
			columnWidths = append(columnWidths, columnWidth)
		}
	}

	yCoordinate := at.Inner.Min.Y

	for i := 0; i < len(at.Rows) && yCoordinate < at.Inner.Max.Y; i++ {
		row := at.Rows[i]
		colXCoordinate := at.Inner.Min.X

		rowStyle := at.TextStyle
		// get the row style if one exists
		if style, ok := at.RowStyles[i]; ok {
			rowStyle = style
		}

		if at.FillRow {
			blankCell := NewCell(' ', rowStyle)
			buf.Fill(blankCell, image.Rect(at.Inner.Min.X, yCoordinate, at.Inner.Max.X, yCoordinate+1))
		}

		for j := 0; j < len(row); j++ {
			col := ParseStyles(row[j], rowStyle)
			// draw row cell
			if len(col) > columnWidths[j] || at.GetCellAlignment(i, j) == AlignLeft {
				for _, cx := range BuildCellWithXArray(col) {
					k, cell := cx.X, cx.Cell
					if k == columnWidths[j] || colXCoordinate+k == at.Inner.Max.X {
						cell.Rune = ELLIPSES
						buf.SetCell(cell, image.Pt(colXCoordinate+k-1, yCoordinate))
						break
					} else {
						buf.SetCell(cell, image.Pt(colXCoordinate+k, yCoordinate))
					}
				}
			} else if at.GetCellAlignment(i, j) == AlignCenter {
				xCoordinateOffset := (columnWidths[j] - len(col)) / 2
				stringXCoordinate := xCoordinateOffset + colXCoordinate
				for _, cx := range BuildCellWithXArray(col) {
					k, cell := cx.X, cx.Cell
					buf.SetCell(cell, image.Pt(stringXCoordinate+k, yCoordinate))
				}
			} else if at.GetCellAlignment(i, j) == AlignRight {
				stringXCoordinate := MinInt(colXCoordinate+columnWidths[j], at.Inner.Max.X) - len(col)
				for _, cx := range BuildCellWithXArray(col) {
					k, cell := cx.X, cx.Cell
					buf.SetCell(cell, image.Pt(stringXCoordinate+k, yCoordinate))
				}
			}
			colXCoordinate += columnWidths[j] + 1
		}

		// draw vertical separators
		separatorStyle := at.Block.BorderStyle

		separatorXCoordinate := at.Inner.Min.X
		verticalCell := NewCell(rune(32), separatorStyle)
		if at.ColSeparator {
			verticalCell = NewCell(VERTICAL_LINE, separatorStyle)
		}
		for i, width := range columnWidths {
			if at.FillRow && i < len(columnWidths)-1 {
				verticalCell.Style.Bg = rowStyle.Bg
			} else {
				verticalCell.Style.Bg = at.Block.BorderStyle.Bg
			}

			separatorXCoordinate += width
			buf.SetCell(verticalCell, image.Pt(separatorXCoordinate, yCoordinate))
			separatorXCoordinate++
		}

		yCoordinate++

		// draw horizontal separator
		horizontalCell := NewCell(HORIZONTAL_LINE, separatorStyle)
		if at.RowSeparator && yCoordinate < at.Inner.Max.Y && i != len(at.Rows)-1 {
			buf.Fill(horizontalCell, image.Rect(at.Inner.Min.X, yCoordinate, at.Inner.Max.X, yCoordinate+1))
			yCoordinate++
		}
	}
}
