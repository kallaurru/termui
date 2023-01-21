package main

import (
	"errors"
	"fmt"
	ui "github.com/kallaurru/termui/v3"
	"github.com/kallaurru/termui/v3/tmpl"
	"github.com/kallaurru/termui/v3/widgets"
	"image"
	"log"
	"time"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	var focus, isReal bool
	isReal = true
	focus = true

	logStack := widgets.NewLogStack(5)
	logStack.AddWarnLogRecord("This is a warning message")
	logStack.AddErrLogRecord(errors.New("this is error test message"))
	logStack.AddInfoLogRecord("This is info message")

	logbookSchema := tmpl.NewGridSchema(ui.NewAdaptiveSize(80), ui.NewAdaptiveSize(20))
	// row 0
	g := makeGauge()
	logbookSchema.AddCell(0, makeCellDetail(0, 0, ui.NewAdaptiveSize(30), false, makeSchemaCellWithGauge(g)))
	logbookSchema.AddCell(0, makeCellDetail(0, 1, ui.NewAdaptiveSize(40), false, makeCentralZoneSchema()))
	logbookSchema.AddCell(0, makeCellDetail(0, 2, ui.NewAdaptiveSize(30), false, makeSchemaCellWithGauge(g)))
	// row 1
	logbookSchema.AddCell(1, makeCellDetail(1, 0, ui.NewAdaptiveSizeMax(), true, logStack))
	x, y := 0, 0
	if isReal {
		x, y = ui.TerminalDimensions()
	} else {
		x, y = 80, 120
	}
	grid := logbookSchema.BuildGrid(x, y)
	grid.Border = true
	ui.Render(grid)

	tickerCount := 1
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	tickerMax := 30
	helper := makeHlpPopup(grid.Inner)
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<F1>":
				ui.Render(helper)
				focus = false
			case "<Backspace>":
				if !focus {
					ui.Clear()
					ui.Render(grid)
				}
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		case <-ticker:
			if tickerCount == tickerMax {
				return
			}
			tickerCount++
			res := float64(tickerCount) / float64(tickerMax) * 100
			g.Percent = int(res)
			ui.Render(g)
		}
	}
}

func makeSchemaCell() *tmpl.GridSchema {
	gauge := widgets.NewGauge()
	gauge.Percent = 10
	gauge.BarColor = ui.ColorBlue
	gauge.Title = "Loaded months"
	gauge.BorderStyle.Fg = ui.ColorWhite
	gauge.TitleStyle.Fg = ui.ColorCyan
	gauge.Label = "Two month loaded"
	gauge.LabelStyle.Fg = ui.ColorGreen

	ls := widgets.NewList()
	ls.Border = false
	ls.Rows = []string{
		"[1] Downloading File 1",
		"",
		"[2] Downloading File 2",
		"",
		"[3] Uploading File 3",
	}

	schema := tmpl.NewGridSchema(ui.NewAdaptiveSizeMin(), ui.NewAdaptiveSize(90))
	firstRow := makeCellDetail(0, 0, ui.NewAdaptiveSizeMax(), true, gauge)
	secondRow := makeCellDetail(1, 0, ui.NewAdaptiveSizeMax(), true, ls)

	schema.AddCell(0, firstRow)
	schema.AddCell(1, secondRow)

	return schema
}

func makeSchemaCellWithGauge(g *widgets.Gauge) *tmpl.GridSchema {
	fin := makeFinanceWidget()
	ll := makeLossLimitWidget()
	pb := makeBudgetPosi()

	schema := tmpl.NewGridSchema(
		ui.NewAdaptiveSizeMin(),
		ui.NewAdaptiveSizeTwenty(),
		ui.NewAdaptiveSize(40),
		ui.NewAdaptiveSizeThird())

	firstRow := makeCellDetail(0, 0, ui.NewAdaptiveSizeMax(), true, g)
	secondRow := makeCellDetail(1, 0, ui.NewAdaptiveSizeMax(), true, fin)
	row3 := makeCellDetail(2, 0, ui.NewAdaptiveSizeMax(), true, ll)
	row4 := makeCellDetail(3, 0, ui.NewAdaptiveSizeMax(), true, pb)

	schema.AddCell(0, firstRow)
	schema.AddCell(1, secondRow)
	schema.AddCell(2, row4)
	schema.AddCell(3, row3)

	return schema
}

func makeCentralZoneSchema() *tmpl.GridSchema {
	schema := tmpl.NewGridSchema(
		ui.NewAdaptiveSizeThird(),
		ui.NewAdaptiveSize(40),
		ui.NewAdaptiveSizeThird())

	first := widgets.NewBarChart()
	first.Labels = []string{"Op", "In", "Di"}
	first.Data = []float64{100000, 400000, 5000}
	first.MakeGlamourTitle("Incoming")
	first.BarWidth = 5
	first.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	first.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	first.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	th := widgets.NewBarChart()
	th.Labels = []string{"Bf", "Tx", "Sh"} // bf - комиссия за оп, tx - налоги, sh - комиссия за плечи
	th.Data = []float64{1030, 30000, 1000}
	th.MakeGlamourTitle("Expenses")
	th.BarWidth = 5
	th.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	th.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	th.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	second := makeLossLimitWidget()

	schema.AddCell(0, makeCellDetail(0, 0, ui.NewAdaptiveSizeMax(), true, first))
	schema.AddCell(1, makeCellDetail(1, 0, ui.NewAdaptiveSizeMax(), true, second))
	schema.AddCell(2, makeCellDetail(2, 0, ui.NewAdaptiveSizeMax(), true, th))

	return schema
}

func makeList() *widgets.List {
	ls2 := widgets.NewList()
	ls2.Border = false
	ls2.Rows = []string{
		"[1] Downloading File 7",
		"",
		"[2] Downloading File 8",
		"",
		"[3] Uploading File 9",
	}

	return ls2
}

func makeCellDetail(row, col uint8, size ui.AdaptiveSize, isWidget bool, data interface{}) *tmpl.CellDetail {
	cd := tmpl.NewCellDetail(row, col, size)
	if isWidget {
		w, ok := data.(ui.Drawable)
		if ok {
			cd.SetWidget(w)
			return cd
		}
	}
	w, ok := data.(*tmpl.GridSchema)
	if ok {
		cd.SetSchema(w)
		return cd
	}

	return nil
}

func makeGauge() *widgets.Gauge {
	gauge := widgets.NewGauge()
	gauge.Percent = 10
	gauge.BarColor = ui.ColorCyan
	gauge.Title = "Loaded months"
	gauge.BorderStyle.Fg = ui.ColorWhite
	gauge.TitleStyle.Fg = ui.ColorWhite
	gauge.Label = "Two month loaded"
	gauge.LabelStyle.Fg = ui.ColorGreen

	return gauge
}

func makeHlpPopup(parent image.Rectangle) *widgets.ATable {
	var w, h int
	var fHlpData widgets.FnHlpRowMarker

	fHlpData = func(w, h *int) [][]string {
		*w = 45
		*h = 6

		return [][]string{
			{"[F1](fg:cyan,mod:bold)", "Show this window"},
			{"[F3](fg:cyan,mod:bold)", "Show advanced data"},
			{"[F4](fg:cyan,mod:bold)", "Edit data"},
			{"[F12](fg:cyan,mod:bold)", "Close application | Close this window"},
		}
	}

	table := widgets.NewATable()
	table.BorderTop = true
	table.RowSeparator = false
	table.ColSeparator = false

	table.ColumnWidths = []int{4, -1}

	table.AddColAlignmentHelperTheme()
	table.Rows = fHlpData(&w, &h)

	dim := ui.MakeCenterPositionWidget(parent, w, h)
	table.SetRect(dim.Min.X, dim.Min.Y, dim.Max.X, dim.Max.Y)

	return table
}

func makeFinanceWidget() *widgets.ATable {
	titlePref := ui.ConvertSymToMarkers("A")
	table := widgets.NewATable()
	table.PaddingTop = 1
	table.Border = true
	table.MakeGlamourTitle(fmt.Sprintf("%s 1782637621", titlePref))
	table.RowSeparator = false
	table.ColSeparator = false

	table.ColumnWidths = []int{2, 12, -1}

	table.AddColAlignment(0, ui.AlignRight)
	table.AddColAlignment(1, ui.AlignLeft)
	table.AddColAlignment(2, ui.AlignRight)
	rowMaker := func() [][]string {
		balPref := ui.ConvertSymToMarkers("B")
		balSuf := ui.ConvertToMonoNumbers(450000)
		balLines := []string{
			fmt.Sprintf("[%s ](fg:white,bg:22,mod:bold)", balPref),
			fmt.Sprintf("[%s %s](fg:white,bg:22,mod:bold)", string(ui.RUR), balSuf),
			"баланс",
		}

		dayPref := ui.ConvertSymToMarkers("D")
		daySuf := ui.ConvertToMonoNumbers(8787)
		dayLines := []string{
			fmt.Sprintf("[%s ](fg:30,mod:bold)", dayPref),
			fmt.Sprintf("[%s %s](fg:30,mod:bold)", string(ui.RUR), daySuf),
			"за день",
		}
		monPref := ui.ConvertSymToMarkers("M")
		monSuf := ui.ConvertToMonoNumbers(9000)
		monLines := []string{
			fmt.Sprintf("[%s ](fg:white,mod:bold)", monPref),
			fmt.Sprintf("[%s %s](fg:green,mod:bold)", string(ui.RUR), monSuf),
			"за месяц",
		}
		yPref := ui.ConvertSymToMarkers("Y")
		ySuf := ui.ConvertToMonoNumbers(9999999)
		yLines := []string{
			fmt.Sprintf("[%s ](fg:white,mod:bold)", yPref),
			fmt.Sprintf("[%s %s](fg:green,mod:bold)", string(ui.RUR), ySuf),
			"за год",
		}

		return [][]string{
			balLines,
			dayLines,
			monLines,
			yLines,
		}
	}
	table.Rows = rowMaker()

	return table
}

func makeLossLimitWidget() *widgets.ATable {
	table := widgets.NewATable()
	table.PaddingTop = 1
	table.Border = true
	table.MakeGlamourTitle("Stop-Loss Limits")

	table.RowSeparator = false
	table.ColSeparator = false

	table.ColumnWidths = []int{12, -1}

	table.AddColAlignmentHelperTheme()
	table.Rows = [][]string{
		{"[Limits:](fg:white,mod:bold)", "[Loss limits](fg:white,mod:bold)"},
		{"[-> D:](fg:white,mod:bold)", "[400](fg:red,mod:bold)"},
		{"[Ext In](fg:white,mod:bold)", "[9000](fg:green,mod:bold)"},
		{"[Last Day %](fg:white,mod:bold)", "[200](fg:green,mod:bold)"},
	}

	return table
}

func makeBudgetPosi() *widgets.ATable {
	table := widgets.NewATable()
	table.PaddingTop = 1
	table.Border = true
	table.MakeGlamourTitle("Budgets. Total - 39283298")
	table.RowSeparator = false
	table.ColSeparator = false

	table.ColumnWidths = []int{7, 10, -1}

	table.AddColAlignment(0, ui.AlignCenter)
	table.AddColAlignment(1, ui.AlignCenter)
	table.AddColAlignment(2, ui.AlignCenter)
	table.Rows = [][]string{
		{"[SBER](fg:white,mod:bold)", "[75 000](fg:white,mod:bold)", "[0](fg:white,mod:bold)"},
		{"[POSI](fg:white,mod:bold)", "[75 000](fg:red,mod:bold)", "[2](fg:green,mod:bold)"},
		{"[TCSG](fg:white,mod:bold)", "[75 000](fg:green,mod:bold)", "[0](fg:white,mod:bold)"},
		{"[PLZL](fg:white,mod:bold)", "[75 000](fg:green,mod:bold)", "[0](fg:white,mod:bold)"},
	}

	return table
}
