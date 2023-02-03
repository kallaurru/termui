package termui

import "sync"

type DataProviderText struct {
	idx    map[uint32]string // не переделывай на массив
	params uint8             // пишем максимальный индекс который был считан при добавлении данных
	cache  string
	mx     sync.RWMutex
}

func NewDataProviderText() *DataProviderText {
	return &DataProviderText{
		idx:    make(map[uint32]string),
		params: 1,
		cache:  "",
		mx:     sync.RWMutex{},
	}
}

func (dpt *DataProviderText) GetText() string {
	return dpt.cache
}

func (dpt *DataProviderText) AddData(data string, address uint32) {
	dpt.idx[address] = data

	if dpt.params < uint8(address) {
		dpt.params = uint8(address)
	}
}

func (dpt *DataProviderText) UpdateData(data string, address ...uint32) {
	var addr uint32 = 0
	var i uint32

	if len(address) > 0 {
		addr = address[0]
	}
	dpt.mx.Lock()

	defer dpt.mx.Unlock()

	dpt.AddData(data, addr)
	cache := make([]interface{}, 0, 2)
	for i = 0; i < uint32(dpt.params); i++ {
		val, ok := dpt.idx[i]
		if ok {
			cache = append(cache, val)
		}
	}

	dpt.cache = MakeStr(dpt.params, cache, true)
}
