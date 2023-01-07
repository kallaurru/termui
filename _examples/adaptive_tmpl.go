package main

import (
	ui "github.com/kallaurru/termui/v3"
	"github.com/kallaurru/termui/v3/tmpl"
	"github.com/kallaurru/termui/v3/widgets"
	"log"
	"time"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	isReal := false

	logStack := widgets.NewLogStack(5)

	logbookSchema := tmpl.NewGridSchema(ui.NewAdaptiveSize(80), ui.NewAdaptiveSizeTwenty())
	// row 0
	logbookSchema.AddCell(0, makeCellDetail(0, 0, ui.NewAdaptiveSize(30), false, makeSchemaCell()))
	logbookSchema.AddCell(0, makeCellDetail(0, 1, ui.NewAdaptiveSize(40), true, makeList()))
	logbookSchema.AddCell(0, makeCellDetail(0, 2, ui.NewAdaptiveSize(30), false, makeSchemaCell()))
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

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		case <-ticker:
			if tickerCount == 30 {
				return
			}
			tickerCount++
		}
	}
}

func makeSchemaCell() *tmpl.GridSchema {
	gauge := widgets.NewGauge()
	gauge.Percent = 10
	gauge.BarColor = ui.ColorBlue

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
