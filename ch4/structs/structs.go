package main

import (
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
}
