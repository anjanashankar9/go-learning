package main

import "fmt"

func main() {
	// To insert comma in a non negative decimal string
	fmt.Println(comma("12345"))
	fmt.Println(comma("345"))
	fmt.Println(comma("123456"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3] + "," + s[n-3:])
}
