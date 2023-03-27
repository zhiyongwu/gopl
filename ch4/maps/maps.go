package main

import (
	"fmt"
	"os"
)

func main() {

	ages := map[string]int{"a": 3, "b": 12}
	age, ok := ages["a"]
	if !ok {
		os.Exit(1)
	}
	fmt.Println(age)

}
