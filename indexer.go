package termui

import (
	"fmt"
	"strings"
)

type Indexer struct {
	iType uint8 // ячейка в таблице. Полный адрес ячейка:строка:колонка
	idx   map[uint32]string
	cols  uint8
}

func NewCellIndexer(cols uint8) *Indexer {
	return &Indexer{
		iType: TYPE_INDEXER_TABLE_CELL,
		idx:   make(map[uint32]string),
		cols:  cols,
	}
}

func NewRowIndexer() *Indexer {
	return &Indexer{
		iType: TYPE_INDEXER_LIST_ROW,
		idx:   make(map[uint32]string),
		cols:  0,
	}
}

func NewTextParamIndexer() *Indexer {
	return &Indexer{
		iType: TYPE_INDEXER_PARAM_IN_TEXT,
		idx:   make(map[uint32]string),
		cols:  0,
	}
}

func (i *Indexer) AddData(data string, address ...uint32) {
	if !ApproveIndexerAddress(i.iType, address...) {
		return
	}

	completedAddress := MakeIndexerAddress(i.iType, address...)
	i.idx[completedAddress] = data
}

func (i *Indexer) GetDataPartialAddress(address ...uint32) (string, bool) {
	if !ApproveIndexerAddress(i.iType, address...) {
		return "", false
	}

	completedAddress := MakeIndexerAddress(i.iType, address...)
	data, ok := i.idx[completedAddress]

	return data, ok
}

func (i *Indexer) GetDataFullAddress(address uint32) (string, bool) {
	data, ok := i.idx[address]

	return data, ok
}

func (i *Indexer) ToTable(rows, cols uint8) [][]string {
	if i.iType != TYPE_INDEXER_TABLE_CELL {
		return [][]string{}
	}

	var (
		maxParams, ir, ic, ip uint8
		out                   [][]string
	)

	out = make([][]string, 0, 2)
	maxParams = 0xff
	ir, ic, ip = 0, 0, 0

	for ir = 0; ir < rows; ir++ {
		rowData := make([]string, 0, 2)
		for ic = 0; ic < cols; ic++ {
			paramsCache := make([]interface{}, 0, 2)
			for ip = 0; ip < maxParams; ip++ {
				data, ok := i.GetDataPartialAddress(uint32(ip), uint32(ir), uint32(ic))
				if ok {
					// есть параметр
					paramsCache = append(paramsCache, data)
					continue
				}
				if len(paramsCache) > 0 {
					str := i.makeStr(ip+1, paramsCache, false)
					rowData = append(rowData, str)
				}
				break
			}
		}
		out = append(out, rowData)
	}
	return [][]string{}
}

func (i *Indexer) ToList(rows uint8) []string {
	if i.iType != TYPE_INDEXER_LIST_ROW {
		return []string{}
	}

	var (
		maxParams, ir, ip uint8
		out               []string
	)

	out = make([]string, 0, 2)
	maxParams = 0xff
	ir, ip = 0, 0

	for ir = 0; ir < rows; ir++ {
		paramsCache := make([]interface{}, 0, 2)
		for ip = 0; ip < maxParams; ip++ {
			data, ok := i.GetDataPartialAddress(uint32(ip), uint32(ir))
			if ok {
				// есть параметр
				paramsCache = append(paramsCache, data)
				continue
			}
			if len(paramsCache) > 0 {
				str := i.makeStr(ip+1, paramsCache, true)
				out = append(out, str)
			}
			break
		}
	}
	return out
}

func (i *Indexer) ToString() string {
	if i.iType != TYPE_INDEXER_PARAM_IN_TEXT {
		return ""
	}
	var (
		maxParams uint8 = 0xff
		ip        uint8
	)

	paramsCache := make([]interface{}, 0, 2)
	for ip = 0; ip < maxParams; ip++ {
		data, ok := i.GetDataPartialAddress(uint32(ip))
		if ok {
			// есть параметр
			paramsCache = append(paramsCache, data)
			continue
		}
		if len(paramsCache) > 0 {
			return i.makeStr(ip+1, paramsCache, true)
		}
	}
	return ""
}

func (i *Indexer) makeStr(countParams uint8, data []interface{}, useSplit32 bool) string {
	var formatStr string

	if useSplit32 {
		formatStr = strings.Repeat("%s ", int(countParams))
		formatStr = strings.TrimRight(formatStr, " ")
	} else {
		formatStr = strings.Repeat("%s", int(countParams))
	}

	return fmt.Sprintf(formatStr, data...)
}
