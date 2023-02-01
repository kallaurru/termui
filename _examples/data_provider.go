package main

import (
	"fmt"
	. "github.com/kallaurru/termui/v3"
	"github.com/kallaurru/termui/v3/widgets"
)

func main() {
	// пример изменения данных по ссылке
	var source *[]string

	source = new([]string)

	*source = append(*source, "lskdfjklksdkfj sdkfjs")
	dpFree(source)
	fmt.Printf("lksdkjf ksldjflk")
}

func dpFree(source *[]string) {
	(*source)[0] = "изменил источник"
}

func mainDPT() {
	// rows - 2, cols - 3
	dpt := makeDPT()
	at := makeATableWidget()
	// linking
	at.Block.MakeGlamourTitle("Data from Data Provider")
	at.Rows = dpt.GetTable()
	dpt.UpdateData("почти", 1, 2, 0)
	fmt.Printf("updated widget. Col 2 - %s %s %s", at.Rows[1][0], at.Rows[1][1], at.Rows[1][2])
}

func mainDPL() {
	dpl := makeDPL()
	l := widgets.NewList()
	dpl.SetExtTarget(&l.Rows)
	dpl.Caching()
	dpl.UpdateData("ночь", 3, 0)
	fmt.Printf("all is end")
}

func makeDPT() *DataProviderTable {

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

	return dpt
}

func makeDPL() *DataProviderList {
	dpl := NewDataProviderList()
	dpl.AddData(ConvertSymToSquaredMarker("D"), 0, 0)

	dpl.AddData(string(RUR), 1, 0)
	dpl.AddData(ConvertToMonoNumbers(182811), 2, 0)

	dpl.AddData("день", 3, 0)
	dpl.Caching()

	return dpl
}

func makeATableWidget() *widgets.ATable {
	at := widgets.NewATable()
	at.Block.MakeGlamourTitle("No data")

	return at
}
