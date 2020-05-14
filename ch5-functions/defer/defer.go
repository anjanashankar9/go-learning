package main

import (
	"os"
	"sync"
)

/*
	As functions grow complex and have to handle more errors, duplication of
	cleanup logic may become a maintenance problem. Go's `defer` mechanism
	makes things simpler.

	Syntactically a defer statement is an ordinary function or method call
	prefixed by the keyword defer. The function and argument expressions
	are evaluated when the statement is executed, but the actual call is
	deferred until the function that contains the defer statement has
	finished, whether normally or abnoramally (by panicking)
	Any number of calls may be deferred, they are executed in the reverse of the order in which they were deferred.

	The right place for a defer statement that releases a resource (for
	example, close file, disconnect from database) is immediately after the resource has been acquired.
*/

//Funciton to read a file
func ReadFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

// Function to unlock a mutex

var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()

	return m[key]
}

/*
The defer statements can also be used to pair "on entry" and "on exit"
actions when debugging a complex function.

Defer statements run after return statements have updated the functions' result
variables. Because an anonymous function can access its enclosing function’s
variables, including named results, a deferred anonymous function can observe
the function’s results.

defer statement in a loop deserves extra scrutiny.
	for _, filename := range filenames {
		f, err := os.Open(filename)
        if err != nil {
		}
		return err
		defer f.Close() // NOTE: risky; could run out of file descriptors
	}

	This function can potentially run out of file descriptors since all the
	files will be closed only after all files have been processed.

	Solution is to move the loop body including the deferred statement
	into another function that is called on each iteration.

	for _, filename := range filenames {
         if err := doFile(filename); err != nil {
			return err
		}
	}
    func doFile(filename string) error {
        f, err := os.Open(filename)
        if err != nil {
			return err
		}
		defer f.Close()
	}

// ...process f...
*/
func main() {
	ReadFile("temp.txt")
	lookup("temp.txt")
}
