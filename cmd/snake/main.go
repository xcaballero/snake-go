package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/xcaballero/snake-go/pkg/shape"
	"github.com/xcaballero/snake-go/pkg/topology"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

var minX = float64(0)
var minY = float64(0)
var maxX = float64(1020)
var maxY = float64(780)

const squareSize = 10

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(minX, minY, maxX, maxY),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	square := shape.NewSquare(&topology.Vector{X: (maxX - minX) / 2, Y: (maxY - minY) / 2}, squareSize, maxX, maxY)
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				square.Move()
			}
		}
	}()

	for !win.Closed() {
		win.Clear(colornames.Black)
		imd := imdraw.New(nil)
		square.Draw(imd)
		imd.Draw(win)
		win.Update()
		listenForArrow(win, square)
	}
	ticker.Stop()
	done <- true
	fmt.Println("Shutting down")
}

func listenForArrow(win *pixelgl.Window, shape shape.Shape) {
	if win.Pressed(pixelgl.KeyUp) {
		shape.Head(topology.North)
		return
	}
	if win.Pressed(pixelgl.KeyDown) {
		shape.Head(topology.South)
		return
	}
	if win.Pressed(pixelgl.KeyLeft) {
		shape.Head(topology.West)
		return
	}
	if win.Pressed(pixelgl.KeyRight) {
		shape.Head(topology.East)
		return
	}
}
