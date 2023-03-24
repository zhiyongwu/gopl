package main

import "fmt"

var graph = map[string]map[string]bool{}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = map[string]bool{}
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("a", "b")
	addEdge("c", "d")
	fmt.Println(graph)
}
