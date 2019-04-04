package geom

import (
	"math"
)

type Sphere struct {
	center Vec
	radius float64

	mat Material
}

func NewSphere(c Vec, r float64, mat Material) Sphere {
	return Sphere{center: c, radius: r, mat: mat}
}

func (s Sphere) Hit(r Ray, tMin, tMax float64) (float64, Vec, Vec, Material) {
	oc := r.Or.Sub(s.center)
	a := r.Dir.Dot(r.Dir)
	b := oc.Dot(r.Dir)
	c := oc.Dot(oc) - s.radius*s.radius
	discriminant := b*b - a*c

	if discriminant <= 0 {
		return 0, Vec{}, Vec{}, s.mat
	}

	var p Vec
	t := (-b - math.Sqrt(b*b-a*c)) / a
	if tMax > t && t > tMin {
		p = r.AtParam(t)
		return t, p, p.Sub(s.center).ToUnit(), s.mat
	}

	t = (-b + math.Sqrt(b*b-a*c)) / a
	if tMax > t && t > tMin {
		p = r.AtParam(t)
		return t, p, p.Sub(s.center).ToUnit(), s.mat
	}
	return 0, Vec{}, Vec{}, s.mat
}
