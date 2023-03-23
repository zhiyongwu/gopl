package main

import "fmt"

func main() {
	x, y := 7, 8
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	fmt.Println(x, y)
}

func f() int {
	return 3
}

func g(a int) int {
	return a
}
