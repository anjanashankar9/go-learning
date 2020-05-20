package main

import (
	"fmt"
	"image/color"
	"math"
)

/*
Go has support for object oriented programming.
Go's definition - An object is simply a value or
variable that has methods, and a method is a function
associated with a particular type.
An OOP program is one that uses methods to express the
properties and operations of each data structure so that
clients need not access the object's representation directly.

A method is declared with a variant of the ordinary
function declaration in which an extra parameter appears
before the function name. The parameter attaches the function
to the type of that parameter.
*/

type Point struct{ X, Y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

/*
p is called the method's receiver.
I n Go, no special name like this or self is used
for the receiver, we use receiver names just as any
parameter.
*/
func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call
	// The expression p.Distance is called a selector,
	// because it selects the appropriate Distance method
	// for the receiver p of type Point.
	// Selectors are used to select fields of struct types
	// as in p.X Since methods and fields inhabit the same
	// namespace declaring a method X on the struct type Point
	// would be ambiguous and will result in compile time error.

	var a IntList

	fmt.Println(a.Sum())
	useColoredPoint()
	methodValue()
	methodExpression()
}

/*
Since each type has its own namespace for methods, we can
use the name Distance for other methods so long as they
belong to different types.
*/
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0

	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}

	return sum
}

/*
Path is a named slice type, not a struct type, yet we can still
define methods for it.
Methods may be declared on any named type defined in the same package
so long as its underlying type is neither a pointer nor an interface.
*/

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

/*
Convention dictates that if any method of Point has a pointer
receiver, then all methods of Point should have pointer receiver,
even ones that don't strictly need it.

Named types (Point) and pointers to them (*Point) are the only
types that may appear in a receiver declaration.
Furthermore, to avoid ambiguities, method declarations are not
permitted on named types that are themselves pointer types.

	type P *int
	func (P) f() { ... } // compile error: invalid receiver type

The (*Point).ScaleBy method can be called by providing a
*Point receiver like below:

r := &Point{1,2}
r.ScaleBy(2)

OR

p := Point{1,2}
pptr := &p
pptr.ScaleBy(2)

OR

p := Point{1,2}
(&p).ScaleBy(2)

The last 2 cases are ungainly.
Go helps us here. If the receiver p is a variale of type Point
but the method requires a *Point receiver, we can use shorthand

p.ScaleBy(2)

and the compiler will perform an implicit &p on the variable.
We cannot call a *Point method on a non addressable Point receiver,
because there is no way to obtain the address of a temporary
variable.

Point{1,2}.ScaleBy(2) //compile error: can't take address of Point literal.

The compiler inserts an implicit * operation for us.
pptr.Distance(q) = (*pptr).Distance(q)

*/

/*
Nil is a valid receiver.
Just as some functions allow nil pointers as arguments,
so do some methods for their receiver,
especially if nil is a meaningful zero value of the type
as with maps and slices.

In the below function, nil represents an empty list.
When you define a type whose methods allow nil as a receiver
value, it's worth pointing this out explicitly in its
documentation comments, as we did above.
*/

// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}

	return list.Value + list.Tail.Sum()
}

/*
Struct Embedding
*/
type ColoredPoint struct {
	Point
	Color color.RGBA
}

type ColoredPoint2 struct {
	*Point
	Color color.RGBA
}

/*
ColoredPoint could be defined as a strcut of three fields, but instead we
embedded a Point to provide the X and Y fields.
Embedding lets us take a syntactiv shortcut to defining a ColorePoint
that contains all the fields of Point, plus some more.
If we want, we can select the fields of ColoredPoint that were contributed by
the embedded Point without meaning Point
*/

func useColoredPoint() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // 1
	cp.Point.Y = 2
	fmt.Println(cp.Y) // 2

	/* A similar mechanism applies to the methods of Point. We can call methods
	of the embedded Point fields using a receiver type of ColoredPoint, even though ColoredPoint has no declared methods.
	*/
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	var p = ColoredPoint{Point{1, 1}, red}

	var q = ColoredPoint{Point{5, 4}, blue}

	fmt.Println(p.Distance(q.Point)) //"5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	/*
		The methods of Point have been promoted to ColoredPoint.
		Embedding allows complex types with many methods to be built
		up by the composition of several fields, each providing
		a few methods.
	*/
	fmt.Println(p.Distance(q.Point)) //"10"
	/*
		The type of an anonymous field may be a pointer to a named type,
		in which case fields and methods are promoted indirectly from
		the pointed to object.
		Adding another level of indirection lets us share common structures
		and vary the relationships between objects dynamically.
		The declaration of ColoredPoint2 above embeds a *Point
	*/

	p2 := ColoredPoint2{&Point{1, 1}, red}
	q2 := ColoredPoint2{&Point{5, 4}, blue}

	fmt.Println(p2.Distance(*q2.Point)) //"5"
	q2.Point = p2.Point                 // p and q now share the same Point
	p2.ScaleBy(2)
	fmt.Println(*p2.Point, *q2.Point) //"{2 2} {2 2}"

	/*
		A struct type may have more than one anonymous field.

		type ColoredPoint struct {
			Point
			color.RGBA
		}

		A value of this type would have all the methods of Point,
		all the methods of RGBA, and any additional methods declared on
		ColorePoint directly.

		When the compiler resolves a selector, it first looks for
		a directly declared method, then for methods promoted once from
		embedded fields, then for methods promoted twice from embedded
		fields and so on. The compiler reports an error if the selector was
		ambiguous because two methods were promoted from the same rank

		Methods can be delared only on named types (like Point)
		and pointers to them (*Point)
		With embedding it's possible and sometimes useful for unnamed
		struct types to have methods too.
	*/

}

func methodValue() {
	/*
		Usually we select and call a method in the same expression,
		e.g- p.Distance(), but it is possible to separate these two
		though not recommended.

		The selector p.Distance yields a method value, a function that
		binds a method (Point.Distance) to a specific receiver value p.
		This function can then be invoked without a receiver value, it
		needs only the non-receiver arguments.
	*/
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q)) // 5

	/*
		Method values are useful when a package's API calls for a function
		value, and the client's desired behavior for that function to call
		a method on a specific receiver.

		For example, the function time.AfterFunc calls a function value after
		a specified delay.
		type Rocket struct { ... }
		func (r *Rocket) Launch() {}
		r := new(Rocket)
		time.AfterFunc(10 * time.Second, func() { r.Launch() })

		The method value syntax is shorter.
		time.AfterFunc(10 * time.Second, r.Launch)
	*/

}

func methodExpression() {
	/*
		Related to method value is the method Expression. When
		calling a method, as opposed to an ordinary function, we must supply
		the receiver in a special way using the selector syntax.
		A method expression, written T.f or (*T).f where T is a type,
		yields a function value with a regular first parameter taking the place
		of the receiver, so it can be called in the usual way.
	*/
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance   //method expression
	fmt.Println(distance(p, q))  //5
	fmt.Printf("%T\n", distance) //func(main.Point, main.Point) float64

	/*
		Method expressions can be useful when you need a value to represent a
		choice among several methods belonging to the same type so that you can
		call the chosen method with many different receivers.

		In the following example, the variable op represents either the addition
		or the subtraction method of type Point, and Path.TranslateBy calls it
		for each Point in the Path.
	*/

}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type PathME []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point

	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}
