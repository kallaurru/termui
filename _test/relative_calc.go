package main

import (
	"github.com/kallaur/termui/v3"
	"github.com/kallaur/termui/v3/tmpl"
	"log"
)

func main() {
	var (
		rows uint8 = 3
	)

	l, ok := termui.CalcRelativeHeight(
		rows,
		tmpl.NewAdaptiveSizeMin(),
		tmpl.NewAdaptiveSizeTwoPercentile())
	if !ok {
		log.Println("No calc relative height")
		return
	}
	if len(l) != int(rows) {
		log.Println("No equals count")
	}
}
