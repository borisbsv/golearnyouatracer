package draw

import (
	"fmt"
	"testing"

	"github.com/templarrei/golearnyouatracer/geom"
	"github.com/templarrei/golearnyouatracer/material"
)

//  NaN Value for col geom.Vector{NaN, NaN, NaN} (old col geom.Vector{8.795573499432999, 11.6273440996598, 15.875}) on 44,86
// u: 0.433950
// v: 0.440185
//  r: {[2.999652060605133 2.9984592255218807 2.001888713872986] [-0.6870332978812014 -0.6481014087066526 -0.32855716646412225]}
//  color: [0.28049495309384076 0.36829697185630444 0.5]

func TestColoWithEdgeCase(t *testing.T) {
	col := geom.NewVec(8.795573499432999, 11.6273440996598, 15.875)
	r := geom.NewRay(geom.NewVec(2.999652060605133, 2.9984592255218807, 2.001888713872986), geom.NewVec(-0.6870332978812014, -0.6481014087066526, -0.32855716646412225))
	hittables := NewList(
		geom.NewSphere(geom.NewVec(0, 0, -1), 0.5, &material.Lambertian{Albedo: geom.NewVec(0.1, 0.2, 0.5)}),
		geom.NewSphere(geom.NewVec(0, -100.5, -1), 100, &material.Lambertian{Albedo: geom.NewVec(0.8, 0.8, 0)}),
		geom.NewSphere(geom.NewVec(1, 0, -1), 0.5, material.NewMetal(geom.NewVec(0.8, 0.6, 0.2), 0.3)),
		geom.NewSphere(geom.NewVec(-1, 0, -1), 0.5, material.NewGlass(1.5)),
		geom.NewSphere(geom.NewVec(-1, 0, -1), -0.45, material.NewGlass(1.5)),
	)
	col = col.Add(color(r, hittables, 0))
	fmt.Println(col)
	t.Fail()
}
