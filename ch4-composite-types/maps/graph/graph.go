package main

import "fmt"

/*
The value of type map can itself be a composite type, such as map or slice.
Following code represents a graph with key type string and value type
is a map[string] bool, representing a set of strings.
Conceptually this maps a vertex to its successors in a directed graph.
*/

var graph = make(map[string]map[string]bool)

/*
This function populates a map. It initiates each value as its key appears for the first time
*/
func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}

	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
func main() {
	addEdge("A", "B")
	addEdge("A", "C")
	addEdge("C", "D")
	fmt.Println(hasEdge("B", "C"))
	fmt.Println(hasEdge("A", "C"))
}
