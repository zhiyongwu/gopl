package embed

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) Draw() {
	fmt.Println("aaaa")
}

type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
	//X      int
}

func main() {
	w := Wheel{
		Circle: Circle{
			Point:  Point{1, 1},
			Radius: 2,
		},
		Spokes: 10,
		//X:      11,
	}
	fmt.Println(w.Circle.Point.X)
	fmt.Println(w.X)
}
