package main

import "fmt"

func main() {
	fmt.Println(fToC(100))
	fmt.Println(fToC(0))
}
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
