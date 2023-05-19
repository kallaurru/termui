package tmpl

import (
	. "github.com/kallaurru/termui/v3"
	w "github.com/kallaurru/termui/v3/widgets"
)

func MakeTableGroup(cw []int, align []Alignment, wts ...*WidgetTheme) (*w.ATable, *DataProviderTable) {
	widget := w.MakeDefaultATable(cw, align, wts...)
	dp := NewDataProviderTable()
	dp.SetMaxColsCount(uint8(len(cw)))

	return widget, dp
}
