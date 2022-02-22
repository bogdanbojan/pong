package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int = 800, 600

type color struct {
	r, g, b byte
}

type pos struct {
	x, y float32
}

type ball struct {
	pos
	radius int
	xv     float32
	yv     float32
	color  color
}

type paddle struct {
	pos
	w     int
	h     int
	color color
}

func (b *ball) draw(pixels []byte) {
	for y := -b.radius; y < b.radius; y++ {
		for x := -b.radius; x < b.radius; x++ {
			if x*x+y*y < b.radius*b.radius {
				setPixel(int(b.x)+x, int(b.y)+y, b.color, pixels)
			}
		}
	}
}

func (p *paddle) draw(pixels []byte) {
	startX := int(p.x) - p.w/2
	startY := int(p.y) - p.h/2

	for y := 0; y < p.h; y++ {
		for x := 0; x < p.w; x++ {
			setPixel(startX+x, startY+y, p.color, pixels)
		}
	}
}

func setPixel(x, y int, c color, pixels []byte) {
	index := (y*winWidth + x) * 4

	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}

}

func main() {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tex.Destroy()

	pixels := make([]byte, winWidth*winHeight*4)

	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			setPixel(x, y, color{byte(0), byte(0), 0}, pixels)
		}
	}

	tex.Update(nil, pixels, winWidth*4)
	renderer.Copy(tex, nil, nil)
	renderer.Present()

	p1 := paddle{pos{100, 100}, 20, 100, color{255, 255, 255}}
	ball := ball{pos{300, 300}, 20, 0, 0, color{255, 255, 255}}

	// Game loop
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		p1.draw(pixels)
		ball.draw(pixels)
		sdl.Delay(16)
	}

}
