package main

import (
	"bufio"
	"os"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	// This reads from standard input and writes to standard output
	counts := make(map[string] int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println(line, " ", n)
		}
	}


	/*
	The program below can either read from standard input or from a list
	of named files.
	It reads one line at a time from the file and processes it.
	*/
	counts = make(map[string] int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println(line, " ", n)
		}
	}

	/*
	The program below can either read from standard input or from a list
	of named files.
	It reads all the lines from the file into the memory and processes it.
	 */
	counts = make(map[string] int)
	files = os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			data, err := ioutil.ReadFile(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			for _, line := range strings.Split(string(data), "\n") {
				counts[line]++
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println(line, " ", n)
		}
	}

}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}