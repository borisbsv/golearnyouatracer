package material

import "github.com/templarrei/golearnyouatracer/geom"

// Lambertian is a diffuse material
type Lambertian struct {
	Albedo geom.Vec
}

func (l Lambertian) Scatter(in geom.Ray, p, n geom.Vec) (geom.Ray, geom.Vec, bool) {
	target := p.Add(n).Add(geom.RandVecInSphere())
	scattered := geom.NewRay(p, target.Sub(p).ToUnit())
	return scattered, l.Albedo, true
}
