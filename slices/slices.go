package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ay := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		ax := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			ax[x] = uint8((x ^ y) / 2)
		}
		ay[y] = ax
	}
	return ay
}

func main() {
	pic.Show(Pic)
}
