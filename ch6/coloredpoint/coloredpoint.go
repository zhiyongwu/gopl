package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-p.X, q.Y-q.Y)
}

type ColoredPoint struct {
	Point
	color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	var q1 ColoredPoint
	cp.Distance(q1.Point)
}
