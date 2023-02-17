package termui

import (
	"fmt"
	"strings"
)

// GetAddressElements - элементы адреса:
// - один элемент - номер параметра для текста, номер параметра для 0 строки списка, 0 строки и 0 колонки таблицы
// - два элемента - номер строки, номер параметра, 0 - номер колонки для таблицы
// - три элемента - номер строки, номер колонки, номер параметра
func GetAddressElements(address ...uint32) (p, r, c uint32) {
	switch len(address) {
	default:
		return 0, 0, 0
	case 1:
		// для текста актуален только номер параметра
		return address[0], 0, 0
	case 2:
		// для списка актуальны строка и параметр. Или таблица с одной колонкой
		return address[1], address[0], 0
	case 3:
		return address[2], address[0], address[1]
	}
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

// ParseDataProviderAddress - порядок элементов возврата как в в параметрах функции
// MakeDataProviderAddress: param, row, col
func ParseDataProviderAddress(address uint32) (uint32, uint32, uint32) {
	return (address >> 16) & 0xff, (address >> 8) & 0xff, address & 0xff
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

func MakeStrStaticW(data []interface{}, rw []uint8, useSplit32 bool) string {
	var formatStr, item string

	countParams := len(data)
	if countParams == 0 {
		return ""
	}

	lRW := len(rw)
	formatArr := make([]string, 0, 2)
	for i := 0; i < countParams; i++ {
		if i < lRW && rw[i] > 0 {
			item = fmt.Sprintf("%%%ds", rw[i])
		} else {
			item = "%s"
		}
		formatArr = append(formatArr, item)
	}
	if useSplit32 {
		formatStr = strings.Join(formatArr, " ")
	} else {
		formatStr = strings.Join(formatArr, "")
	}

	return fmt.Sprintf(formatStr, data...)
}
