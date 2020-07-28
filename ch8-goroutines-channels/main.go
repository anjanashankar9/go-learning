package main

import (
	"fmt"
	"time"
)

/*
In Go, each concurrently executing activity is called a goroutine.
When a program starts, its only goroutine is the one that calls the main
function, so we call it the main goroutine. New goroutines are created by the
go statement. Syntactically, a go state- ment is an ordinary function or method
call prefixed by the keyword go. A go statement causes the function to be
called in a newly created goroutine. The go statement itself completes
immediately:

Consider the following program:
After several seconds of animation, the fib(45) call returns and the main function prints its result:
	Fibonacci(45) = 1134903170

The main function then returns. When this happens, all goroutines are abruptly
terminated and the program exits. Other than by returning from main or exiting
the program, there is no programmatic way for one goroutine to stop another,
but as we will see later, there are ways to communicate with a goroutine to
request that it stop itself.

*/
func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
