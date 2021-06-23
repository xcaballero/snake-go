package main

import (
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
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

	square := NewSquare(&Vector{X: (maxX - minX) / 2, Y: (maxY - minY) / 2})
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

type Vector struct {
	X float64
	Y float64
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Square struct {
	position  Vector
	direction Direction
	size      float64
}

func NewSquare(center *Vector) *Square {
	return &Square{
		position:  *center,
		direction: East,
		size:      squareSize,
	}
}

func (s *Square) Draw(imd *imdraw.IMDraw) {
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(s.position.X-s.size, s.position.Y-s.size))
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(s.position.X-s.size, s.position.Y+s.size))
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(s.position.X+s.size, s.position.Y+s.size))
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(s.position.X+s.size, s.position.Y-s.size))
	imd.Polygon(0)
}

func (s *Square) Move() {
	switch s.direction {
	case North:
		s.position.Y += 2 * s.size
	case East:
		s.position.X += 2 * s.size
	case South:
		s.position.Y -= 2 * s.size
	case West:
		s.position.X -= 2 * s.size
	}
	s.position.X = math.Mod(s.position.X, maxX)
	s.position.Y = math.Mod(s.position.Y, maxY)
}
