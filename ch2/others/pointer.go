package main

import "fmt"

func main() {
	a := f()
	fmt.Println(a)
	fmt.Println(a.a)
	b := *a
	fmt.Println(b)
	fmt.Println(b.a)

}

func f() *S {
	s := S{"s", 1}
	return &s
}

type S struct {
	a string
	b int
}
