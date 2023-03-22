package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	result := strings.Join(os.Args[1:], " ")
	fmt.Println(result)
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}
