package shape

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/xcaballero/snake-go/pkg/topology"
)

type square struct {
	position  topology.Vector
	direction topology.Direction
	size      float64
	maxX      float64
	maxY      float64
}

// NewSquare returns
func NewSquare(center *topology.Vector, size float64, maxX float64, maxY float64) Shape {
	return &square{
		position:  *center,
		direction: topology.East,
		size:      size,
		maxX:      maxX,
		maxY:      maxY,
	}
}

func (s *square) Draw(imd *imdraw.IMDraw) {
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

func (s *square) Move() {
	switch s.direction {
	case topology.North:
		s.position.Y += 2 * s.size
	case topology.East:
		s.position.X += 2 * s.size
	case topology.South:
		s.position.Y -= 2 * s.size
	case topology.West:
		s.position.X -= 2 * s.size
	}
	s.position.X = math.Mod(s.position.X, s.maxX)
	s.position.Y = math.Mod(s.position.Y, s.maxY)
}

func (s *square) Head(directrion topology.Direction) {
	s.direction = directrion
}
