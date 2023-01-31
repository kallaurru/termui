package termui

type DataProviderList struct {
	idx   map[uint32]string
	rows  uint8
	cache []string
}

func NewDataProviderList() *DataProviderList {
	return &DataProviderList{
		idx:   make(map[uint32]string),
		rows:  1,
		cache: make([]string, 2),
	}
}

func (dpl *DataProviderList) Caching() {
	var p, r int

	cache := make([]string, 0, 2)
	for r = 0; r < int(dpl.rows); r++ {
		paramsCache := make([]interface{}, 0, 2)
		for p = 0; p < 0xff; p++ {
			addr := MakeDataProviderAddress(uint32(p), uint32(r), 0)
			val, ok := dpl.idx[addr]
			if !ok {
				break
			}
			paramsCache = append(paramsCache, val)
		}
		str := MakeStr(uint8(len(paramsCache)), paramsCache, true)
		cache = append(cache, str)
	}
	dpl.cache = cache
}

func (dpl *DataProviderList) GetList() []string {
	return dpl.cache
}

func (dpl *DataProviderList) AddData(data string, param uint32, row uint32) {
	address := MakeDataProviderAddress(param, row, 0)
	dpl.idx[address] = data

	if dpl.rows < uint8(row+1) {
		dpl.rows = uint8(row + 1)
	}
}

// UpdateData - адрес подается след образом - row, col, param
func (dpl *DataProviderList) UpdateData(data string, address ...uint32) {
	var paramMax = 0xff // максимальное количество параметров
	var pr int          // счетчик параметров при обходе

	p, r, _ := dpl.getAddrElements(address...)

	if uint8(r) >= dpl.rows {
		return
	}
	if len(dpl.cache) == 0 {
		dpl.Caching()
	}
	dpl.AddData(data, p, r)

	cacheParams := make([]interface{}, 0, 2)

	for pr = 0; pr < paramMax; pr++ {
		addr := MakeDataProviderAddress(uint32(pr), r, 0)
		value, ok := dpl.idx[addr]
		if !ok {
			break
		}
		cacheParams = append(cacheParams, value)
	}
	str := MakeStr(uint8(pr), cacheParams, true)

	dpl.cache[int(r)] = str
}

func (dpl *DataProviderList) getAddrElements(address ...uint32) (p, r, c uint32) {
	switch len(address) {
	default:
		return 0, 0, 0
	case 1:
		return 0, address[0], 0
	case 2:
		return address[1], address[0], 0
	}
}
