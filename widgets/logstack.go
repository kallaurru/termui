package widgets

import (
	"container/list"
	. "github.com/kallaurru/termui/v3"
	"image"
)

const (
	logStackDefDepth    uint8 = 5
	logStackDefDepthMax uint8 = 12
)

type LogStack struct {
	Block
	depth  uint8 // кол-во линий отображаемых для пользователя
	stack  *list.List
	fields []int
}

func NewLogStack(depth ...uint8) *LogStack {
	d := logStackDefDepth
	if len(depth) > 0 {
		d = depth[0]
		if d > logStackDefDepthMax {
			d = logStackDefDepthMax
		}
	}
	ls := &LogStack{
		Block:  *NewBlock(),
		depth:  d,
		stack:  list.New(),
		fields: []int{16, 4, -1}, // 0 и -1 можно ставить на поля в строках которых нет описания стилей
	}

	ls.Block.Border = false

	return ls
}

/** ИНТЕРФЕЙС Widget */

func (ls *LogStack) GetId() string {
	return ls.Block.GetName()
}

func (ls *LogStack) SetId(name string) {
	ls.Block.SetName(name)
}

func (ls *LogStack) AddLogRecordAsIs(lr LogRecord) {
	ls.stack.PushFront(lr)
	ls.manageStack()
}

func (ls *LogStack) AddInfoLogRecord(message string) {
	ls.addLogRecord(message, LogRecTypeInfo)
}

func (ls *LogStack) AddWarnLogRecord(message string) {
	ls.addLogRecord(message, LogRecTypeWarn)
}

func (ls *LogStack) AddErrLogRecord(err error) {
	ls.addLogRecord(err.Error(), LogRecTypeErr)
}

func (ls *LogStack) Draw(buf *Buffer) {
	ls.Block.Draw(buf)

	fLogRecToCells := func(rec *LogRecord) [][]Cell {
		return [][]Cell{
			ParseStyles(rec.GetMomentFormat(), Theme.Table.Text),
			ParseStyles(rec.GetRecType(), Theme.Table.Text),
			ParseStyles(rec.GetMsg(), Theme.Table.Text),
		}
	}

	y := ls.Inner.Min.Y // сквозная координата y по всей площади
	// проходим по записям в хранилище. Строки в таблице
	for el := ls.stack.Front(); el != nil; el = el.Next() {
		x := ls.Inner.Min.X // сквозная координата по всей площади
		rec := el.Value.(LogRecord)
		// колонки в строке
		cols := fLogRecToCells(&rec)
		ls.fields = ConvWidthRelativeToAbs(
			ls.Inner.Min.X,
			ls.Inner.Max.X,
			ls.fields,
			[]string{rec.GetMomentFormat(), rec.GetRecType(), rec.GetMsg()},
		)

		for jc, col := range cols {
			// за каждый проход заполняем одну колонку в строке
			maxXField := ls.fields[jc] - 1
			maxXData := len(col) - 1
			for _, cx := range BuildCellWithXArray(col) {
				k, cell := cx.X, cx.Cell
				buf.SetCell(cell, image.Pt(x+k, y))
				// случай если данные превышают длину поля
				if k > maxXField {
					// вставляем заполнитель превышения данных длины поля
					buf.SetCell(NewCell(ELLIPSES, Theme.Table.Text), image.Pt(x+k-1, y))
					break
				}
			}
			// позиционируем X. Если длина данных больше заявленной длины поля, то позиционируем
			// по длине поля.
			if maxXField > maxXData {
				tmp := x + maxXData
				diff := maxXField - maxXData
				for i := 0; i < diff; i++ {
					buf.SetCell(NewCell(' '), image.Pt(tmp+1, y))
				}
			}
			x += maxXField + 2 // 2 место под разделитель колонок
		}
		y++
	}
}

func (ls *LogStack) addLogRecord(message, t string) {
	rec := NewLogRecord(message, t)
	ls.stack.PushFront(rec)
	ls.manageStack()
}

func (ls *LogStack) manageStack() {
	if ls.stack.Len() > int(ls.depth) {
		el := ls.stack.Back()
		ls.stack.Remove(el)
	}
}
