package main

import (
	"fmt"
	"os"
)

/* This program prints its command line arguments */

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
	i := 1
	s, sep = "", ""
	for i < len(os.Args) {
		s += sep + os.Args[i]
		sep = " "
		i++
	}

	fmt.Println(s)

	i = 1
	s, sep = "", ""
	for {
		if i < len(os.Args) {
			s += sep + os.Args[i]
			sep = " "
		} else {
			break
		}
		i++
	}

	fmt.Println(s)

}
