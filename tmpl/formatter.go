package tmpl

import (
	"fmt"
	"github.com/kallaurru/termui/v3"
)

func FormatAmount(amount int32, isMono bool, addCurSym bool) string {
	var amountStr string
	if isMono {
		amountStr = termui.ConvertToMonoNumbers(amount)
	} else {
		amountStr = termui.ConvertToFullNumbers(amount)
	}

	if addCurSym {
		return fmt.Sprintf("%s%s", string(termui.RUR), amountStr)
	}
	return amountStr
}

func FormatStrAsMarkers(in string, asSquared bool) string {
	if asSquared {
		return termui.ConvertSymToSquaredMarker(in)
	} else {
		return termui.ConvertSymToMarkers(in)
	}
}

func FormatStrWithStyle(in string, style termui.Style) string {
	styleStr := termui.StyleToString(style)
	if len(styleStr) > 0 {
		return fmt.Sprintf(
			"%s%s%s%s",
			string(termui.TokenBeginStyledText),
			in,
			string(termui.TokenEndStyledText),
			styleStr)
	}

	return fmt.Sprintf(
		"%s%s%s%s",
		string(termui.TokenBeginStyledText),
		in,
		string(termui.TokenEndStyledText),
		styleStr)
}
