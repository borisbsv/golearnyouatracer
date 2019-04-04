package material

import "github.com/templarrei/golearnyouatracer/geom"

type Metal struct {
	Albedo geom.Vec
	fuzz   float64
}

func NewMetal(albedo geom.Vec, fuzz float64) Metal {
	if fuzz > 1 {
		fuzz = 1
	}
	return Metal{
		Albedo: albedo,
		fuzz:   fuzz,
	}
}

func (m Metal) Scatter(in geom.Ray, p, n geom.Vec) (geom.Ray, geom.Vec, bool) {
	reflected := in.Dir.ToUnit().Reflect(n)
	scattered := geom.NewRay(p, reflected.Add(geom.RandVecInSphere().Scale(m.fuzz)))

	return scattered, m.Albedo, scattered.Dir.Dot(n) > 0
}
