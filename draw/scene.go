package draw

import (
	"fmt"
	"io"
	"math"
	"math/rand"

	"github.com/borisbsv/golearnyouatracer/geom"
)

const bias = 0.001

type Hittable interface {
	Hit(r geom.Ray, tMin, tMax float64) (t float64, p geom.Vec, n geom.Vec, m geom.Material)
}

type Scene struct {
	w, h float64
}

func NewScene(w, h float64) Scene {
	return Scene{w: w, h: h}
}

func (s Scene) WritePPM(w io.Writer, h Hittable, samples float64, c Camera) error {
	fmt.Fprintf(w, "P3\n%f %f\n255\n", s.w, s.h)

	for j := s.h - 1; j >= 0; j-- {
		for i := 0.0; i < s.w; i++ {
			col := geom.NewVec(0, 0, 0)
			for sm := 0.0; sm < samples; sm++ {
				u := (i + rand.Float64()) / s.w
				v := (j + rand.Float64()) / s.h
				r := c.Ray(u, v)
				col = col.Add(color(r, h, 0))
			}
			// Apply Gamma
			col = col.Scale(1 / samples).Gamma(2)
			ir := int(255.99 * col.R())
			ig := int(255.99 * col.G())
			ib := int(255.99 * col.B())
			if _, err := fmt.Fprintf(w, "%d %d %d\n", ir, ig, ib); err != nil {
				return err
			}
		}
	}
	return nil
}

func color(r geom.Ray, h Hittable, depth int) geom.Vec {
	if depth > 50 {
		return geom.NewVec(0, 0, 0)
	}

	if t, p, n, mat := h.Hit(r, bias, math.MaxFloat64); t > 0 {
		scattered, attenuation, ok := mat.Scatter(r, p, n)
		if !ok {
			return geom.NewVec(0, 0, 0)
		}
		return attenuation.Mul(color(scattered, h, depth+1))
	}
	t := 0.5 * (r.Dir.ToUnit().Y() + 1)
	white := geom.NewVec(1, 1, 1).Scale(1 - t)
	blue := geom.NewVec(0.5, 0.7, 1).Scale(t)
	return white.Add(blue)
}
