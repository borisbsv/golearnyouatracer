package material

import (
	"math"
	"math/rand"

	"github.com/borisbsv/golearnyouatracer/geom"
)

type Glass struct {
	refractionIndex float64
}

func NewGlass(refractionIndex float64) Glass {
	return Glass{refractionIndex: refractionIndex}
}

func (g Glass) Scatter(in geom.Ray, p, n geom.Vec) (geom.Ray, geom.Vec, bool) {
	var outwardNormal geom.Vec
	var niOverNT, cosine float64
	if in.Dir.Dot(n) > 0 {
		outwardNormal = n.Inv()
		niOverNT = g.refractionIndex
		cosine = (g.refractionIndex * in.Dir.Dot(n)) / in.Dir.Len()
	} else {
		outwardNormal = n
		niOverNT = 1 / g.refractionIndex
		cosine = -in.Dir.Dot(n) / in.Dir.Len()
	}

	refracted, ok := in.Dir.Refract(outwardNormal, niOverNT)
	var reflectProbability float64
	if ok {
		reflectProbability = shlick(cosine, g.refractionIndex)
	} else {
		reflectProbability = 1
	}

	attenuation := geom.NewVec(1, 1, 1)
	if rand.Float64() < reflectProbability {
		reflected := in.Dir.Reflect(n)
		return geom.NewRay(p, reflected), attenuation, true
	}
	return geom.NewRay(p, refracted), attenuation, true
}

func shlick(cosine, refractionIndex float64) float64 {
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 *= r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
