package termui

import "sync"

type DataProviderDashboard struct {
	lib   *DashboardLib
	cache [][]string
	mx    sync.RWMutex
}

func NewDataProviderDashboard(lib *DashboardLib) *DataProviderDashboard {
	dash := &DataProviderDashboard{
		lib:   lib,
		mx:    sync.RWMutex{},
		cache: make([][]string, 0, lib.Rows()),
	}
	// сразу создаем кэш для отображения всех значков панели
	rows, cols := lib.Rows(), lib.Cols()
	idx := lib.Idx()
	// проставляем размеры кэша
	for r := 0; r < rows; r++ {
		dash.cache[r] = make([]string, cols)
	}
	// заполняем данные
	for _, item := range idx {
		dash.cache[item.R][item.C] = item.String()
	}

	return dash
}

func (dpd *DataProviderDashboard) GetTable() [][]string {
	return dpd.cache
}

func (dpd *DataProviderDashboard) On(name string) {
	val, ok := dpd.lib.GetItem(name)
	if !ok {
		return
	}
	dpd.insertInCache(val.R, val.C, val.String())
}

func (dpd *DataProviderDashboard) Off(name string) {
	val, ok := dpd.lib.GetItem(name)
	if !ok {
		return
	}
	dpd.insertInCache(val.R, val.C, string(EMPTY))
}

func (dpd *DataProviderDashboard) insertInCache(r, c uint32, data string) {
	if len(dpd.cache) <= int(r) {
		return
	}
	cols := dpd.cache[r]
	if len(cols) <= int(c) {
		return
	}
	dpd.cache[r][c] = data
}
