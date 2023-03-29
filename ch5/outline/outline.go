package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, _ := http.Get(url)
		doc, _ := html.Parse(resp.Body)
		resp.Body.Close()
		var stack []string
		outline(stack, doc)
	}

}

func outline(stack []string, n *html.Node) {

	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func f1(s *[]string) {
	fmt.Printf("%T", *s)
}
