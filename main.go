package main

import (
	"fmt"
	"os"

	"github.com/templarrei/golearnyouatracer/draw"
	"github.com/templarrei/golearnyouatracer/geom"
	"github.com/templarrei/golearnyouatracer/material"
)

func main() {
	f, _ := os.Create("test.ppm")
	defer f.Close()

	const x, y float64 = 200, 100

	scene := draw.NewScene(x, y)
	cam := draw.Camera{
		LLCorner:   geom.NewVec(-2, -1, -1),
		Horizontal: geom.NewVec(4, 0, 0),
		Vertical:   geom.NewVec(0, 2, 0),
		Origin:     geom.NewVec(0, 0, 0),
	}

	l := draw.NewList(
		geom.NewSphere(geom.NewVec(0, 0, -1), 0.5, &material.Lambertian{Albedo: geom.NewVec(0.1, 0.2, 0.5)}),
		geom.NewSphere(geom.NewVec(0, -100.5, -1), 100, &material.Lambertian{Albedo: geom.NewVec(0.8, 0.8, 0)}),
		geom.NewSphere(geom.NewVec(1, 0, -1), 0.5, material.NewMetal(geom.NewVec(0.8, 0.6, 0.2), 0.3)),
		geom.NewSphere(geom.NewVec(-1, 0, -1), 0.5, material.NewGlass(1.5)),
		geom.NewSphere(geom.NewVec(-1, 0, -1), -0.45, material.NewGlass(1.5)),
	)

	fmt.Println(scene.WritePPM(f, l, 100, cam))
}
