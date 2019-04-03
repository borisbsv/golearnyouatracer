package geom

type Ray struct {
	Or  Vec
	Dir Vec
}

func NewRay(or, dir Vec) Ray {
	return Ray{or, dir}
}

func (r Ray) AtParam(t float64) Vec {
	return r.Or.Add(r.Dir.Scale(t))
}
