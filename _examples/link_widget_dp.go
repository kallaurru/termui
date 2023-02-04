package main

import (
	"errors"
	"fmt"
	"github.com/kallaurru/termui/v3"
	"github.com/kallaurru/termui/v3/widgets"
	"log"
)

func main() {
	// formatter test
	amount := "17777777"
	amount2 := "83832"

	// example format length string
	fmt.Printf(
		"%5s %14s %5s -%s\n",
		"POSI",
		amount,
		"100",
		"21")
	fmt.Printf(
		"%5s %14s %5s -%s\n",
		"OZON",
		amount2,
		"11",
		"1")
}

func mainOld2() {
	st := termui.NewStyleBgFree(termui.ColorWhite, termui.ModifierBold)
	in := "Ljdskjfksjl slkdjfskl"

	fstr := termui.FormatStrWithStyle(in, st)

	fmt.Printf("%s", fstr)
	amount := int32(398729)
	decimal := int32(9)

	sumStr := termui.FormatAmount(amount, decimal, true, true)
	sumStr = termui.FormatAmount(amount, decimal, false, false)
	decimal2 := int32(90)
	sumStr = termui.FormatAmount(amount, decimal2, true, false)
	fmt.Println(sumStr)
}

func mainOld() {
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
