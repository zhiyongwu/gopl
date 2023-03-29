package main

import "fmt"

func f(arr ...int) int {
	var result int

	for _, a := range arr {
		result += a
	}
	return result
}

func main() {
	a := []int{1, 2, 3, 4}
	var b = []int{1, 2, 3, 5}
	fmt.Println(f(1, 2, 3, 4))
	fmt.Println(f(a...))
	fmt.Println(f(b...))

}
