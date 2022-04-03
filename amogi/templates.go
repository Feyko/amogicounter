package amogi

import (
	"image"
)

var Amogi = []Amogus{
	{
		{false, true, true, true},
		{true, true, false, false},
		{true, true, true, true},
		{false, true, false, true},
	},
	{
		{false, true, true, true},
		{true, true, false, false},
		{true, true, true, true},
		{false, true, true, true},
		{false, true, false, true},
	},
	{
		{false, true, true, true},
		{true, true, false, false},
		{true, true, true, true},
		{true, true, true, true},
		{false, true, false, true},
	},
}

var AmogiSizes []image.Point

type Amogus [][]pixel

func (amogus Amogus) Size() image.Point {
	return image.Point{len(amogus[0]), len(amogus)}
}

type pixel bool

func init() {
	for _, amongus := range Amogi {
		Amogi = append(Amogi, flipHorizontally(amongus))
		Amogi = append(Amogi, flipVertically(amongus))
		Amogi = append(Amogi, flipVertically(flipHorizontally(amongus)))
	}
	AmogiSizes = computeSizes(Amogi)
}

func flipHorizontally(amogus Amogus) Amogus {
	var out Amogus
	for _, line := range amogus {
		out = append(out, reverseSlice(line))
	}
	return out
}

func flipVertically(amogus Amogus) Amogus {
	return reverseSlice(amogus)
}

func reverseSlice[T any](slice []T) []T {
	sliceLen := len(slice)
	out := make([]T, sliceLen)

	for i, n := range slice {
		j := sliceLen - i - 1
		out[j] = n
	}

	return out
}

func computeSizes(amogi []Amogus) []image.Point {
	out := make([]image.Point, 0)
	for _, amongus := range amogi {
		size := amongus.Size()
		if !sizeIsKnown(size, out) {
			out = append(out, size)
		}
	}
	return out
}

func sizeIsKnown(size image.Point, knownSizes []image.Point) bool {
	for _, known := range knownSizes {
		if known == size {
			return true
		}
	}
	return false
}
