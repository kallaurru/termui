package tmpl

import . "github.com/kallaur/termui/v3"

const (
	CellTmplAdded     uint8 = 0x00
	CellTmplBadRowIdx uint8 = 0x01
	CellTmplIsExists  uint8 = 0x02
	CellTmplOtherErr  uint8 = 0x04
	adaptiveSizeMin   int   = 10
	adaptiveSizeMax   int   = 100
)

// AdaptiveSize тип для указания адаптивных размеров под виджеты.
// числа от 10 до 100. Остальные числа ограничиваются в этом диапазоне
type AdaptiveSize int

func NewAdaptiveSize(in int) AdaptiveSize {
	// ограничиваем от 10 до 100
	if in <= adaptiveSizeMin {
		return AdaptiveSize(adaptiveSizeMax)
	}
	if in >= adaptiveSizeMax {
		return AdaptiveSize(adaptiveSizeMax)
	}
	return AdaptiveSize(in)
}

func NewAdaptiveSizeFirstPercentile() AdaptiveSize {
	return AdaptiveSize(25)
}

func NewAdaptiveSizeTwoPercentile() AdaptiveSize {
	return AdaptiveSize(50)
}

func NewAdaptiveSizeThreePercentile() AdaptiveSize {
	return AdaptiveSize(75)
}

func NewAdaptiveSizeTwenty() AdaptiveSize {
	return AdaptiveSize(2 * adaptiveSizeMin)
}

func NewAdaptiveSizeThird() AdaptiveSize {
	return AdaptiveSize(3 * adaptiveSizeMin)
}

func NewAdaptiveSizeMin() AdaptiveSize {
	return AdaptiveSize(adaptiveSizeMin)
}

func NewAdaptiveSizeMax() AdaptiveSize {
	return AdaptiveSize(adaptiveSizeMax)
}

func (as AdaptiveSize) ToInt() int {
	return int(as)
}

func (as AdaptiveSize) ToUint8() uint8 {
	return uint8(as)
}

type GridTmpl struct {
	rows  uint8          // количество строк в шаблоне
	sizes []AdaptiveSize // высоты строк
	cells map[uint16]*CellDetail
}

func NewGridTmpl(rCount uint8, rSize ...AdaptiveSize) *GridTmpl {
	var (
		counter       = 0
		defSize uint8 = 0
		allS    uint8 = 0
		s             = make([]uint8, 0, int(rCount))
	)

	l := len(rSize)
	defSize = max / rCount

	for counter = 0; counter < int(rCount); counter++ {
		if l == 0 {
			allS += defSize
			s[counter] = defSize
			continue
		}
		if l-1 > counter {
			size := rSize[counter]
			allS += size
			s[counter] = size
			continue
		}
		sz := max - allS

	}
	return &GridTmpl{
		rows:  rCount,
		cells: make(map[uint16]*CellDetail),
		sizes: make([]uint8, 0, int(rCount)),
	}
}

// AddCell - индексация строк и столбцов начинается с 0
func (gt *GridTmpl) AddCell(row, col, size AdaptiveSize, widget Drawable) uint8 {
	if row > gt.rows {
		return CellTmplBadRowIdx
	}
	cd := NewCellDetail(row, col, size, widget)
	idx := cd.idx

	_, ok := gt.cells[idx]
	if ok {
		return CellTmplIsExists
	}
	gt.cells[idx] = cd

	return CellTmplAdded
}
