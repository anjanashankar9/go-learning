package main

import "fmt"

/*
Arrays are inflexible because of their fixed size.
For this reason, other than special cases like sha256's fixed size hash,
arrays are seldom used as function parameters.
Instead we use slices.

Slices represent variable-length sequences with elements of same type.
A slice type looks like an array type without the size.

Slice is a lightweight data structure that gives access to subsequence
of the elements of an array known as the slice underlying array.

A slice has three components: a pointer, a length, and a capacity.
*/
func main() {
	months := [...]string{1: "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December"}

	fmt.Println(months[1])  //January
	fmt.Println(months[12]) //December

	fmt.Println(months[4:7]) //[April May June]

	summer := months[6:9]
	fmt.Println(summer)

	// Slicing beyond capacity causes a panic.
	// Slcing beyond length extends the slice, so the result may be
	// longer than the original
	// This will cause a panic
	// fmt.Println(summer[:20])

	fmt.Println(summer[:5]) //June July August September October]

	// Since a slice contains a pointer to an element of an array
	// passing the slice to a function permits the function to modify
	// the underlying array elements.
	reverseList := []int{1, 2, 3, 4, 5}
	reverse(reverseList)
	fmt.Println(reverseList)

	// Unlike arrays slices are not comparable, so == cannot be used to test
	// whether two slices contain the same elements.
	// bytes.Equal is highly optimized to compare two byte slices but
	// for other types it needs to be done ourself

	// zero value of slice type is nil.
	// A nil slice type has no underlying array

	/* var s []int    // len(s) == 0, s == nil
	s = nil        // len(s) == 0, s == nil
	s = []int(nil) // len(s) == 0, s == nil
	s = []int() // len(s) == 0, s != nil
	So the best way to test if a slice is empty is to check if its length is 0
	*/

	// make creates a slice of a specified type, length and capacity.
	// Capacity can be omitted.
	// When omitted it is equal to length
	// make([]T, len)
	// make([]T, len, cap)

	//append function appends to slices.
	reverseList = append(reverseList, 6)
	fmt.Println(reverseList)

	fmt.Println("End of Main")
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
