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

func (r Ray) Color() Vec {
	uDir := r.Dir.ToUnit()
	t := 0.5 * (uDir.Y() + 1)
	return NewVec(1, 1, 1).Scale(1 - t).Add(NewVec(0.5, 0.7, 1).Scale(t))
}
