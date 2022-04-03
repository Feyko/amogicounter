package highlight

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
)

func Highlight(img image.Image, spots []image.Rectangle, opacity float64) image.Image {
	shadowed := overlayShadow(img, opacity)
	for _, spot := range spots {
		originalArea := imaging.Crop(img, spot)
		shadowed = imaging.Paste(shadowed, originalArea, spot.Min)
	}
	return shadowed
}

func overlayShadow(img image.Image, opacity float64) image.Image {
	size := img.Bounds().Size()
	shadow := imaging.New(size.X, size.Y, color.Black)
	return imaging.OverlayCenter(img, shadow, opacity)
}
