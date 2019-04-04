package draw

import (
	"math"

	"github.com/templarrei/golearnyouatracer/geom"
)

type Camera struct {
	LLCorner, Horizontal, Vertical, Origin geom.Vec

	theta, halfW, halfH float64
}

func NewCamera(vFOV, aspectRatio float64) Camera {
	theta := vFOV * math.Pi / 180
	halfH := math.Tan(theta / 2)
	halfW := aspectRatio * halfH

	return Camera{
		LLCorner:   geom.NewVec(-halfW, -halfH, -1),
		Horizontal: geom.NewVec(halfW*2, 0, 0),
		Vertical:   geom.NewVec(0, halfH*2, 0),
		Origin:     geom.NewVec(0, 0, 0),
		theta:      theta,
		halfW:      halfW,
		halfH:      halfH,
	}
}

func (c Camera) Ray(u, v float64) geom.Ray {
	return geom.NewRay(
		c.Origin,
		c.LLCorner.
			Add(c.Horizontal.Scale(u)).
			Add(c.Vertical.Scale(v)).
			ToUnit(),
	)
}
