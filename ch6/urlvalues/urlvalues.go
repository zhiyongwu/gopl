package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")
	fmt.Println(m)
	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])
}
