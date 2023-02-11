package main

import (
	"fmt"
	"github.com/kallaurru/termui/v3"
)

func main() {
	dpl := makeDPLSize()
	l := dpl.GetList()

	for _, el := range l {
		fmt.Printf("%s\n", el)
	}
}

func makeDPLSize() *termui.DataProviderList {
	dpl := termui.NewDataProviderList()
	dpl.UseSplit32().SetWidthParams([]uint8{5, 22, 7})

	dpl.AddData(termui.ConvertSymToSquaredMarker("POSI"), 0, 0)
	dpl.AddData(
		termui.FormatAmount(3883, 32, true, true),
		1,
		0)
	dpl.AddData("day dat", 2, 0)

	dpl.AddData(termui.ConvertSymToSquaredMarker("WUSH"), 0, 1)
	dpl.AddData(
		termui.FormatAmount(3873783, 32, true, true),
		1,
		1)
	dpl.AddData("ball", 2, 1)

	dpl.Caching()

	return dpl
}

func makeDPLSize2() *termui.DataProviderList {
	dpl := termui.NewDataProviderList()
	dpl.UseSplit32().SetWidthParams([]uint8{5, 22, 7})

	dpl.AddData("DAY", 0, 0)
	dpl.AddData(
		termui.FormatAmount(3883, 32, true, true),
		1,
		0)
	dpl.AddData("day dat", 2, 0)

	dpl.AddData("MONTH", 0, 1)
	dpl.AddData(
		termui.FormatAmount(3873783, 32, true, true),
		1,
		1)
	dpl.AddData("buld", 2, 1)

	dpl.Caching()

	return dpl
}
