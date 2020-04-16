package main

import (
	"flag"
	"fmt"
	"os"
)

var help = flag.Bool("help", false, "Show help")
var boolFlag = false
var stringFlag = "Hello There!"
var intFlag int

func main() {
	flag.BoolVar(&boolFlag, "boolFlag", false, "A boolean flag")
	flag.StringVar(&stringFlag, "stringFlag", "Hello There!", "A string flag")
	flag.IntVar(&intFlag, "intFlag", 4, "An integer flag")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("Boolean Flag is ", boolFlag)
	fmt.Println("String Flag is ", stringFlag)
	fmt.Println("Int Flag is ", intFlag)
}
