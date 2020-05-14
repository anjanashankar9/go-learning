package main

import (
	"fmt"
	"os"
	"runtime"
)

/*
When the go runtime detects mistakes, it panics.
During a typical panic, normal execution stops, all deferred
function calls in that goroutine are executed, and the program
crashes with a log message.
This log message includes the panic value, which is usually an error
message of some sort, and for each goroutine, a stack trace showing
the function calls that were active at the time of panic.

Not all panics come from the runtime. The built-in panic function
may be called directly, it accepts any value as an argument.
A panic is the best thing to do when some "impossible" situation
occurs.

Go's panic resembles exceptions in other languages, however, the
situations in which panic is used are quite different.
Since panic causes the program to crash, it is used for grave errors
such as logical inconsistency in the program.

When a panic occurs all deferred functions are run in the reverse order,
starting with those of the topmost function on the stack and proceeding
upto main.
*/
func main() {
	defer printStack()
	f(3)
}
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)

	f(x - 1)
	/*
				Output
				f(3)
		f(2)
		f(1)
		defer 1
		defer 2
		defer 3
		panic: runtime error: integer divide by zero

		goroutine 1 [running]:
		main.f(0x0)
		        ../github/go-learning/ch5/panic/panic.go:32 +0x1dc
		main.f(0x1)
		        ../github/go-learning/ch5/panic/panic.go:35 +0x17d
		main.f(0x2)
		        ../github/go-learning/ch5/panic/panic.go:35 +0x17d
		main.f(0x3)
		        ../github/go-learning/ch5/panic/panic.go:35 +0x17d
		main.main()
		        ../github/go-learning/ch5/panic/panic.go:29 +0x2a
		exit status 2
	*/

	/* A panic occurs during the call to f(0), causing the three deferred
	calls to fmt.Printf to run. Then the runtime terminates the program
	printing the panic message and a stack dump to the standard error stream.

	It is possible for a function to recover from a panic so that it does not terminate the program.
	*/

}

/*
For diagnostic purposes, the runtime package lets the programmer dump the stack
using the same machinery.
By deferring a call to printStack in main

goroutine 1 [running]:
     main.printStack()
         ../defer.go:20
     main.f(0)
         ../defer.go:27
     main.f(1)
         ../defer.go:29
     main.f(2)
         ../defer.go:29
     main.f(3)
         ../defer.go:29
     main.main()
		 ../defer.go:15

additional text is printed to the stdout.

runtime.Stack can print information about functions that have already
been unwound. Go's panic mechanism runs the deferred functions before it unwinds the stack.
*/
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])

}
