package termui

import (
	"fmt"
	"strings"
)

const (
	TYPE_INDEXER_TABLE_CELL    uint8 = 0x00 // address - param_number:row:col
	TYPE_INDEXER_LIST_ROW      uint8 = 0x01 // address - param_number:row
	TYPE_INDEXER_PARAM_IN_TEXT uint8 = 0x02 // address - param_number
)

func ApproveIndexerAddress(iType uint8, address ...uint32) bool {
	if iType == TYPE_INDEXER_TABLE_CELL && len(address) == 3 {
		return true
	}

	if iType == TYPE_INDEXER_LIST_ROW && len(address) == 2 {
		return true
	}

	if iType == TYPE_INDEXER_PARAM_IN_TEXT && len(address) == 1 {
		return true
	}

	return false
}

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

// MakeIndexerAddress @deprecated
func MakeIndexerAddress(iType uint8, addrElements ...uint32) uint32 {
	ok := ApproveIndexerAddress(iType, addrElements...)
	if !ok {
		return 0
	}
	switch len(addrElements) {
	case 1:
		return addrElements[0]
	case 2:
		row := addrElements[0] << 8
		col := addrElements[1]

		return row | col

	case 3:
		param := addrElements[0] << 16
		row := addrElements[1] << 8
		col := addrElements[2]

		return param | row | col
	}

	return 0
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
