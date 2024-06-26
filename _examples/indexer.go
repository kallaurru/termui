package main

import (
	. "github.com/kallaurru/termui/v3"
	"github.com/kallaurru/termui/v3/widgets"
	"log"
)

func main() {
	makeIndexerTable()
	//makeIndexerList()
}

func makeIndexerTable() {
	dpt := NewDataProviderTable()

	dpt.AddData(ConvertSymToSquaredMarker("D"), 0, 0, 0)

	dpt.AddData(string(RUR), 0, 1, 0)
	dpt.AddData(ConvertToMonoNumbers(182811), 0, 1, 1)

	dpt.AddData("день", 0, 2, 0)

	// вторая строка
	dpt.AddData(ConvertSymToSquaredMarker("M"), 1, 0, 0)

	dpt.AddData(string(RUR), 1, 1, 0)
	dpt.AddData(ConvertToMonoNumbers(777777), 1, 1, 1)

	dpt.AddData("месяц", 1, 2, 0)

	dpt.Caching()
	aTabl := widgets.NewATable()
	aTabl.Rows = dpt.GetTable()

	dpt.UpdateData(ConvertToMonoNumbers(999999), 1, 1, 1)

	log.Println("indexer is ready", len(aTabl.Rows))
}

func makeIndexerList() {
	dpl := NewDataProviderList()
	dpl.AddData(ConvertSymToSquaredMarker("D"), 0, 0)

	dpl.AddData(string(RUR), 1, 0)
	dpl.AddData(ConvertToMonoNumbers(182811), 2, 0)

	dpl.AddData("день", 3, 0)
	dpl.Caching()
	wl := widgets.NewListES()
	wl.Rows = dpl.GetListPtr()
	dpl.UpdateData("год", 0, 3)

	log.Println("indexer is ready", len(*wl.Rows))

}

func makeIndexerText() {

}
