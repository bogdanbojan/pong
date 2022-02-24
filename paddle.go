package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type paddle struct {
	pos
	w     float32
	h     float32
	speed float32
	score int
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

	numX := lerp(p.x, getCenter().x, 0.2)
	drawNumber(pos{numX, 35}, p.color, 10, p.score, pixels)
}

func (p *paddle) aiUpdate(b *ball, elapsedTime float32) {
	p.y = b.y
}

func (p *paddle) update(keyState []uint8, controllerAxis int16, elapsedTime float32) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		p.y -= p.speed * elapsedTime
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		p.y += p.speed * elapsedTime
	}

	if math.Abs(float64(controllerAxis)) > 1500 {
		pct := float32(controllerAxis) / 32767.0
		p.y += p.speed * pct * elapsedTime
	}

}
