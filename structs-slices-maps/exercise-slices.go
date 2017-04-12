package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	screen := make([][]uint8, dy)
	for y := range screen {
		screen[y] = make([]uint8, dx)
	}

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			screen[y][x] = uint8(x ^ y)
		}
	}
	return screen
}

func main() {
	pic.Show(Pic)
}
