package main

import (
	"fmt"
	"math"
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
	cam := draw.NewCamera(90, x/y)
	// l := draw.NewList(
	//     geom.NewSphere(geom.NewVec(0, 0, -1), 0.5, &material.Lambertian{Albedo: geom.NewVec(0.1, 0.2, 0.5)}),
	//     geom.NewSphere(geom.NewVec(0, -100.5, -1), 100, &material.Lambertian{Albedo: geom.NewVec(0.8, 0.8, 0)}),
	//     geom.NewSphere(geom.NewVec(1, 0, -1), 0.5, material.NewMetal(geom.NewVec(0.8, 0.6, 0.2), 0.3)),
	//     geom.NewSphere(geom.NewVec(-1, 0, -1), 0.5, material.NewGlass(1.5)),
	//     geom.NewSphere(geom.NewVec(-1, 0, -1), -0.45, material.NewGlass(1.5)),
	// )

	r := math.Cos(math.Pi / 4)
	l := draw.NewList(
		geom.NewSphere(geom.NewVec(-r, 0, -1), r, &material.Lambertian{Albedo: geom.NewVec(0, 0, 1)}),
		geom.NewSphere(geom.NewVec(r, 0, -1), r, &material.Lambertian{Albedo: geom.NewVec(1, 0, 0)}),
	)

	fmt.Println(scene.WritePPM(f, l, 100, cam))
}
