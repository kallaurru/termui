package main

import (
	"fmt"
	"github.com/kallaurru/termui/v3"
)

func main() {
	var input int32 = 1209039
	fmt.Println(termui.ConvertToMonoNumbers(input))
	fmt.Println(termui.ConvertToFullNumbers(input))
	fmt.Println(termui.ConvertSymToMarkers("BAL"))
}
