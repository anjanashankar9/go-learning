package main

import "fmt"

/*
Constant is a package level declaration (package main)
The name of each package level entity is visible not only
throughout the source file that contains its declaration,
but throughout all the files of the package.
 */
const boilingF = 212.0

func main() {
	/*
	In contrast to package declarations, local declarations
	are visible only within the scope in which they are
	declared.
	 */
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)

}
