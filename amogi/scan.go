package amogi

import (
	"github.com/disintegration/imaging"
	"image"
)

func ScanAmogi(img image.Image) []image.Rectangle {
	spots := make([]image.Rectangle, 0)
	for _, size := range AmogiSizes {
		spots = append(spots, countAmogiForSize(img, size)...)
	}
	return spots
}

func countAmogiForSize(img image.Image, size image.Point) []image.Rectangle {
	max := getMaximumSubimageCoords(img.Bounds().Size(), size)
	spots := make([]image.Rectangle, 0)
	for y := 0; y <= max.Y; y++ {
		for x := 0; x <= max.X; x++ {
			coords := image.Point{x, y}
			areaToScan := sizeToRectangle(coords, size)
			subImage := imaging.Crop(img, areaToScan)
			if IsAmogus(subImage) {
				spots = append(spots, areaToScan)
			}
		}
	}
	return spots
}

func getMaximumSubimageCoords(imgSize, subImgSize image.Point) image.Point {
	return image.Point{imgSize.X - subImgSize.X + 1, imgSize.Y - subImgSize.Y + 1}
}

func sizeToRectangle(coords image.Point, size image.Point) image.Rectangle {
	return image.Rectangle{
		Min: coords,
		Max: image.Point{
			coords.X + size.X,
			coords.Y + size.Y,
		},
	}
}

func IsAmogus(img image.Image) bool {
	for _, amogus := range Amogi {
		if fitsAmogus(img, amogus) {
			return true
		}
	}
	return false
}

func fitsAmogus(img image.Image, amogus Amogus) bool {
	return isSizeOfAmogus(img.Bounds().Size(), amogus) && looksLikeAmogus(img, amogus)
}

func isSizeOfAmogus(size image.Point, amogus Amogus) bool {
	return size.X == len(amogus[0]) && size.Y == len(amogus)
}

func looksLikeAmogus(img image.Image, amogus Amogus) bool {
	bodyColor := img.At(1, 0)
	for y, line := range amogus {
		for x, pixel := range line {
			if (bodyColor == img.At(x, y)) != pixel {
				return false
			}
		}
	}
	return true
}
