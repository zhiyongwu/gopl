package main

import (
	"fmt"
	"os"
)

func main() {
	var result, temp string
	for i := 1; i < len(os.Args); i++ {
		temp = os.Args[i]
		result += " " + temp
	}
	fmt.Println(result)
}
