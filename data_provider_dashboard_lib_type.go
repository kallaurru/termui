package termui

import (
	"fmt"
	"strings"
)

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

func (dl *DashboardLib) AddString(r, c uint32, name, data string, styles ...Style) *DashboardLib {
	dl.addItem(r, c, data, name, styles...)
	return dl
}

func (dl *DashboardLib) AddCircleLineItem(r, c uint32, name, line string, styles ...Style) *DashboardLib {
	circled := ConvertSymToMarkers(line)
	dl.addItem(r, c, circled, name, styles...)

	return dl
}

func (dl *DashboardLib) AddSquaredLineItem(r, c uint32, name, line string, styles ...Style) *DashboardLib {
	var squared string

	line = strings.ToUpper(line)
	squared = ConvertSymToSquaredMarker(line)
	dl.addItem(r, c, squared, name, styles...)

	return dl
}

func (dl *DashboardLib) AddIconFree(r, c uint32, icon int32, name string) *DashboardLib {
	dl.addItem(r, c, string(icon), name)

	return dl
}

func (dl *DashboardLib) AddBoldNumber(r, c uint32, numb int32, name string, styles ...Style) *DashboardLib {
	dl.addItem(r, c, ConvertToBoldNumbers(numb), name, styles...)

	return dl
}

func (dl *DashboardLib) AddRomeNumber(r, c uint32, numb int32, name string, styles ...Style) *DashboardLib {
	dl.addItem(r, c, ConvertToRomeNumbers(numb), name, styles...)

	return dl
}

func (dl *DashboardLib) AddMonoNumber(r, c uint32, numb int32, name string, styles ...Style) *DashboardLib {
	dl.addItem(r, c, ConvertToMonoNumbers(numb), name, styles...)

	return dl
}

func (dl *DashboardLib) AddFullNumber(r, c uint32, numb int32, name string, styles ...Style) *DashboardLib {
	dl.addItem(r, c, ConvertToFullNumbers(numb), name, styles...)

	return dl
}

func (dl *DashboardLib) addItem(r, c uint32, data, name string, styles ...Style) {
	var (
		newRow        = uint8(r)
		newCol        = uint8(c)
		p      uint32 = 0
	)
	if name == "" {
		name = fmt.Sprintf("%d", MakeDataProviderAddress(p, r, c))
	}

	item := NewDashboardLibsItem(data, r, c, p, styles...)
	dl.lib[name] = item

	if dl.r < newRow+1 {
		dl.r = newRow + 1
	}

	if dl.c < newCol+1 {
		dl.c = newCol + 1
	}
}
