package widgets

import (
	. "github.com/kallaurru/termui/v3"
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
	at.PaddingTop = 1
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
