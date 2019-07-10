package draw

import (
	"math"
	"math/rand"
	"time"

	"github.com/templarrei/golearnyouatracer/geom"
)

type Camera struct {
	LLCorner, Horizontal, Vertical, Origin geom.Vec
	u, v                                   geom.Vec
	lensRadius                             float64
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewCamera(
	lFrom,
	lAt,
	vUp geom.Vec, vFOV,
	aspectRatio,
	focusDist,
	aperture float64) Camera {
	theta := vFOV * math.Pi / 180
	halfH := math.Tan(theta / 2)
	halfW := aspectRatio * halfH
	w := lFrom.Sub(lAt).ToUnit()
	u := vUp.Cross(w).ToUnit()
	v := w.Cross(u)

	x := u.Scale(halfW * focusDist)
	y := v.Scale(halfH * focusDist)
	z := w.Scale(focusDist)

	llCorner := lFrom.
		Sub(x).
		Sub(y).
		Sub(z)

	return Camera{
		Origin:     lFrom,
		LLCorner:   llCorner,
		Horizontal: x.Scale(2),
		Vertical:   y.Scale(2),
		lensRadius: aperture / 2,

		u: u,
		v: v,
	}
}

func (c Camera) Ray(s, t float64) geom.Ray {
	rd := randomInUnitDisc().Scale(c.lensRadius)
	offset := c.u.
		Scale(rd.X()).
		Add(c.v.Scale(rd.Y()))

	return geom.NewRay(
		c.Origin.Add(offset),
		c.LLCorner.
			Add(c.Horizontal.Scale(s)).
			Add(c.Vertical.Scale(t)).
			Sub(c.Origin).
			Sub(offset).
			ToUnit(),
	)
}

func randomInUnitDisc() geom.Vec {
	var p geom.Vec
	for {
		p = geom.NewVec(rand.Float64(), rand.Float64(), 0).Scale(2).Sub(geom.NewVec(1, 1, 0))
		if p.Dot(p) < 1 {
			return p
		}
	}
}
