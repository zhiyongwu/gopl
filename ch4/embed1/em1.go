package main

import (
	"fmt"
	"gopl/ch4/embed"
)

func main() {
	w := embed.Wheel{
		Circle: embed.Circle{
			Point:  embed.Point{X: 1, Y: 1},
			Radius: 10,
		},
		Spokes: 0,
	}
	fmt.Println(w.X)
	w.Draw()
}
