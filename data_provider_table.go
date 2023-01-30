package termui

type DataProviderTable struct {
	idx   map[uint32]string
	cols  uint8      // количество колонок, если есть
	rows  uint8      // количество строк, есть всегда
	cache [][]string // для текста это всегда одна строка. 0 0 для текста, 0-n 0 для списка, 0-n 0-x для таблицы
}

func NewDataProviderTable() *DataProviderTable {
	return &DataProviderTable{
		idx:   make(map[uint32]string),
		cols:  1,
		rows:  1, // в любом случае 1 строка будет всегда
		cache: make([][]string, 0, 2),
	}
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
			colStr := MakeStr(uint8(p), cacheCol, false)
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

func (dpt *DataProviderTable) AddData(data string, row, col, param uint32) *DataProviderTable {
	var (
		newCols = uint8(row)
		newRows = uint8(col)
	)
	address := MakeDataProviderAddress(param, row, col)
	dpt.idx[address] = data
	if dpt.cols < newCols {
		dpt.cols = newCols
	}

	if dpt.rows < newRows {
		dpt.rows = newRows
	}

	return dpt
}

func (dpt *DataProviderTable) UpdateData(data string, address ...uint32) {
	var r, c, p uint32 = 0, 0, 0
	var iParam int

	if len(address) == 3 {
		p = address[0]
		r = address[1]
		c = address[2]
	}

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
	colStr := MakeStr(uint8(iParam), cacheCol, false)
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
