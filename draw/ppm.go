package draw

import (
	"fmt"
	"io"
	"math"

	"github.com/templarrei/golearnyouatracer/geom"
)

type Hittable interface {
	Hit(r geom.Ray, tMin, tMax float64) (t float64, p geom.Vec, n geom.Vec)
}

type Scene struct {
	w, h float64
}

func NewScene(w, h float64) Scene {
	return Scene{w: w, h: h}
}

func (s Scene) WritePPM(w io.Writer, h Hittable) error {
	fmt.Fprintf(w, "P3\n%f %f\n255\n", s.w, s.h)

	llCorner := geom.NewVec(-2, -1, -1)
	horizontal := geom.NewVec(4, 0, 0)
	vertical := geom.NewVec(0, 2, 0)
	origin := geom.NewVec(0, 0, 0)

	for j := s.h - 1; j >= 0; j-- {
		for i := 0.0; i < s.w; i++ {
			u := i / s.w
			v := j / s.h

			r := geom.NewRay(
				origin,
				llCorner.
					Add(horizontal.Scale(u)).
					Add(vertical.Scale(v)),
			)

			c := color(r, h)
			ir := int(255.99 * c.R())
			ig := int(255.99 * c.G())
			ib := int(255.99 * c.B())
			if _, err := fmt.Fprintf(w, "%d %d %d\n", ir, ig, ib); err != nil {
				return err
			}
		}
	}
	return nil
}

func color(r geom.Ray, h Hittable) geom.Vec {
	if t, _, n := h.Hit(r, 0, math.MaxFloat64); t > 0 {
		return geom.NewVec(n.X()+1, n.Y()+1, n.Z()+1).Scale(0.5)
	}
	t := 0.5 * (r.Dir.ToUnit().Y() + 1)
	white := geom.NewVec(1, 1, 1).Scale(1 - t)
	blue := geom.NewVec(0.5, 0.7, 1).Scale(t)
	return white.Add(blue)
}
