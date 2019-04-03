package geom

func hitsSphere(center Vec, r float64, ray Ray) bool {
	oc := ray.Or.Sub(center)
	a := ray.Dir.Dot(ray.Dir)
	b := 2 * oc.Dot(ray.Dir)
	c := oc.Dot(oc) - r*r
	discriminant := b*b - 4*a*c

	return discriminant > 0
}
