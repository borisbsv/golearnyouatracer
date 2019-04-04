package draw

import (
	"math"

	"github.com/templarrei/golearnyouatracer/geom"
)

type Camera struct {
	LLCorner, Horizontal, Vertical, Origin geom.Vec
}

func NewCamera(lFrom, lAt, vUp geom.Vec, vFOV, aspectRatio float64) Camera {
	theta := vFOV * math.Pi / 180
	halfH := math.Tan(theta / 2)
	halfW := aspectRatio * halfH

	w := lFrom.Sub(lAt).ToUnit()
	u := vUp.Cross(w).ToUnit()
	v := w.Cross(u)

	llCorner := lFrom.
		Sub(u.Scale(halfW)).
		Sub(v.Scale(halfH)).
		Sub(w)
	return Camera{
		Origin:     lFrom,
		LLCorner:   llCorner,
		Horizontal: u.Scale(2 * halfW),
		Vertical:   v.Scale(2 * halfH),
	}
}

func (c Camera) Ray(s, t float64) geom.Ray {
	return geom.NewRay(
		c.Origin,
		c.LLCorner.
			Add(c.Horizontal.Scale(s)).
			Add(c.Vertical.Scale(t)).
			Sub(c.Origin).
			ToUnit(),
	)
}
