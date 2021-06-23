package main

import (
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
	for !win.Closed() {
		win.Clear(colornames.Black)
		imd := imdraw.New(nil)
		square.Draw(imd)
		imd.Draw(win)
		win.Update()
		square.Move()
		time.Sleep(1 * time.Second)
	}
}
