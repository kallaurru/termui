package termui

import (
	"strconv"
	"strings"
)

// ConvertToRomeNumbers Актуальны числа от 1 до 12, 50, 100, 500, 1000
func ConvertToRomeNumbers(in int32) string {
	switch in {
	case 1:
		return string(ROME_1)
	case 2:
		return string(ROME_2)
	case 3:
		return string(ROME_3)
	case 4:
		return string(ROME_4)
	case 5:
		return string(ROME_5)
	case 6:
		return string(ROME_6)
	case 7:
		return string(ROME_7)
	case 8:
		return string(ROME_8)
	case 9:
		return string(ROME_9)
	case 10:
		return string(ROME_10)
	case 11:
		return string(ROME_11)
	case 12:
		return string(ROME_12)
	case 50:
		return string(ROME_50)
	case 100:
		return string(ROME_100)
	case 500:
		return string(ROME_500)
	case 1000:
		return string(ROME_1000)
	default:
		return ""
	}
}

func ConvertToMonoNumbers(in int32) string {
	m := makeMapMonoNumbers()

	line := []rune(strconv.Itoa(int(in)))

	cache := make([]int32, 0, len(line))

	for _, code := range line {
		mono, ok := m[code]
		if ok {
			cache = append(cache, mono)
		} else {
			cache = append(cache, EMPTY)
		}
	}

	return string(cache)
}

func ConvertToFullNumbers(in int32) string {
	m := makeMapFullNumbers()

	line := []rune(strconv.Itoa(int(in)))

	cache := make([]int32, 0, len(line))

	for _, code := range line {
		mono, ok := m[code]
		if ok {
			cache = append(cache, mono)
		} else {
			cache = append(cache, EMPTY)
		}
	}

	return string(cache)
}

func ConvertSymToMarkers(in string) string {
	m := makeMapCircleLat()
	line := []rune(in)
	cache := make([]int32, 0, len(line))
	for _, code := range line {
		mark, ok := m[code]
		if ok {
			cache = append(cache, mark, EMPTY)
		} else {
			cache = append(cache, EMPTY, EMPTY)
		}
	}

	return strings.TrimRight(string(cache), " ")
}

func ConvertSymToSquaredMarker(in string) string {
	m := makeMapSquaredLat()
	line := []rune(in)
	cache := make([]int32, 0, len(line))
	for _, code := range line {
		mark, ok := m[code]
		if ok {
			cache = append(cache, mark, EMPTY)
		} else {
			cache = append(cache, EMPTY, EMPTY)
		}
	}

	return strings.TrimRight(string(cache), " ")
}

func makeMapMonoNumbers() map[int32]int32 {
	var (
		convMap               = make(map[int32]int32)
		startLatNumber  int32 = 0x0030
		startMonoNumber int32 = 0x1d7f6
		count                 = 10
	)

	for i := 0; i < count; i++ {
		diff := int32(i)
		convMap[startLatNumber+diff] = startMonoNumber + diff
	}

	return convMap
}

func makeMapFullNumbers() map[int32]int32 {
	var (
		convMap              = make(map[int32]int32)
		startLatNumber int32 = 0x0030
		startFull      int32 = 0xff10
		count                = 10
	)

	for i := 0; i < count; i++ {
		diff := int32(i)
		convMap[startLatNumber+diff] = startFull + diff
	}

	return convMap
}

func makeMapCircleLat() map[int32]int32 {
	var (
		convMap              = make(map[int32]int32)
		startKeyNumber int32 = 0x0041
		startValNumber int32 = 0x24b6
		count                = 26
	)

	for i := 0; i < count; i++ {
		diff := int32(i)
		convMap[startKeyNumber+diff] = startValNumber + diff
	}

	// conv small
	startKeyNumber = 0x0061
	startValNumber = 0x24d0

	for i := 0; i < count; i++ {
		diff := int32(i)
		convMap[startKeyNumber+diff] = startValNumber + diff
	}

	return convMap
}

func makeMapSquaredLat() map[int32]int32 {
	var (
		convMap              = make(map[int32]int32)
		startKeyNumber int32 = 0x0041
		startValNumber int32 = 0x1f130
		count                = 26
	)

	for i := 0; i < count; i++ {
		diff := int32(i)
		convMap[startKeyNumber+diff] = startValNumber + diff
	}

	return convMap
}
