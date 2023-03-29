package main

import "fmt"

func c() (i int) {
	defer func() { i++; fmt.Println(i) }()
	i = 1
	return i
}

func main() {
	fmt.Println(c())
}
