package termui

import (
	"strconv"
	"strings"
)

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
