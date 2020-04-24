package main

import "fmt"

/*
The hash table is an unordered collection of key/value pairs in
which all keys are distinct and the value associated with a given
key can be retrieved, removed or updated using a constant number of
key comparisons on an average, no matter how large the hash table is

A map is a referece to hash table in Go.
All the keys in the given map are of the same type and
all the values in the given map are of the same type.
However, the keys and values could be of different types.

The key type must be comparable using ==
*/
func main() {
	ages1 := make(map[string]int) //mapping from strings to int
	ages1["alice"] = 26
	ages1["bob"] = 31

	ages2 := map[string]int{
		"alice": 26,
		"bob":   31,
	}

	fmt.Println(ages1)          //map[alice:26 bob:31]
	fmt.Println(ages2)          //map[alice:26 bob:31]
	fmt.Println(ages2["alice"]) //26
	delete(ages1, "alice")
	fmt.Println(ages1) //map[bob:31]

	/*
		Map operations are safe even if the element isn't in the map.
		A map lookup using a key that is not present returns the zero value
		for it's type
	*/
	ages1["charlie"]++
	// The above is equivalent to ages1["charlie"] = ages1["charlie"] + 1
	fmt.Println(ages1)

	// However the map element is not a variable and hence we cannot
	//take its address.
	// This is not possible. It results into a compile error
	// _ = &ages1["bob"]

	// Enumerating all the key/value pairsin map
	for name, age := range ages1 {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// The order of map iteration is unspecified, and different
	// implementations might use a different hash function leading
	// to a different ordering

	// Explicit sorting is required to enumerate the key/value pairs in order
	// Most operations on maps, including lookup, delete, len and range loops
	// are safe to perform on a nil map reference, since it behaves like an
	// empty map. But storing to a nil map causes panic
	var ages map[string]int

	fmt.Println(ages == nil) //true
	//ages["carol"] = 21       //panic: assignment to entry in nil map

	// Accessing a map element by subscripting always yields a value.
	// If the key is present in the map, you get the corresponding value,
	// If not, you get the zero value for the element type
	fmt.Println(ages["bob"])

	// Sometimes you might have to distinguish between a nonexistent element
	// and an element that happens to have the value zerip.
	// This can be done using a test like this
	age, ok := ages["bob"]

	if !ok {
		fmt.Println("bob is not a key in this map")
	} else {
		fmt.Println(age == 0)
	}

	// Maps cannot be compared to each other
	// The only legal comparison is with nil
	fmt.Println(equal(ages1, ages2))
	fmt.Println("End of Main")
}

/*
Function to test whether two maps contain the same keys
and the same associated values
*/
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
		// Using !ok to distinguish the missing and present but zero case.
	}

	return true
}
