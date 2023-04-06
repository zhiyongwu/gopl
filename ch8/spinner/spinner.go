package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib1(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func fib1(x int) int {

	if x < 2 {
		return x
	}
	a, b := 1, 1
	for i := 0; i < x; i++ {
		a, b = b, a+b
	}
	return a

}
