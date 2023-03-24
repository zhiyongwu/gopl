package main

import "fmt"

func main() {
	s := "hello 你"
	fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		fmt.Println(i)
	}
}
