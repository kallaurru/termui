package termui

import "strings"

type DashboardLib struct {
	lib map[string]*DashboardLibItem
	r   uint8
	c   uint8
}

func NewDashboardLib() *DashboardLib {
	return &DashboardLib{
		lib: make(map[string]*DashboardLibItem),
		r:   0,
		c:   0,
	}
}

func (dl *DashboardLib) Idx() map[uint32]*DashboardLibItem {
	idx := make(map[uint32]*DashboardLibItem, dl.Size())
	for _, item := range dl.lib {
		addr := MakeDataProviderAddress(item.P, item.R, item.C)
		idx[addr] = item
	}

	return idx
}

func (dl *DashboardLib) Rows() int {
	return int(dl.r)
}

func (dl *DashboardLib) Cols() int {
	return int(dl.c)
}

func (dl *DashboardLib) Size() int {
	return len(dl.lib)
}

func (dl *DashboardLib) Names() []string {
	keys := make([]string, len(dl.lib))
	i := 0
	for k := range dl.lib {
		keys[i] = k
		i++
	}
	return keys
}

func (dl *DashboardLib) GetItem(name string) (*DashboardLibItem, bool) {
	val, ok := dl.lib[name]
	return val, ok
}

func (dl *DashboardLib) AddCircleLineItem(r, c, p uint32, name, line string, styles ...Style) *DashboardLib {
	circled := ConvertSymToMarkers(line)
	dl.addItem(r, c, p, circled, name, styles...)

	return dl
}

func (dl *DashboardLib) AddSquaredLineItem(r, c, p uint32, name, line string, styles ...Style) *DashboardLib {
	var squared string

	line = strings.ToUpper(line)
	squared = ConvertSymToSquaredMarker(line)
	dl.addItem(r, c, p, squared, name, styles...)

	return dl
}

func (dl *DashboardLib) AddIconFree(r, c, p uint32, icon int32, name string) *DashboardLib {
	dl.addItem(r, c, p, string(icon), name)

	return dl
}

func (dl *DashboardLib) AddBoldNumber(r, c, p uint32, numb int32, name string, styles ...Style) *DashboardLib {
	dl.addItem(r, c, p, ConvertToBoldNumbers(numb), name, styles...)

	return dl
}

func (dl *DashboardLib) AddRomeNumber(r, c, p uint32, numb int32, name string, styles ...Style) *DashboardLib {
	dl.addItem(r, c, p, ConvertToRomeNumbers(numb), name, styles...)

	return dl
}

func (dl *DashboardLib) AddMonoNumber(r, c, p uint32, numb int32, name string, styles ...Style) *DashboardLib {
	dl.addItem(r, c, p, ConvertToMonoNumbers(numb), name, styles...)

	return dl
}

func (dl *DashboardLib) AddFullNumber(r, c, p uint32, numb int32, name string, styles ...Style) *DashboardLib {
	dl.addItem(r, c, p, ConvertToFullNumbers(numb), name, styles...)

	return dl
}

func (dl *DashboardLib) addItem(r, c, p uint32, data, name string, styles ...Style) {
	var (
		newRow = uint8(r)
		newCol = uint8(c)
	)

	item := NewDashboardLibsItem(data, r, c, p, styles...)
	dl.lib[name] = item

	if dl.r < newRow+1 {
		dl.r = newRow + 1
	}

	if dl.c < newCol+1 {
		dl.c = newCol + 1
	}
}
