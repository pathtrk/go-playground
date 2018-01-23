package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image takes width and height for its initialization
type Image struct {
	width  int
	height int
}

// Bounds shows image boundary in coordinates
func (i *Image) Bounds() (r image.Rectangle) {
	return image.Rect(0, 0, i.width, i.height)
}

// ColorModel returns Image type's color mode which is color.RGBAModel
func (i *Image) ColorModel() (m color.Model) {
	return color.RGBAModel
}

// At takes x and y coordinnates and the color of the position
func (i *Image) At(x, y int) (c color.Color) {
	v := uint8((x ^ y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := &Image{200, 200}
	pic.ShowImage(m)
}
