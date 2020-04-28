package main

import (
	"fmt"
	"time"
)

/*
	A struct is an aggregate data type that groups together zero or more named values of arbitrary type as a single entity.
	Each value is called field.
*/

type Employee struct {
	// Fields are usually written one per line
	// Fields name precedes its type
	// Consecutive fields of same type can be combined.
	// Field order is significant to type identity, so typically you only
	// combine the declarations of related fields.
	// The name of a struct field is exported if it begins with a capital letter
	// A struct type may contain a mixture of exported and unexported fields
	// An aggregate value cannot contain itself.
	// Thus a named struct cannot define a field of same type
	// But it can define  a field of pointer to the type
	// This lets us create recursive data structures like linked lists and trees
	ID        int
	Name      string
	Address   string
	DOB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee

	// The individual fields of dilbert are accessed using dot notation.
	// Because dilbert is a variable its fields are variables too, so we may
	// assign to a field as follows

	dilbert.Salary += 5000
	position := &dilbert.Position
	*position = "Senior" + *position

	// The dot notation also works with a pointer to a struct
	var employeeOfTheMonth = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"

	// The struct type with no fields is called the empty struct, written
	// struct{}. It has size zero and carries no information.

	// Struct Literal

	type Point struct{ X, Y int }
	// A value of a struct type can be written using a struct literal that
	// specifies value for its fields.
	p := Point{1, 2}
	fmt.Println(p)

	// There are 2 forms of struct literal. One as shown above.
	// A struct value is initialized by listing some or all of the field
	// names and their corresponding values
	// This is used most often
	p1 := Point{X: 1, Y: 2}
	fmt.Println(p1)
	// If field is omitted in this kind of literal, it is set to the zero value
	// for its type. Because names are provided the order of fields does not
	// matter.
	// 2 forms cannot be mixed.

	// Structs valies can be passed as arguments to functions and returned from
	// them.
	// For efficiency, larger struct types are usually passed to or returned from functions indirectly using pointer.

	// If all the fields of a struct are comparable, the struct itself is
	// comparable. They can be compared using == or !=
	// Comparable structs may be used as the key type of a map.

	// Go lets us use one named struct type as an anonymous field of another
	// struct type
	type Circle struct {
		X, Y, Radius int
	}

	type Wheel struct {
		X, Y, Radius, Spokes int
	}

	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	// As you see a wheel has all features of a citrcle plus Spokes.
	// So the definition of Wheel above and Wheel1 below is equivalent.
	type Wheel1 struct {
		Circle Circle
		Spokes int
	}

	var w1 Wheel1
	w1.Circle.X = 8
	w1.Circle.Y = 8
	w1.Circle.Radius = 5
	w1.Spokes = 20

	// As you can see accessing the fields of wheel becomes more cumbersome.
	// Imaging a multi hierarchy struct definition.
	// This is where anonymous fields play an important role.
	// These are fields with a type but no name.
	// We can define Wheel as below

	type Wheel2 struct {
		Circle
		Spokes int
	}

	var w2 Wheel2
	w2.X = 8
	w2.Y = 8
	w2.Radius = 5
	w2.Spokes = 20

	// There is no such shorthand for the struct literal syntax.
	// The following statements will not compile.

	//w2 = Wheel2{X: 8, Y: 8, Radius: 5, Spokes: 20}
	//w2 = Wheel2{8, 8, 5, 20}
	// Correct struct literal
	w2 = Wheel2{Circle{8, 8, 5}, 20}
	// or
	w2 = Wheel2{
		Circle: Circle{
			X:      8,
			Y:      8,
			Radius: 5,
		},
		Spokes: 20,
	}

	// You cannot have two anonymous fields of same type in a struct.
	// As the anonymous fields have implicit names.
	// Just as the name of the field is implicitly determined by its type, so
	// too is the visibility of the field.
	fmt.Println("End of Main")
}
