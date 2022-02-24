package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"time"
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
	radius float32
	xv     float32
	yv     float32
	color  color
}

func (b *ball) draw(pixels []byte) {
	for y := -b.radius; y < b.radius; y++ {
		for x := -b.radius; x < b.radius; x++ {
			if x*x+y*y < b.radius*b.radius {
				setPixel(int(b.x+x), int(b.y+y), b.color, pixels)
			}
		}
	}
}

func getCenter() pos {
	return pos{float32(winWidth) / 2, float32(winHeight) / 2}
}

func (b *ball) update(leftp *paddle, rightp *paddle, elapsedTime float32) {
	b.x += b.xv * elapsedTime
	b.y += b.yv * elapsedTime

	if b.y-b.radius < 0 || int(b.y+b.radius) > winHeight {
		b.yv = -b.yv
	}

	if b.x < 0 || int(b.x) > winWidth {
		b.pos = getCenter()
	}

	if b.x < leftp.x+leftp.w/2 {
		if b.y > leftp.y-leftp.h/2 && b.y < leftp.y+leftp.h/2 {
			b.xv = -b.xv

		}
	}
	if b.x > rightp.x-rightp.w/2 {
		if b.y > rightp.y-rightp.h/2 && b.y < rightp.y+rightp.h/2 {
			b.xv = -b.xv

		}
	}

}

type paddle struct {
	pos
	w     float32
	h     float32
	speed float32
	color color
}

func (p *paddle) draw(pixels []byte) {
	startX := int(p.x - p.w/2)
	startY := int(p.y - p.h/2)

	for y := 0; y < int(p.h); y++ {
		for x := 0; x < int(p.w); x++ {
			setPixel(startX+x, startY+y, p.color, pixels)
		}
	}
}

func (p *paddle) aiUpdate(b *ball, elapsedTime float32) {
	p.y = b.y
}

func (p *paddle) update(keyState []uint8, elapsedTime float32) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		p.y -= p.speed * elapsedTime
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		p.y += p.speed * elapsedTime
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

// Can make it so that it only clears where obj is drawn.
func clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
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

	player1 := paddle{pos{50, 100}, 20, 100, 300, color{255, 255, 255}}
	player2 := paddle{pos{float32(winWidth) - 50, 100}, 20, 100, 300, color{255, 255, 255}}
	ball := ball{pos{300, 300}, 20, 400, 400, color{255, 255, 255}}

	keyState := sdl.GetKeyboardState()

	var frameStart time.Time
	var elapsedTime float32

	// Game loop
	for {
		frameStart = time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		clear(pixels)

		player1.update(keyState, elapsedTime)
		player2.aiUpdate(&ball, elapsedTime)
		ball.update(&player1, &player2, elapsedTime)

		player1.draw(pixels)
		player2.draw(pixels)
		ball.draw(pixels)

		tex.Update(nil, pixels, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		//sdl.Delay(16)
		elapsedTime = float32(time.Since(frameStart).Seconds())
		fmt.Println(elapsedTime)
	}

}
