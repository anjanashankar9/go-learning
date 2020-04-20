package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int
	/* By default the initial value of array elements
	are set to the zero value for the element type
	*/
	fmt.Printf("%d %d %d ", a[0], a[1], a[2])

	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q = [3]int{1, 2, 3}
	for _, v := range q {
		fmt.Printf("%d\n", v)
	}

	/*
		The size of an array is part of its type so [3]int and [4]int are different types. The size must be a constant expression, that is, an expression whose value can be computed as the program is being compiled.
	*/

	symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: "T"}
	fmt.Println(RMB, symbol[RMB])

	r := [...]int{99: -1}

	fmt.Println(r[0])   //0
	fmt.Println(r[99])  //-1
	fmt.Println(len(r)) //100

	/*
		If an array’s element type is comparable then the array type is comparable too, so we may directly compare two arrays of that type using the == operator, which reports whether all cor- responding elements are equal. The != operator is its negation.
	*/
	e := [2]int{1, 2}
	f := [...]int{1, 2}
	g := [2]int{1, 3}
	fmt.Println(e == f, e == g, f == g) // "true false false"
	// h := [3]int{1, 2}
	// fmt.Println(e == h) // compile error: cannot compare [2]int == [3]int

}
