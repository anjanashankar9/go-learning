package main

/*
A type switch enables a multi-way branch based on the interface value's dynamic
type. The nil case matches if x == nil, and the default case matches if no
other case does. Example:
	switch x.(type) {
		case nil:      // ...
		case int, uint // ...
		case bool:     // ...
		case string:   // ...
		default:       // ...
	}

As with an ordinary switch statement, cases are considered in order, and when a
match is found, the case's body is executed.
Case order becomes significant when one or more case types are interfaces,
since there is a possibility of two cases matching.

Sometimes, the logic for a case needs access to the value extracted by the type assertion.
	switch x := x.(type) { ... }

Like a switch statement, a type switch implicitly creates a lexical block, so the dec- laration of the new variable called x does not conflict with a variable x in an outer block. Each case also implicitly creates a separate lexical block.

*/

func main() {

}
