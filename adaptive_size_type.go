package termui

const (
	adaptiveSizeZero int = 0
	adaptiveSizeMin  int = 10
	adaptiveSizeMax  int = 100
)

// AdaptiveSize тип для указания адаптивных размеров под виджеты.
// числа от 10 до 100. Остальные числа ограничиваются в этом диапазоне
// Далее при построении сетки виджета эти размеры будут разделены на 100
// что бы определить относительный размер каждой клетки в сетке
type AdaptiveSize int

func NewAdaptiveSize(in int) AdaptiveSize {
	if in == adaptiveSizeZero || in == adaptiveSizeMax {
		return AdaptiveSize(in)
	}
	// ограничиваем от 10 до 100
	if in <= adaptiveSizeMin {
		return AdaptiveSize(adaptiveSizeMin)
	}
	if in > adaptiveSizeMax {
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

func (as AdaptiveSize) FloatSize() float64 {
	res := float64(as) / float64(adaptiveSizeMax)
	if res > 1 {
		return 0.0
	}

	return float64(as) / float64(adaptiveSizeMax)
}

func (as AdaptiveSize) Sum(num AdaptiveSize) AdaptiveSize {
	sum := int(as) + int(num)
	if sum <= adaptiveSizeMax {
		return AdaptiveSize(sum)
	}
	diff := adaptiveSizeMax - sum

	return AdaptiveSize(diff)
}

func (as AdaptiveSize) IsMax() bool {
	return as == NewAdaptiveSizeMax()
}
