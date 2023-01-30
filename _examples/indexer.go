package main

import (
	. "github.com/kallaurru/termui/v3"
	"log"
)

func main() {
	makeIndexerTable()
}

func makeIndexerTable() {
	dpt := NewDataProviderTable()

	dpt.AddData(ConvertSymToSquaredMarker("D"), 0, 0, 0)
	dpt.AddData(string(RUR), 0, 0, 1)
	dpt.AddData(ConvertToMonoNumbers(182811), 1, 0, 1)

	dpt.AddData("день", 0, 0, 2)

	// вторая строка
	dpt.AddData(ConvertSymToSquaredMarker("M"), 0, 1, 0)
	dpt.AddData(string(RUR), 0, 1, 1)
	dpt.AddData(ConvertToMonoNumbers(777777), 1, 1, 1)

	dpt.AddData("месяц", 0, 1, 2)

	tabl := dpt.GetTable()
	log.Println("indexer is ready", len(tabl))
}

func makeIndexerList() {

}

func makeIndexerText() {

}
