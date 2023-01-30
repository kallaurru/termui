package termui

type DataProviderList struct {
	idx   map[uint32]string
	rows  uint8
	cache []string
}

func NewDataProviderList() *DataProviderList {
	return &DataProviderList{
		idx:   make(map[uint32]string),
		rows:  0,
		cache: make([]string, 2),
	}
}

func (dpl *DataProviderList) GetList() []string {
	return dpl.cache
}

func (dpl *DataProviderList) AddData(data string, param uint32, row uint32) {
	address := MakeDataProviderAddress(param, row, 0)
	dpl.idx[address] = data

	if dpl.rows < uint8(row) {
		dpl.rows = uint8(row)
	}
}

func (dpl *DataProviderList) UpdateData(data string, address ...uint32) {
	var p, r uint32 = 0, 0
	var paramMax = 0xff // максимальное количество параметров
	var pr int          // счетчик параметров при обходе

	if len(address) < 2 {
		p = address[0]
	}
	if len(address) > 1 {
		r = address[1]
	}
	dpl.AddData(data, p, r)

	cacheParams := make([]interface{}, 0, 2)

	for pr = 0; pr < paramMax; pr++ {
		addr := MakeDataProviderAddress(uint32(paramMax), r, 0)
		value, ok := dpl.idx[addr]
		if !ok {
			break
		}
		cacheParams = append(cacheParams, value)
	}
	str := MakeStr(uint8(pr), cacheParams, true)
	if len(dpl.cache) <= int(r) {
		return
	}
	dpl.cache[int(r)] = str
}
