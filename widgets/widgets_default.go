package widgets

import (
	. "github.com/kallaurru/termui/v3"
	"image"
)

func MakeDefaultGauge(gts ...GaugeTheme) *Gauge {
	var gt GaugeTheme

	g := NewGauge()
	if len(gts) > 0 {
		gt = gts[0]
	} else {
		gt = GaugeTheme{
			Bar:   ColorWhite,
			Label: NewStyleBgFree(ColorGreenBgDunkel, ModifierBold),
		}
	}

	g.Percent = 0
	g.UploadTheme(&gt)

	return g
}

func MakeDefaultListES() *ListES {
	cpl := NewListES()
	cpl.Block.SetEmptyHeader()
	cpl.Border = true

	return cpl
}

func MakeDefaultATable(cw []int, align []Alignment, wts ...*WidgetTheme) *ATable {
	var wt *WidgetTheme
	var alignDef = AlignCenter

	if len(wts) > 0 {
		wt = wts[0]
	} else {
		wt = NewMyDefaultWidgetTheme()
	}

	at := NewATable()

	at.WidgetTheme = wt
	at.Block.SetEmptyHeader()
	at.Border = true
	at.RowSeparator = false
	at.ColSeparator = false

	at.ColumnWidths = cw
	alignLen := len(align)
	for i := 0; i < len(cw); i++ {
		if i >= alignLen {
			at.AddColAlignment(i, alignDef)
		} else {
			at.AddColAlignment(i, align[i])
		}
	}

	return at
}

// MakeDefaultHotkeyWindow - всплывающее окно со справочной информацией
// @parent - размер приложения на экране
// @w, @h - размеры в которые должно вписываться это всплывающее окно
// @rows - содержимое
func MakeDefaultHotkeyWindow(parent image.Rectangle, w, h int, cw []int, rows [][]string) *ATable {
	positionWidget := MakeCenterPositionWidget(parent, w, h)
	realColWidth := ConvWidthRelativeToAbs(positionWidget.Min.X, positionWidget.Max.X, cw, []string{""})

	at := MakeDefaultATable(realColWidth, []Alignment{AlignRight, AlignLeft})
	at.MakeGlamourTitle("Application hotkeys")
	at.Rows = rows
	at.SetRect(positionWidget.Min.X, positionWidget.Min.Y, positionWidget.Max.X, positionWidget.Max.Y)
	at.PaddingTop = 1

	return at
}
