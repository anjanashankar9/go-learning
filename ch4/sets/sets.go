package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Go does not provide a set type, but since the keys of
map are distinct, a map can serve this purpose.
The function below removes duplicate lines from the given input
*/
func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true

			fmt.Println(line)
		}
	}

	err := input.Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
