package main

/*
A variable or method of an object is said to be encapsulated
if it is inaccessible to clients of the object.
Encapsulation, sometimes called `information hiding`, is a key aspect of object-oriented programming.

Go has only one mechanism to control the visibility of names:
capitalized identifiers are exported from the package in which they are
defined, and uncapitalized names are not. The same mechanism that
limits access to members of a package also limits access to
the fields of a struct or the methods of a type. As a consequence, to encapsulate an object, we must make it a struct.

That’s the reason the IntSet type from the previous section was
declared as a struct type even though it has only a single field:
    type IntSet struct {
        words []uint64
	}

We could instead define IntSet as a slice type as follows,
though of course we’d have to replace each occurrence of s.words
by *s in its methods:
	type IntSet []uint64

Although this version of IntSet would be essentially equivalent,
it would allow clients from other packages to read and modify the slice
directly. Put another way, whereas the expression *s could be used in any
package, s.words may appear only in the package that defines IntSet.

Another consequence of this name-based mechanism is that the unit of
encapsulation is the package, not the type as in many other languages.
The fields of a struct type are visible to all code within the same package.
Whether the code appears in a function or a method makes no difference.

Encapsulation provides three benefits.
1. Because clients cannot directly modify the object’s variables,
one need inspect fewer statements to understand the possible values of
those variables.
2. Hiding implementation details prevents clients from depending on things
that might change, which gives the designer greater freedom to evolve the
implementation without breaking API compatibility.
3. In many cases the most important benefit, is that it prevents clients
from setting an object’s variables arbitrarily.
Because the object’s variables can be set only by functions in the same
package, the author of that package can ensure that all those functions
maintain the object’s internal invariants.

Functions that merely access or modify internal values of a type, such as
the methods of the Logger type from log package, below, are called getters
and setters. However, when naming a getter method, we usually omit the Get
prefix.
This preference for brevity extends to all methods,
not just field accessors, and to other redundant prefixes as well,
such as Fetch, Find, and Lookup.

	package log
    type Logger struct {
        flags  int
        prefix string
        // ...
	}
    func (l *Logger) Flags() int
    func (l *Logger) SetFlags(flag int)
    func (l *Logger) Prefix() string
    func (l *Logger) SetPrefix(prefix string)

Go style does not forbid exported fields. Of course, once exported,
a field cannot be unexported without an incompatible change to the API,
so the initial choice should be deliberate and should consider the
complexity of the invariants that must be maintained, the likelihood
of future changes, and the quantity of client code that would be
affected by a change.

Encapsulation is not always desirable. By revealing its representation
as an int64 number of nanoseconds, time.Duration lets us use all the usual
arithmetic and comparison operations with durations, and even to define constants of this type:

	const day = 24 * time.Hour
    fmt.Println(day.Seconds()) // "86400"
*/

func main() {
}
