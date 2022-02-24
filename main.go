package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

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

	var controllerHandlers []*sdl.GameController
	for i := 0; i < sdl.NumJoysticks(); i++ {
		controllerHandlers = append(controllerHandlers, sdl.GameControllerOpen(i))
		defer controllerHandlers[i].Close()
	}

	pixels := make([]byte, winWidth*winHeight*4)

	player1 := paddle{pos{50, 100}, 20, 100, 300, 0, color{255, 255, 255}}
	player2 := paddle{pos{float32(winWidth) - 50, 100}, 20, 100, 300, 0, color{255, 255, 255}}
	ball := ball{pos{300, 300}, 20, 400, 400, color{255, 255, 255}}

	keyState := sdl.GetKeyboardState()

	var frameStart time.Time
	var elapsedTime float32
	var controllerAxis int16

	// Game loop
	for {
		frameStart = time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		for _, c := range controllerHandlers {
			if c != nil {
				controllerAxis = c.Axis(sdl.CONTROLLER_AXIS_LEFTY)
			}
		}

		if state == play {
			player1.update(keyState, controllerAxis, elapsedTime)
			player2.aiUpdate(&ball, elapsedTime)
			ball.update(&player1, &player2, elapsedTime)
		} else if state == start {
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				if player1.score == 3 || player2.score == 3 {
					player1.score = 0
					player2.score = 0
				}
				state = play
			}
		}

		clear(pixels)
		player1.draw(pixels)
		player2.draw(pixels)
		ball.draw(pixels)

		tex.Update(nil, pixels, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		elapsedTime = float32(time.Since(frameStart).Seconds())
		if elapsedTime < .005 {
			sdl.Delay(5 - uint32(elapsedTime*1000.0))
			elapsedTime = float32(time.Since(frameStart).Seconds())
		}
	}

}
