package main

import (
	"fmt"
	"github.com/kallaurru/termui/v3"
	"log"
)

func main() {
	var input int32 = 1209039
	var amount, dec int64 = -130000, 78
	fmtStr := termui.FormatAmountToMonoNumbers(amount, int32(dec), true)
	log.Println(fmtStr)
	fmt.Println(termui.ConvertToMonoNumbers(input))
	log.Println(termui.ConvertToFullNumbers(input))
	log.Println(termui.ConvertSymToMarkers("BAL"))
}
