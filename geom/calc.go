package geom

import "math"

func hitsSphere(center Vec, r float64, ray Ray) float64 {
	oc := ray.Or.Sub(center)
	a := ray.Dir.Dot(ray.Dir)
	b := 2 * oc.Dot(ray.Dir)
	c := oc.Dot(oc) - r*r
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1
	}
	return (-b - math.Sqrt(discriminant)) / (2 * a)
}
