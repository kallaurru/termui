package termui

import (
	"fmt"
)

func FormatAmount(amount int32, isMono bool, addCurSym bool) string {
	var amountStr string
	if isMono {
		amountStr = ConvertToMonoNumbers(amount)
	} else {
		amountStr = ConvertToFullNumbers(amount)
	}

	if addCurSym {
		return fmt.Sprintf("%s%s", string(RUR), amountStr)
	}
	return amountStr
}

func FormatStrAsMarkers(in string, asSquared bool) string {
	if asSquared {
		return ConvertSymToSquaredMarker(in)
	} else {
		return ConvertSymToMarkers(in)
	}
}

func FormatStrWithStyle(in string, style Style) string {

	styleStr := StyleToString(style)
	if len(styleStr) > 0 {
		return fmt.Sprintf(
			"%s%s%s%s",
			string(TokenBeginStyledText),
			in,
			string(TokenEndStyledText),
			styleStr)
	}

	return fmt.Sprintf(
		"%s%s%s%s",
		string(TokenBeginStyledText),
		in,
		string(TokenEndStyledText),
		styleStr)
}
