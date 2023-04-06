package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter

	c.Write([]byte("hello"))
	a := A
	fmt.Println(c)
	fmt.Println(a)
}

var A interface{}
