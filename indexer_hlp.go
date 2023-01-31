package termui

import (
	"fmt"
	"strings"
)

func MakeDataProviderAddress(param, row, col uint32) uint32 {
	var address uint32 = 0
	if param+row+col == 0 {
		return address
	}
	if param > 0 {
		address |= param << 16
	}
	if row > 0 {
		address |= row << 8
	}

	address |= col

	return address
}

func MakeStr(countParams uint8, data []interface{}, useSplit32 bool) string {
	var formatStr string

	if useSplit32 {
		formatStr = strings.Repeat("%s ", int(countParams))
		formatStr = strings.TrimRight(formatStr, " ")
	} else {
		formatStr = strings.Repeat("%s", int(countParams))
	}

	return fmt.Sprintf(formatStr, data...)
}
