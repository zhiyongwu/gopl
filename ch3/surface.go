package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func cornor(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy

}

func plot() string {
	var result string
	result += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := cornor(i+1, j)
			bx, by := cornor(i, j)
			cx, cy := cornor(i, j+1)
			dx, dy := cornor(i+1, j+1)
			result += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}

	}
	result += "</svg>"
	return result
}
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		result := plot()
		writer.Header().Set("Content-Type", "image/svg+xml")
		io.Copy(writer, strings.NewReader(result))
	})
	http.ListenAndServe("localhost:8080", nil)
}
