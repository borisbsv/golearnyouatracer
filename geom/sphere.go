package geom

import (
	"math"
)

type Sphere struct {
	center Vec
	radius float64
}

func NewSphere(c Vec, r float64) Sphere {
	return Sphere{center: c, radius: r}
}

func (s Sphere) Hit(r Ray, tMin, tMax float64) (float64, Vec, Vec) {
	oc := r.Or.Sub(s.center)
	a := r.Dir.Dot(r.Dir)
	b := oc.Dot(r.Dir)
	c := oc.Dot(oc) - s.radius*s.radius
	discriminant := b*b - a*c

	if discriminant <= 0 {
		return 0, Vec{}, Vec{}
	}

	var p Vec
	t := (-b - math.Sqrt(b*b-a*c)) / a
	if tMax > t && t > tMin {
		p = r.AtParam(t)
		return t, p, p.Sub(s.center).ToUnit()
	}

	t = (-b + math.Sqrt(b*b-a*c)) / a
	if tMax > t && t > tMin {
		p = r.AtParam(t)
		return t, p, p.Sub(s.center).ToUnit()
	}
	return 0, Vec{}, Vec{}
}
