package main

import (
	"errors"
	"github.com/kallaurru/termui/v3"
	"github.com/kallaurru/termui/v3/widgets"
	"log"
)

func main() {
	mw := make(map[string]interface{})
	mdp := make(map[string]interface{})

	name := "fin_result_acc_1"

	w := widgets.NewATable()
	w.SetName(name)
	dpt := lwdpMakeDataProviderTable()
	// линкуем виджет и дата провайдер
	w.Rows = dpt.GetTable()

	mw[name] = w
	mdp[name] = dpt

	newAmount := int32(377373)

	name = w.GetName()
	dp := mdp[name]
	err := handleUpdateDayAmount(dp, newAmount)
	if err != nil {
		log.Fatalln(err)
	}
}

func handleUpdateDayAmount(dp interface{}, amount int32) error {
	// w - ATable, dp = dpTable

	datap, ok := dp.(*termui.DataProviderTable)
	if !ok {
		return errors.New("no correct dp type")
	}
	amountStr := termui.ConvertToMonoNumbers(amount)
	datap.UpdateData(amountStr, 0, 1, 1)
	return nil
}

func lwdpMakeDataProviderTable() *termui.DataProviderTable {

	dpt := termui.NewDataProviderTable()

	dpt.AddData(termui.ConvertSymToSquaredMarker("D"), 0, 0, 0)

	dpt.AddData(string(termui.RUR), 0, 1, 0)
	dpt.AddData(termui.ConvertToMonoNumbers(182811), 0, 1, 1)

	dpt.AddData("день", 0, 2, 0)

	// вторая строка
	dpt.AddData(termui.ConvertSymToSquaredMarker("M"), 1, 0, 0)

	dpt.AddData(string(termui.RUR), 1, 1, 0)
	dpt.AddData(termui.ConvertToMonoNumbers(777777), 1, 1, 1)

	dpt.AddData("месяц", 1, 2, 0)

	dpt.Caching()

	return dpt
}
