package main

import "fmt"

/*
Inteface types express generalizations or abstractions about the
behaviors of other types. By generalizing, interfaces let us write functions
that are more flexible and adaptable because they are not tied to the
details of one particular implementation.

Go's interfaces are satisified implicitly, that is, there is no need to
declare all the interfaces that a given concrete type satisfies; simply
possessing the necessary methods is enough.
This design lets you create new interfaces that are satisfied by existing
concrete types without changing existing types. This is particularly useful
for types defined in packages that you don't control.

An interface os am abstract type. It doesn't expose the representation
of its values, or the set of basic operations they support, it reveals
only some of their methods.

The freedom to substitute one type for another that satisfies the same
interface is called substitutablity.
*/

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

/*
The *Bytecounter satisfies the io.Writer contract, we can pass it
to Fprintf, which does its string formatting oblivious to this change.
The Bytecounter correctly accumulates the length of the result.

Another interface that is widely used is fmt.Stringer.
This is satisified by declaring a String method
*/

/*
An interface type specifies a set of methods that a concrete type must
possess to be considered an instance of that interface.

Looking farther, we find declarations of new interface types as combinations of existing ones. Here are two examples:
    type ReadWriter interface {
        Reader
		Writer
	}
    type ReadWriteCloser interface {
        Reader
        Writer
        Closer
	}

The syntax used above, which resembles struct embedding, lets us
name another interface as a shorthand for writing out all of
its methods. This is called embedding an interface.
We could have written io.ReadWriter without embedding, like this:
	type ReadWriter interface {
        Read(p []byte) (n int, err error)
        Write(p []byte) (n int, err error)
	}

or even using a mixture of the two styles:
    type ReadWriter interface {
        Read(p []byte) (n int, err error)
        Writer
	}

All three declarations have the same effect.
The order in which the methods appear is immaterial.
All that matters is the set of methods.

A type is said to satisfy an interface if it possesses all the
methods the interface requires.
There is no need to declare the relationship between a concrete
type and the interfaces it satisifies.

A concrete type may satisfy many unrelated interfaces.
*/
func main() {
	var c ByteCounter

	c.Write([]byte("hello"))

	fmt.Println(c) //5

	c = 0 //reset the counter

	var name = "Dolly"

	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) //12
}
