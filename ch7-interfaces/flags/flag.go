package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

/*
flag.Value is another standard interface. It helps us define
new notations for command-line flags.
*/

var period = flag.Duration("period", 1*time.Second, "sleep period")
var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println("Woke up")
	/*
		$ go build flag.go
		$ ./flag -period 3s
	*/

	fmt.Println(*temp)
	printType()
}

/*
It is easy to define new flag notations for our own data types.
We need only define a type that satisfies the flag.Value interface.
Below is the declaration
package flag
    // Value is the interface to the value stored in a flag.
    type Value interface {
        String() string
        Set(string) error
	}

The string method formats the flag's value to use in command-line
help messages, thus every flag.Value is also a fmt.Stringer.
The Set method parses its string argument and updates the flag value.
In effect the Set method is the inverse of the String method,
and it is good practice for them to use the same notation.
*/

type Celsius float64
type Fahrenheit float64

//*celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

func fToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)
	// The call to fmt.Sscanf parses a floating-point number (value) and a
	// string (unit) from the input s. Although one must usually check Sscanf’s
	// error result, in this case we don’t need to because if there was a
	// problem, no switch case will match.

	switch unit {
	case "C":
		f.Celsius = Celsius(value)
		return nil

	case "F":
		f.Celsius = fToC(Fahrenheit(value))
		return nil
	}

	return fmt.Errorf("Invalid temperature %q", s)
}

func (f *celsiusFlag) String() string {
	return "Just for implementing the interface"
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

/*
Conceptually, a value of an interface type or interface value,
has two components, a concrete type and a value of that type.
These are called the interface's dynamic type and dynamic value.

Two interface values are equal if both are nil, or if their
dynamic types are identical and their dynamic values are equal
according to the usual behavior of == for that type.
Because interface values are comparable, they may be used as the
keys of a map or as the operand of a switch statement.

However, if two interface values are compared and have the same
dynamic type, but that type is not comparable (a slice, for instance)
then the comparison fails with a panic.

When handling errors or during debugging, it is helpful to report
the dynamic type of an interface value.
We use the fmt pacakge's %T verb

Internally, fmt uses reflection to obtain the name of the
interface's dynamic type.
*/
func printType() {
	var w io.Writer
	fmt.Printf("%T\n", w) // "<nil"

	w = os.Stdout
	fmt.Printf("%T\n", w) // "*os.File"

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w) // "*bytes.Buffer"
}

/*
A nil interface value, which contains no value at all, is not the same as an
interface value containing a pointer that happens to be nil.
*/
