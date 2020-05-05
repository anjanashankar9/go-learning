package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

/*
A function declaration has a name, a list of parameters,
an optional list of results and a body

	func name(parameter-list) (result-list) {
		body
	}

The parameter list specifies the names and types of the function's parameters
which are the local variables whose values or arguments are supplied by the caller.
The result list specifies the types of the values that the function returns.
If the function returns one unnamed result or no results at all, the parantheses
are optional and usually omitted.

Leaving off the result list entirely declares a function that does not return
any value and is called only for its effects.
*/

// The following declarations are equivalent:
// func f(i, j, k int, s, t string) { /*...*/ }
// func f(i int, j int, k int, s string, t string){ /*...*/ }

/*
The type of function is called its sgnature.
Two functions have the same signature if they have the same sequence of
parameter types and the same sequence of result types.
The names of parameters and results don't affect the type, nor does whether
or not they were declared using the factored form.

Go has no concept of dafault parameter values nor any way to specify
arguments by name, so the names of parameters and results don't matter
to the caller except at documentation.

Arguments are passed by value, so the function receives a copy of each
argument.

Occasionally you may encounter a function declaration without a body, indicating
that the function is implemented in a language other than Go.
Such a declaration defines the function signature

package math
func Sin(x float64) float64 //implemented in assembly language
*/

/*
Recursion
Functions may call themselves, either directly or indirectly.
Typical Go implementations use variable-size stacks that start small
and grow as needed up to a limit on the order of a gigabyte.
This lets go users use recursion without worry about overflow.
*/

/*
A function can return more than one result.
The result of calling a multi-valued function is a tuple of values.
The caller of such a function must explicitly assignt the values to variables
if any of them are to be used.

To ignore one of the values, assign it to the blank idenrifier.

In a function with named results, the operands of a return statement may be omitted. This is called a bare return.
In the following function each of return statement is equivalent to
	return words, images, err

Bare returns are best used sparingly as they rarely make the code easier to
understand.
*/
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	_, err = html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		err = fmt.Errorf("parsing HTML : %s", err)
		return
	}

	words, images = 1, 2

	return
}

/*
Errors are an important part of a package's API or an application's
user interface and failure is just one of several expected behaviors.

A function which has failure as an expected behavior returns an additional
result, conventionally the last one.
If the failure has only one possible cause, the result is a boolean,
usually called ok.

	value, ok := cache.Lookup(key)
	if !ok {
		// ...cache[key] does not exist.
	}
More often the failure may have a variety of causes for which the caller
will need an explanation. In such cases the type of the result is error.
The built-in type error is an interface type. An error may be nil or non-nil,
nil implies success and non nil implies failure.
Usually when a function returns a non nil error, the other results are undefined and should be ignored.
A non nil error has an error message string which we obtain by calling its
Error method or print by calling
	fmt.Println(err) or fmt.Printf("%v", err)

When a function call returns an error, it's the responsibility of the caller
to check and take appropriate action.
There are a number of possibilities
1. Propagate the error as is or construct a new error message.
		return nil, err

		return nil, fmt.Errorf("parsing %s as HTML : %v", url, err)
	The fmt.Errorf function formats an error message using fmt.Sprintf
	and returns a new error value.
	Because error messages are frequently chained together,
	message strings should not be capitalized and newlines should be
	avoided.

2. Retry the errors, possibly with a delay between tries.
*/

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}

		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

/*
3. If progress is impossible the caller can print the error and stop the
program gracefully. But this should usually be done by the main program.
		if err := WaitForServer(url); err != nil {
    		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		 	os.Exit(1)
		}
OR a more convenient way to achieve the same effect
		if err := WaitForServer(url); err != nil {
        	log.Fatalf("Site is down: %v\n", err)
		}

4. Log the error and continue.

5. Ignore completely.

In Go, after checking the error, failure is usuakly dealt with before success.
*/

/*
End of File
The io package guarantees that any read failure caused by an end-of-file
condition is always reported by a distinguished error, io.EOF, which is
defined as follows:
		if err == io.EOF {
			//Finished reading the file
		}
*/

/*
Functions are first class values in Go. that is, like other values, function
values have types, and they may be assigned to variables or passed to or
returned from functions.
*/
func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func main() {
	f := square
	fmt.Println(f(3))

	f = negative
	fmt.Println(f(3))

	// The following is a compile error: can't assign f(int, int) int to f(int) int
	//f = product
	g := product
	fmt.Println(g(2, 3))

	// The zero value of a function type is nil. Calling a nil value function
	// causes a panic.
	var f1 func(int) int
	//f1(3) //would result in panic: call of nil function
	f1 = square
	fmt.Println(f1(3))

	// Function values may be compared to nil.
	// But they are not comparable so they may not be compared against each
	// other or used as keys in the map.

	/*
		Anonymous function
		Named functions can be declared only at the package level, but we can use
		a function literal to denote a function valiue without any expression.
		A function literal is written like a function declaration, but without a name
		following the func keyword. It is an expression and its value is called anonymous functions.
	*/
	s := sqaures()
	fmt.Println(s()) // 1
	fmt.Println(s()) // 4
	fmt.Println(s()) // 9
	fmt.Println(s()) // 16
	s1 := sqaures()
	fmt.Println(s1())
	/*
		A call to squares creates a local variable x and returns an
		anonymous function that, each time it is called, increments x
		and returns its square.
		A second call to squares would create a second variable x and
		return a new anonymous function.
		This demonstrates that function values are not just code but can
		maintain a state.
		The anonymous inner function can access and update the local variables
		of the enclosing function squares. These hidden variable references are
		why we classify functions as reference types and why function values
		are not comparable.
		Function values like these are implemented using a technique called CLOSURES.
	*/
}

// For demonstrating anonymous functions.
func sqaures() func() int {
	// squares returns a function that returns
	// the next square number each time it is called.
	var x int

	return func() int {
		x++
		return x * x
	}
}
