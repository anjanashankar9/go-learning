package basename2

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(basename("a"))
	fmt.Println(basename("a.go"))
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("a/b.c.go"))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}