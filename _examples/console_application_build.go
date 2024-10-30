package main

import (
	ui "github.com/kallaurru/termui/v3"
	tmpl "github.com/kallaurru/termui/v3/tmpl"
	"github.com/kallaurru/termui/v3/widgets"
	"log"
	"math"
	"time"
)

func main() {

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	// создаем виджеты

	sinFloat64 := (func() []float64 {
		n := 400
		data := make([]float64, n)
		for i := range data {
			data[i] = 1 + math.Sin(float64(i)/5)
		}
		return data
	})()

	sl := widgets.NewSparkline()
	sl.Data = sinFloat64[:100]
	sl.LineColor = ui.ColorCyan
	sl.TitleStyle.Fg = ui.ColorWhite

	slg := widgets.NewSparklineGroup(sl)
	slg.Title = "Sparkline"

	lc := widgets.NewPlot()
	lc.Title = "braille-mode Line Chart"
	lc.Data = append(lc.Data, sinFloat64)
	lc.AxesColor = ui.ColorWhite
	lc.LineColors[0] = ui.ColorYellow

	gs := buildGsSlice()

	ls := widgets.NewList()
	ls.Rows = []string{
		"[1] Downloading File 1",
		"",
		"",
		"",
		"[2] Downloading File 2",
		"",
		"",
		"",
		"[3] Uploading File 3",
	}
	ls.Border = false

	p := widgets.NewParagraph()
	p.Text = "<> This row has 3 columns\n<- Widgets can be stacked up like left side\n<- Stacked widgets are treated as a single widget"
	p.Title = "Demonstration"

	grid := cAppBuildGrid(slg, lc, ls, p)

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
			if tickerCount == 100 {
				return
			}
			for _, g := range gs {
				g.Percent = (g.Percent + 3) % 100
			}
			slg.Sparklines[0].Data = sinFloat64[tickerCount : tickerCount+100]
			lc.Data[0] = sinFloat64[2*tickerCount:]
			ui.Render(grid)
			tickerCount++
		}
	}
}

func cAppBuildGrid(slg *widgets.SparklineGroup, lc *widgets.Plot, ls *widgets.List, p *widgets.Paragraph) *ui.Grid {
	root, ok := tmpl.NewAppGridSchema(true,
		true,
		ui.NewAdaptiveSizeTwoPercentile(),
		ui.NewAdaptiveSizeTwoPercentile())

	if !ok {
		panic(interface{}("not created schema"))
	}
	gs := buildGsSlice()
	row0 := buildRow0Block(slg, lc)
	row1 := buildRow1Block(gs, ls, p)

	root.AddItem(row0)
	root.AddItem(row1)
	maxX, maxY := ui.TerminalDimensions()
	grid, ok := root.Grid(0, 0, maxX, maxY)
	if !ok {
		panic(interface{}("not built grid"))
	}
	return grid
}

func buildRow0Block(slg *widgets.SparklineGroup, lc *widgets.Plot) *tmpl.AppGridSchema {
	// создаем виджеты
	schema, ok := tmpl.NewAppGridSchema(false,
		false,
		ui.NewAdaptiveSizeTwoPercentile(),
		ui.NewAdaptiveSizeTwoPercentile())
	if !ok {
		panic(interface{}("not created schema"))
	}

	schema.AddItem(slg)
	schema.AddItem(lc)
	schema.SetDeep(0)

	return schema
}

func buildRow1Block(gs []*widgets.Gauge, ls *widgets.List, p *widgets.Paragraph) *tmpl.AppGridSchema {
	schema, ok := tmpl.NewAppGridSchema(false,
		false,
		ui.NewAdaptiveSizeFirstPercentile(),
		ui.NewAdaptiveSizeFirstPercentile(),
		ui.NewAdaptiveSizeTwoPercentile())

	schema.SetDeep(0)

	if !ok {
		panic(interface{}("not created schema"))
	}

	schema.AddItem(ls)
	midRowSchema, ok := tmpl.NewAppGridSchema(false,
		true,
		ui.NewAdaptiveSize(30),
		ui.NewAdaptiveSize(30),
		ui.NewAdaptiveSize(40))

	midRowSchema.SetDeep(1)
	if !ok {
		panic(interface{}("not created schema"))
	}
	for _, gsItem := range gs {
		midRowSchema.AddItem(gsItem)
	}
	schema.AddItem(midRowSchema)
	schema.AddItem(p)

	return schema
}

func buildGsSlice() []*widgets.Gauge {
	gs := make([]*widgets.Gauge, 3)
	for i := range gs {
		gs[i] = widgets.NewGauge()
		gs[i].Percent = i * 10
		gs[i].BarColor = ui.ColorRed
	}
	return gs
}
