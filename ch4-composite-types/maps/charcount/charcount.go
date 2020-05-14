package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

/*
Sometimes we need map whose keys are slices but
because a map's keys must be comparable, this cannot be
expressed directly. It can be done in 2 steps
First, a helper function k that maps each key to a string
with a property that k(x) == k(y) if and only if we consider
x and y equivalent.
Then we create a map whose keys are strings, applying the helper
function to each key before we access the map
*/

func main() {
	counts := make(map[rune]int) // counts of unicode characters

	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings

	invalid := 0 // count ofinvalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")

	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\nlen\tcount\n")

	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
