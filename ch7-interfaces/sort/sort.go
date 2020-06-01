package main

import (
	"fmt"
	"sort"
)

/*
An in-place sort algorithm needs 3 things, length of the sequence
means of comparing two elements, and a way to swap two elements
- so they are the three methods of sort.Interface

	package sort

	type Interface interface {
		Len() int
		Less(i, j int) bool // i, j are indices of sequnce elements
		Swap(i, j int)
	}

To sort any sequence, we need to define a type that implements these
three methods, then apply sort.Sort to an instance of that type.

As an example, consider sorting a slice of strings.
The new type StringSlice and its Len, Less and Swap methods are defined below:
*/

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	// We can sort a slice of strings, names, by converting the slice
	// to a StringSlice
	names := []string{"a", "b", "z", "c", "v"}
	fmt.Println(names)
	// The conversion yields a slice value with the same length,
	// capacity and unerlying array as names but with a type that
	// has three methods required for sorting.
	sort.Sort(StringSlice(names))
	fmt.Println(names)
	// Sorting a slice of strings is so common that the sort package
	// provides the StringSlice type, as well as a function called
	// Strings so that the call above can be simplified to
	names = []string{"a", "b", "z", "c", "v"}
	fmt.Println(names)
	sort.Strings(names)
	fmt.Println(names)
}
