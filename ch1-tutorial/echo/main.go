package main

import (
	"os"
	"fmt"
	"strings"
)

/* This program prints its command line arguments */

func main() {
	var s, sep string

	/* Implementation 1
	Each time in the loop, the string s gets completely new contents.
	The += operator makes a new string by concatenating to the old string,
	a space character, and the next argument, then assigns the new string
	to s. The old contents of s are no longer in use,
	so they will be garbage-collected.
	This could be costly if the data involved is large
	 */
	for i:=1; i< len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	/* A simpler and more efficient solution is to use the Join function from
	the strings package.
	  */
	fmt.Println(strings.Join(os.Args[1:], " "))

}
