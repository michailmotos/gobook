//3.4 computes an SVG rendering of a 3-D surface function & writes SVG data to a client
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		handler(w)
	}) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//handler serves the SVG content to the client
func handler(w http.ResponseWriter) {
	var disp string
	disp = fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	w.Write([]byte(disp))
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
			if red {
				disp = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:#ff0000'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
				w.Write([]byte(disp))
			} else {
				disp = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:#0000ff'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
				w.Write([]byte(disp))
			}

		}
	}
	disp = fmt.Sprintln("</svg>")
	w.Write([]byte(disp))
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	var color bool = false
	// Compute surface height z.
	z := f(x, y)
	if z < 0 || z > 2 {
		color = true
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, color
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
