package termui

import (
	"fmt"
	"strings"
	"sync"
)

type DataProviderTable struct {
	idx      map[uint32]string
	cols     uint8      // количество колонок, если есть
	rows     uint8      // количество строк, есть всегда
	cache    [][]string // для текста это всегда одна строка. 0 0 для текста, 0-n 0 для списка, 0-n 0-x для таблицы
	mx       sync.RWMutex
	useSplit bool
}

func NewDataProviderTable() *DataProviderTable {
	return &DataProviderTable{
		idx:      make(map[uint32]string),
		cols:     1,
		rows:     1, // в любом случае 1 строка будет всегда
		cache:    make([][]string, 0, 2),
		mx:       sync.RWMutex{},
		useSplit: false,
	}
}

func (dpt *DataProviderTable) UseSplit32() *DataProviderTable {
	dpt.useSplit = true
	return dpt
}

func (dpt *DataProviderTable) Caching() {
	var c, p, r int
	cache := make([][]string, 0, 2)
	for r = 0; r < int(dpt.rows); r++ {
		cacheRow := make([]string, 0, 2)
		for c = 0; c < int(dpt.cols); c++ {
			cacheCol := make([]interface{}, 0, 2)
			for p = 0; p < 0xff; p++ {
				addr := MakeDataProviderAddress(uint32(p), uint32(r), uint32(c))
				val, ok := dpt.idx[addr]
				if !ok {
					break
				}
				cacheCol = append(cacheCol, val)
			}
			colStr := MakeStr(uint8(p), cacheCol, dpt.useSplit)
			cacheRow = append(cacheRow, colStr)
		}
		cache = append(cache, cacheRow)
	}

	dpt.cache = cache
}

func (dpt *DataProviderTable) GetTable() [][]string {
	return dpt.cache
}

func (dpt *DataProviderTable) GetRow(row int) []string {
	if len(dpt.cache) > row {
		return dpt.cache[row]
	}
	return nil
}

// AddStyledData - если мы хотим вставить стилизованный текст как 2 параметра:
// его нужно обернуть в спец символы
func (dpt *DataProviderTable) AddStyledData(data string, row, col, param uint32) *DataProviderTable {
	data = fmt.Sprintf("%s%s%s", string(TokenBeginStyledText), data, string(TokenEndStyledText))
	return dpt.AddData(data, row, col, param)
}

//AddData - добавить данные в провайдер
func (dpt *DataProviderTable) AddData(data string, row, col, param uint32) *DataProviderTable {
	var (
		newRows = uint8(row)
		newCols = uint8(col)
	)
	address := MakeDataProviderAddress(param, row, col)
	dpt.idx[address] = data
	if dpt.cols < newCols+1 {
		dpt.cols = newCols + 1
	}

	if dpt.rows < newRows+1 {
		dpt.rows = newRows + 1
	}

	return dpt
}

func (dpt *DataProviderTable) UpdateData(data string, address ...uint32) {
	var iParam int

	p, r, c := GetAddressElements(address...)

	// проверка индексов строки и колонки
	if !dpt.approveElementsAddress(r, c) {
		return
	}
	// проверить если существуют данные и они подготовлены под стиль, значит добавляем токены стиля
	dpt.mx.Lock()

	defer dpt.mx.Unlock()

	dpt.AddData(data, r, c, p)

	cacheCol := make([]interface{}, 0, 2)
	for iParam = 0; iParam < 0xff; iParam++ {
		addr := MakeDataProviderAddress(uint32(iParam), r, c)
		val, ok := dpt.idx[addr]
		if !ok {
			break
		}
		cacheCol = append(cacheCol, val)
	}
	colStr := MakeStr(uint8(iParam), cacheCol, dpt.useSplit)
	dpt.insertInCache(r, c, colStr)
}

func (dpt *DataProviderTable) insertInCache(row, col uint32, data string) {
	if len(dpt.cache) <= int(row) {
		return
	}
	cols := dpt.cache[row]
	if len(cols) <= int(col) {
		return
	}
	dpt.cache[row][col] = data
}

func (dpt *DataProviderTable) approveElementsAddress(r, c uint32) bool {
	if uint8(r) >= dpt.rows || uint8(c) >= dpt.cols {
		return false
	}

	return true
}

func (dpt *DataProviderTable) isOnlyStyledText(p, r, c uint32) bool {
	// если в переменной только стилизованный текст без строки стиля - true
	address := MakeDataProviderAddress(p, r, c)
	data, ok := dpt.idx[address]
	if !ok {
		return false
	}
	has := strings.HasSuffix(data, string(tokenEndStyle))
	if has {
		return false
	}

	return strings.HasPrefix(data, string(TokenBeginStyledText))
}
