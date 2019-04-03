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
	for j := ny - 1; j >= 0; j-- {
		for i := 0.0; i < nx; i++ {
			c := geom.NewVec(i/nx, j/ny, 0.2)
			ir := int(255.99 * c.R())
			ig := int(255.99 * c.G())
			ib := int(255.99 * c.B())
			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}
}
