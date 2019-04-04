package material

import "github.com/templarrei/golearnyouatracer/geom"

type Dielectric struct {
	refractionIndex float64
}

func NewDielectric(refractionIndex float64) Dielectric {
	return Dielectric{refractionIndex: refractionIndex}
}

func (d Dielectric) Scatter(in geom.Ray, p, n geom.Vec) (geom.Ray, geom.Vec, bool) {
	reflected := in.Dir.Reflect(n)

	var outwardNormal geom.Vec
	var niOverNT float64
	if in.Dir.Dot(n) > 0 {
		outwardNormal = n.Inv()
		niOverNT = d.refractionIndex
	} else {
		outwardNormal = n
		niOverNT = 1 / d.refractionIndex
	}

	attenuation := geom.NewVec(1, 1, 1)
	refracted, ok := in.Dir.Refract(outwardNormal, niOverNT)
	if ok {
		return geom.NewRay(p, refracted), attenuation, true
	}
	return geom.NewRay(p, reflected), attenuation, true
}
