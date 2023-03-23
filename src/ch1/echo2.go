package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	var result string
	for _, each := range os.Args[1:] {
		result += " " + each
	}
	fmt.Println(result)
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}
