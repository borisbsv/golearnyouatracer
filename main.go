package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/borisbsv/golearnyouatracer/draw"
	"github.com/borisbsv/golearnyouatracer/geom"
	"github.com/borisbsv/golearnyouatracer/material"
)

func main() {
	f, _ := os.Create("test.ppm")
	defer f.Close()

	const x, y float64 = 200, 100

	scene := draw.NewScene(x, y)
	lFrom := geom.NewVec(3, 3, 2)
	lAt := geom.NewVec(0, 0, -1)
	vUp := geom.NewVec(0, 1, 0)
	cam := draw.NewCamera(
		lFrom,
		lAt,
		vUp,
		90,
		x/y,
		lFrom.Sub(lAt).Len(),
		0.01,
	)

	l := randomScene()
	fmt.Println(scene.WritePPM(f, l, 100, cam))
}

func randomScene() draw.Hittable {
	k := draw.NewList(
		geom.NewSphere(geom.NewVec(0, -1000, 0), 1000, material.Lambertian{Albedo: geom.NewVec(0.5, 0.5, 0.5)}),
	)

	for a := -11.0; a < 11; a++ {
		for b := -11.0; b < 11; b++ {
			mat := rand.Float64()
			center := geom.NewVec(a+0.9*rand.Float64(), 0.2, b+0.9*rand.Float64())
			if center.Sub(geom.NewVec(4, 0.2, 0)).Len() > 0.9 {
				switch {
				case mat < 0.8:
					k.Append(
						geom.NewSphere(
							center,
							0.2,
							material.Lambertian{
								Albedo: geom.NewVec(rand.Float64()*rand.Float64(), rand.Float64()*rand.Float64(), rand.Float64()*rand.Float64()),
							},
						),
					)
				case mat < 0.95:
					k.Append(
						geom.NewSphere(
							center,
							0.2,
							material.NewMetal(
								geom.NewVec(0.5*(1+rand.Float64()), 0.5*(1+rand.Float64()), 0.5*rand.Float64()),
								0.3,
							),
						),
					)
				default:
					k.Append(geom.NewSphere(center, 0.2, material.NewGlass(1.5)))
				}
			}
		}
	}

	k.Append(geom.NewSphere(geom.NewVec(0, 1, 0), 1, material.NewGlass(1.5)))
	k.Append(geom.NewSphere(geom.NewVec(-4, 0, 1), 1, material.Lambertian{Albedo: geom.NewVec(0.4, 0.2, 0.1)}))
	k.Append(geom.NewSphere(geom.NewVec(4, 1, 0), 1, material.NewMetal(geom.NewVec(0.7, 0.6, 0.5), 0)))

	return k
}

func practiceScene() draw.Hittable {
	return draw.NewList(
		geom.NewSphere(geom.NewVec(0, 0, -1), 0.5, &material.Lambertian{Albedo: geom.NewVec(0.1, 0.2, 0.5)}),
		geom.NewSphere(geom.NewVec(0, -100.5, -1), 100, &material.Lambertian{Albedo: geom.NewVec(0.8, 0.8, 0)}),
		geom.NewSphere(geom.NewVec(1, 0, -1), 0.5, material.NewMetal(geom.NewVec(0.8, 0.6, 0.2), 0.3)),
		geom.NewSphere(geom.NewVec(-1, 0, -1), 0.5, material.NewGlass(1.5)),
		geom.NewSphere(geom.NewVec(-1, 0, -1), -0.45, material.NewGlass(1.5)),
	)
}
