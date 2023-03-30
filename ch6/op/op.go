package main

import (
	"fmt"
)

func main() {
	p := Point{1, 2}
	offset := Point{1, 1}
	fmt.Println(p.translate(offset, true))
	fmt.Println(p.translate(offset, false))

}

type Point struct {
	X, Y float64
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}
func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p Point) translate(offset Point, add bool) Point {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	return op(p, offset)
}
