package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z := corner(i+1, j)
			bx, by, z := corner(i, j)
			cx, cy, z := corner(i, j+1)
			dx, dy, z := corner(i+1, j+1)
			redness := int(10 * (z + 7))
			fmt.Println(redness)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:rgb(%d,0,%d)'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, redness, 255-redness)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return 90 * math.Sin(r) / r
}
