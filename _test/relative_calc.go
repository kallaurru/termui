package main

import (
	. "github.com/kallaur/termui/v3"
	"log"
)

func main() {
	var (
		rows uint8 = 5
	)

	l := CalcRelativeHeight(
		rows,
		NewAdaptiveSizeMin(),
		NewAdaptiveSizeTwoPercentile())

	if len(l) != int(rows) {
		log.Println("No equals count")
	}
}
