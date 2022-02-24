package main

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

func (b *ball) update(leftp *paddle, rightp *paddle, elapsedTime float32) {
	b.x += b.xv * elapsedTime
	b.y += b.yv * elapsedTime

	if b.y-b.radius < 0 || int(b.y+b.radius) > winHeight {
		b.yv = -b.yv
	}

	if b.x < 0 {
		rightp.score++
		b.pos = getCenter()
		state = start
	} else if int(b.x) > winWidth {
		leftp.score++
		b.pos = getCenter()
		state = start
	}

	if b.x-b.radius < leftp.x+leftp.w/2 {
		if b.y > leftp.y-leftp.h/2 && b.y < leftp.y+leftp.h/2 {
			b.xv = -b.xv
			b.x = leftp.x + leftp.w/2.0 + b.radius

		}
	}
	if b.x+b.radius > rightp.x-rightp.w/2 {
		if b.y > rightp.y-rightp.h/2 && b.y < rightp.y+rightp.h/2 {
			b.xv = -b.xv
			b.x = rightp.x - rightp.w/2.0 - b.radius

		}
	}

}
