package main

import (
	. "github.com/kallaurru/termui/v3"
	"log"
)

func main() {
	makeIndexerTable()
}

func makeIndexerTable() {
	i := NewCellIndexer(3)

	i.AddData(ConvertSymToSquaredMarker("D"), 0, 0, 0)
	i.AddData(string(RUR), 0, 0, 1)
	i.AddData(ConvertToMonoNumbers(182811), 1, 0, 1)

	i.AddData("день", 0, 0, 2)

	// вторая строка
	i.AddData(ConvertSymToSquaredMarker("M"), 0, 1, 0)
	i.AddData(string(RUR), 0, 1, 1)
	i.AddData(ConvertToMonoNumbers(777777), 1, 1, 1)

	i.AddData("месяц", 0, 1, 2)

	tabl := i.ToTable(2, 3)
	log.Println("indexer is ready", len(tabl))
}

func makeIndexerList() {

}

func makeIndexerText() {

}
