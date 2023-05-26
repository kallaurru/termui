package termui

import "sync"

type DataProviderList struct {
	idx   map[uint32]string
	rows  uint8
	cache []string
	mx    sync.RWMutex
	// ширина поля под параметр. Распространяется на параметры с одинаковым индексом во всех строках.
	// поля где ширина не важна можно пропустить установив 0
	width    []uint8 // ширина поля под параметр.
	useSplit bool    // использовать разделитель пробел между параметрами
}

func NewDataProviderList() *DataProviderList {
	return &DataProviderList{
		idx:      make(map[uint32]string),
		rows:     1,
		cache:    make([]string, 0, 2),
		mx:       sync.RWMutex{},
		width:    []uint8{0},
		useSplit: false,
	}
}

func (dpl *DataProviderList) UseSplit32() *DataProviderList {
	dpl.useSplit = true
	return dpl
}

func (dpl *DataProviderList) SetWidthParams(width []uint8) {
	dpl.width = make([]uint8, len(width), len(width))
	for idx, w := range width {
		dpl.width[idx] = w
	}
}

func (dpl *DataProviderList) SetExtTarget(t *[]string) {
	dpl.cache = *t
}

func (dpl *DataProviderList) Caching() {
	var p, r int
	var str string

	hasWR := dpl.width[0] > 0

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
		if hasWR {
			str = MakeStrStaticW(paramsCache, dpl.width, dpl.useSplit)
		} else {
			str = MakeStr(uint8(len(paramsCache)), paramsCache, dpl.useSplit)
		}
		cache = append(cache, str)
	}
	dpl.cache = cache
}

func (dpl *DataProviderList) GetList() []string {
	return dpl.cache
}

func (dpl *DataProviderList) GetListPtr() *[]string {
	return &dpl.cache
}

func (dpl *DataProviderList) AddData(data string, param uint32, row uint32) {
	address := MakeDataProviderAddress(param, row, 0)
	dpl.idx[address] = data

	if dpl.rows < uint8(row+1) {
		dpl.rows = uint8(row + 1)
	}
}

func (dpl *DataProviderList) UpdateDataFromMap(data map[uint32]string) {
	if data == nil {
		return
	}

	for addr, line := range data {
		p, r, _ := ParseDataProviderAddress(addr)
		dpl.updateData(line, p, r)
	}
}

// UpdateData - адрес подается след образом - row, col, param
func (dpl *DataProviderList) UpdateData(data string, address ...uint32) {
	p, r, _ := GetAddressElements(address...)
	dpl.updateData(data, p, r)
}

func (dpl *DataProviderList) updateData(data string, p, r uint32) {
	var paramMax = 0xff // максимальное количество параметров
	var pr int          // счетчик параметров при обходе

	if uint8(r) >= dpl.rows {
		return
	}
	if len(dpl.cache) == 0 {
		dpl.Caching()
	}
	dpl.mx.Lock()
	defer dpl.mx.Unlock()

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
	hasWR := dpl.width[0] > 0
	str := ""
	if hasWR {
		str = MakeStrStaticW(cacheParams, dpl.width, dpl.useSplit)
	} else {
		str = MakeStr(uint8(pr), cacheParams, dpl.useSplit)
	}

	dpl.cache[int(r)] = str
}
