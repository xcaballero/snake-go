package shape

import (
	"github.com/faiface/pixel/imdraw"
	"github.com/xcaballero/snake-go/pkg/topology"
)

// Shape represents an object that can be Drawn and Moved
type Shape interface {
	Draw(imd *imdraw.IMDraw)
	Move()
	Head(directrion topology.Direction)
}
