package termui

import (
	"fmt"
)

// FormatAmount - вывод форматированной суммы.
// @decimal - если меньше 0, значит вывод десятичных отключаем
func FormatAmount(amount, decimal int32, isMono bool, addCurSym bool) string {
	var amountStr, decimalStr string
	useDecimal := decimal >= 0

	if decimal < 0 {
		decimal = 0
	}
	if useDecimal {
		if decimal < 10 {
			if isMono {
				decimalStr = fmt.Sprintf("%s%s", ConvertToMonoNumbers(0), ConvertToMonoNumbers(decimal))
			} else {
				decimalStr = fmt.Sprintf("%s%s", ConvertToFullNumbers(0), ConvertToFullNumbers(decimal))
			}
		} else {
			if isMono {
				decimalStr = ConvertToMonoNumbers(decimal)
			} else {
				decimalStr = ConvertToFullNumbers(decimal)
			}
		}
	}

	if isMono {
		amountStr = ConvertToMonoNumbers(amount)
	} else {
		amountStr = ConvertToFullNumbers(amount)
	}

	if addCurSym {
		return fmt.Sprintf("%s%s,%s", string(RUR), amountStr, decimalStr)
	}
	if useDecimal {
		return fmt.Sprintf("%s,%s", amountStr, decimalStr)
	} else {
		return amountStr
	}
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
