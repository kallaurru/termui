package main

import (
	ui "github.com/kallaur/termui/v3"
	"github.com/kallaur/termui/v3/tmpl"
	"log"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	standartLogBookTmpl := tmpl.NewGridTmpl(2, ui.NewAdaptiveSize(80), ui.NewAdaptiveSizeTwenty())

}

func emptyWidget() ui.Drawable {

}
