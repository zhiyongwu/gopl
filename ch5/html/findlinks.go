package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Printf("read file error : %v\n", err)
	}
	defer file.Close()
	doc, err := html.Parse(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1 : %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a)
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
