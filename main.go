package main

import (
	"fmt"
	"os"

	"github.com/templarrei/golearnyouatracer/draw"
	"github.com/templarrei/golearnyouatracer/geom"
)

func main() {
	f, _ := os.Create("test.ppm")
	defer f.Close()

	const x, y float64 = 200, 100

	scene := draw.NewScene(x, y)

	l := draw.NewList(
		geom.NewSphere(geom.NewVec(0, 0, -1), 0.5),
		geom.NewSphere(geom.NewVec(0, -100.5, -1), 100),
	)

	fmt.Println(scene.WritePPM(f, l))
}
