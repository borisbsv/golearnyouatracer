package draw

import "github.com/templarrei/golearnyouatracer/geom"

type Camera struct {
	LLCorner, Horizontal, Vertical, Origin geom.Vec
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
