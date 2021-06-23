package shape

import "github.com/faiface/pixel/imdraw"

// Shape represents an object that can be Drawn and Moved
type Shape interface {
	Draw(imd *imdraw.IMDraw)
	Move()
}
