package main

const winWidth, winHeight int = 800, 600

type gameState int

const (
	start gameState = iota
	play
)

var state = start

type color struct {
	r, g, b byte
}

type pos struct {
	x, y float32
}

func getCenter() pos {
	return pos{float32(winWidth) / 2, float32(winHeight) / 2}
}

func lerp(a float32, b float32, pct float32) float32 {
	return a + (b-a)*pct
}

func setPixel(x, y int, c color, pixels []byte) {
	index := (y*winWidth + x) * 4

	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}

}
func clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}
