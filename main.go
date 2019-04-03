package main

import (
	"fmt"
	"os"

	"github.com/templarrei/golearnyouatracer/geom"
)

func main() {
	f, _ := os.Create("test.ppm")
	defer f.Close()

	const nx, ny float64 = 200, 100
	fmt.Fprintf(f, "P3\n%f %f\n255\n", nx, ny)

	llCorner := geom.NewVec(-2, -1, -1)
	horizontal := geom.NewVec(4, 0, 0)
	vertical := geom.NewVec(0, 2, 0)
	origin := geom.NewVec(0, 0, 0)

	for j := ny - 1; j >= 0; j-- {
		for i := 0.0; i < nx; i++ {
			u := i / nx
			v := j / ny

			r := geom.NewRay(
				origin,
				llCorner.
					Add(horizontal.Scale(u)).
					Add(vertical.Scale(v)),
			)

			c := r.Color()
			ir := int(255.99 * c.R())
			ig := int(255.99 * c.G())
			ib := int(255.99 * c.B())
			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}
}
