//3.3 computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	file, err := os.Create("result.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString("<body>\n")
	file.WriteString("<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
		"width='" + strconv.Itoa(width) + "' height='" + strconv.Itoa(height) + "'>")
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, red := corner(i+1, j)
			bx, by, red := corner(i, j)
			cx, cy, red := corner(i, j+1)
			dx, dy, red := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				fmt.Println("Error: Non finite values found.")
				continue
			}
			file.WriteString("<polygon points='" + strconv.FormatFloat(ax, 'f', 6, 64) + "," +
				strconv.FormatFloat(ay, 'f', 6, 64) + strconv.FormatFloat(bx, 'f', 6, 64) +
				"," + strconv.FormatFloat(by, 'f', 6, 64) +
				strconv.FormatFloat(cx, 'f', 6, 64) + "," +
				strconv.FormatFloat(cy, 'f', 6, 64) + strconv.FormatFloat(dx, 'f', 6, 64) +
				"," + strconv.FormatFloat(dy, 'f', 6, 64) + "'/>\n")
			if red {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:#ff0000'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			} else {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:#0000ff'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}

		}
	}
	fmt.Println("</svg>")
	file.WriteString("</svg>\n</body>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	var color bool = false
	// Compute surface height z.
	z := eggbox(x, y)
	if z < 0 || z > 2 {
		color = true
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, color
}

func eggbox(x, y float64) float64 {
	return 1 + math.Sin(x)*math.Sin(x) + math.Cos(x)*math.Cos(x)
}
