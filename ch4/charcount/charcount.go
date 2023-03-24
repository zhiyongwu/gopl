package main

import (
	"fmt"
	"strconv"
)

func main() {
	counts := map[rune]int{}
	counts['a'] = 3
	fmt.Println(counts)

	a := 'K'
	b := a + 'A' - 'a'
	fmt.Println(strconv.QuoteRune(b))
	fmt.Println(string(b))

}
