package main

import (
	"errors"
	"fmt"
	ui "github.com/kallaurru/termui/v3"
	"github.com/kallaurru/termui/v3/tmpl"
	"github.com/kallaurru/termui/v3/widgets"
	"log"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	var (
		formInFocus = true
		tickerCount = 1
		tickerMax   = 1
	)
	// 2 строки. Основной блок индикаторов и логирование
	rootSchema := tmpl.NewGridSchema(ui.NewAdaptiveSize(80), ui.NewAdaptiveSize(20))

	// индикаторы по первому счету
	rootSchema.AddCell(0, nil)
	// сумма по обоим счетам
	rootSchema.AddCell(0, nil)
	// индикаторы по второму счету
	rootSchema.AddCell(0, nil)

	rootSchema.AddCell(2, makeCellDetail(
		1,
		0,
		ui.NewAdaptiveSizeMax(),
		true,
		makeLogStack()))
}

func makeLogStack() *widgets.LogStack {
	logStack := widgets.NewLogStack(7)
	logStack.AddWarnLogRecord("This is a warning message")
	logStack.AddErrLogRecord(errors.New("this is error test message"))
	logStack.AddInfoLogRecord("This is info message")
	logStack.AddWarnLogRecord("another is a warning message")
	logStack.AddErrLogRecord(errors.New("another is error test message"))
	logStack.AddInfoLogRecord("another is info message")
	logStack.AddErrLogRecord(errors.New("this is error test message"))

	return logStack
}

func makeAccountFinanceSchema(accountId string) *tmpl.GridSchema {
	schema := tmpl.NewGridSchema(
		ui.NewAdaptiveSizeMin(),
		ui.NewAdaptiveSizeTwenty(),
		ui.NewAdaptiveSize(40),
		ui.NewAdaptiveSizeThird())

	row0col0 := makeCellDetail(
		0,
		0,
		ui.NewAdaptiveSizeMax(),
		true,
		makeGauge())
	row1col0 := makeCellDetail(
		1,
		0,
		ui.NewAdaptiveSizeTwoPercentile(),
		true,
		makeAccountFinanceWidget(accountId))
	row1col1 := makeCellDetail(
		1,
		1,
		ui.NewAdaptiveSizeMax(),
		true,
		makeGauge())

	schema.AddCell(0, row0col0)
	schema.AddCell(1, row1col0)
	schema.AddCell(1, row1col1)

	return schema
}

func makeFinanceStatSchema() *tmpl.GridSchema {

}

func makeAccountFinanceWidget(id string) *widgets.ATable {
	titlePref := ui.ConvertSymToMarkers("A")
	title := fmt.Sprintf("%s %s", titlePref, id)

	table := widgets.NewATable()

	table.PaddingTop = 1
	table.Border = true
	table.MakeGlamourTitle(title)
	table.RowSeparator = false
	table.ColSeparator = false

	table.ColumnWidths = []int{2, 12, -1}

	table.AddColAlignment(0, ui.AlignRight)
	table.AddColAlignment(1, ui.AlignLeft)
	table.AddColAlignment(2, ui.AlignRight)

	rowMaker := func() [][]string {
		balPref := ui.ConvertSymToMarkers("B")
		balSuf := ui.ConvertToMonoNumbers(450000)
		balLines := []string{
			fmt.Sprintf("[%s ](fg:white,bg:22,mod:bold)", balPref),
			fmt.Sprintf("[%s %s](fg:white,bg:22,mod:bold)", string(ui.RUR), balSuf),
			"баланс",
		}

		dayPref := ui.ConvertSymToMarkers("D")
		daySuf := ui.ConvertToMonoNumbers(8787)
		dayLines := []string{
			fmt.Sprintf("[%s ](fg:30,mod:bold)", dayPref),
			fmt.Sprintf("[%s %s](fg:30,mod:bold)", string(ui.RUR), daySuf),
			"за день",
		}
		monPref := ui.ConvertSymToMarkers("M")
		monSuf := ui.ConvertToMonoNumbers(9000)
		monLines := []string{
			fmt.Sprintf("[%s ](fg:white,mod:bold)", monPref),
			fmt.Sprintf("[%s %s](fg:green,mod:bold)", string(ui.RUR), monSuf),
			"за месяц",
		}
		yPref := ui.ConvertSymToMarkers("Y")
		ySuf := ui.ConvertToMonoNumbers(9999999)
		yLines := []string{
			fmt.Sprintf("[%s ](fg:white,mod:bold)", yPref),
			fmt.Sprintf("[%s %s](fg:green,mod:bold)", string(ui.RUR), ySuf),
			"за год",
		}

		return [][]string{
			balLines,
			dayLines,
			monLines,
			yLines,
		}
	}
	table.Rows = rowMaker()

	return table
}

func makeLossLimitAccountWidget(accountId string) *widgets.ATable {
	titlePref := ui.ConvertSymToMarkers("LL")
	title := fmt.Sprintf("%s %s", titlePref, accountId)

	table := widgets.NewATable()
	table.PaddingTop = 1
	table.Border = true
	table.MakeGlamourTitle(title)

	table.RowSeparator = false
	table.ColSeparator = false

	table.ColumnWidths = []int{8, 4, -1}

	table.AddColAlignment(0, ui.AlignRight)
	table.AddColAlignment(1, ui.AlignLeft)
	table.AddColAlignment(2, ui.AlignRight)

	rowMaker := func() [][]string {
		// поступления из вне в текущем месяце. Берется определенный процент
		extInAmount := fmt.Sprintf(
			"[%s %s](fg:30,mod:bold)",
			ui.ConvertToMonoNumbers(700),
			string(ui.RUR))
		extInSymb := fmt.Sprintf(
			"[%s](fg:30,mod:bold)",
			ui.ConvertSymToMarkers("EI"),
		)
		extInText := fmt.Sprintf(
			"[%s](fg:white,mod:bold)",
			"% от поступлений на счет в текущем месяце",
		)

		ldAmount := fmt.Sprintf(
			"[%s %s](fg:30,mod:bold)",
			ui.ConvertToMonoNumbers(300),
			string(ui.RUR))
		ldInSym := fmt.Sprintf(
			"[%s](fg:30,mod:bold)",
			ui.ConvertSymToMarkers("LD"),
		)
		ldInText := fmt.Sprintf(
			"[%s](fg:white,mod:bold)",
			"% от прибыли за последний день",
		)

		lpdAmount := fmt.Sprintf(
			"[%s %s](fg:30,mod:bold)",
			ui.ConvertToMonoNumbers(1000),
			string(ui.RUR))

		lpdInSym := fmt.Sprintf(
			"[%s](fg:30,mod:bold)",
			ui.ConvertSymToMarkers("DL"))

		lpdInText := fmt.Sprintf(
			"[%s](fg:white,mod:bold)",
			"стоп-лосс лимит на день")

		firstRow := []string{extInAmount, extInSymb, extInText}
		secondRow := []string{ldAmount, ldInSym, ldInText}
		thirdRow := []string{lpdAmount, lpdInSym, lpdInText}

		return [][]string{
			firstRow,
			secondRow,
			thirdRow,
		}
	}

	table.Rows = rowMaker()
	table.Rows = [][]string{
		{"[Limits:](fg:white,mod:bold)", "[Loss limits](fg:white,mod:bold)"},
		{"[-> D:](fg:white,mod:bold)", "[400](fg:red,mod:bold)"},
		{"[Ext In](fg:white,mod:bold)", "[9000](fg:green,mod:bold)"},
		{"[Last Day %](fg:white,mod:bold)", "[200](fg:green,mod:bold)"},
	}

	return table
}
