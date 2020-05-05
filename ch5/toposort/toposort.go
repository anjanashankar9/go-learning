package main

import (
	"fmt"
	"sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
func topoSort(m map[string][]string) []string {
	var order []string
	var keys []string

	seen := make(map[string]bool)

	var visitAll func(items []string)
	/*
				When an anonymous function requires recursion, we must first declare a
				variable, and then assign the anonymous function to that variable.
				If we combine the 2 steps in the decalration, the function literal would
		        not be within the scope of the variable, so it would have no way to call itself.

	*/
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true

				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)

	return order
}
