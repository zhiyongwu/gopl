package main

import (
	"ch2/temperature"
	"fmt"

	"os"
	"strconv"
)

func main() {

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		f := temperature.CToF(temperature.Celsius(t))
		fmt.Println(f)
	}
}
