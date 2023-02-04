package main

import (
	"github.com/kallaurru/termui/v3"
	"github.com/kallaurru/termui/v3/widgets"
	"image"
	"log"
	"time"
)

/*
Прототип приборной доски для логбука
*/
func main() {

	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	t := widgets.NewATable()
	t.RowSeparator = false
	t.ColSeparator = false
	t.ColumnWidths = []int{3, 3, 3, 3}
	t.AddColAlignment(0, termui.AlignCenter)
	t.AddColAlignment(1, termui.AlignCenter)
	t.AddColAlignment(2, termui.AlignCenter)
	maxX, maxY := termui.TerminalDimensions()
	w, h := 40, 8
	parent := image.Rect(0, 0, maxX, maxY)
	dim := termui.MakeCenterPositionWidget(parent, w, h)

	t.SetRect(dim.Min.X, dim.Min.Y, dim.Max.X, dim.Max.Y)
	dp := makeDashboardDP()
	// linking
	t.Rows = dp.GetTable()

	termui.Render(t)

	ticker := time.NewTicker(time.Second * 2).C
	begin := 0x2581
	end := 0x2587
	current := 0x2581
	uiEvents := termui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return

			}
		case <-ticker:
			if (current + 1) > end {
				current = begin
			} else {
				current += 1
			}
			dp.UpdateData(makeDashboardItem(int32(current)), 1, 0, 0)
			termui.Render(t)

		}
	}
}

// 4 колонки
func makeDashboardDP() *termui.DataProviderTable {
	dpt := termui.NewDataProviderTable()
	dpt.AddData(
		makeDashboardItem(0x33da),
		0,
		0,
		0)
	dpt.AddData(
		makeDashboardItem(0x269d),
		0,
		1,
		0)
	dpt.AddData(
		makeDashboardItem(0x20ad),
		0,
		2,
		0)
	dpt.AddData(
		makeDashboardItem(0x2180),
		0,
		3,
		0)

	dpt.AddData(makeDashboardItem(0x2581), 1, 0, 0)
	dpt.AddData(makeDashboardItem(0x216d), 1, 1, 0)
	dpt.AddData(makeDashboardItem(0x2592), 1, 2, 0)
	dpt.AddData(makeDashboardItem(0x33ff), 1, 3, 0)

	dpt.Caching()
	return dpt
}

func makeDashboardItem(item int32) string {
	// 0x23b0, 0x23b1 - красивые скобки
	// items := []int32 {0x0020, item, 0x0020}
	var style termui.Style
	if item%2 > 0 {
		style = termui.NewStyle(termui.ColorYellow)
	} else {
		style = termui.NewStyle(termui.ColorGreenBgDunkel)
	}

	return termui.FormatStrWithStyle(string(item), style)
}
