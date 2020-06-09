package main

/*
A type assertion is na operation applied to an interface value.
Syntactivally it looks like x.(T), where x is an expression
of an interface type and T is a type, called the asserted type.
A type assertion checks that the dynamic type of its operand matches the
asseted type.

If the asserted type T is a concrete type, then the type assertion
checks whether x's dynamic type is identical to T.
If this check succeeds, the result of the type assertion is x's
dynamic value, whose type is T.
That is, a type assertion to a concrete type extracts the concrete
value from its operand. If the check fails, operation panics.

If the asserted type T is an interface type, then the type assertion checks
whether x's dynamic type satisfies T. If this check succeeds, the dynamic
value is not extracted, the result is still an interface value
with the same type and value components, but the result has interface type T.
That is, a type assertion to an interface type changes the type of the
expression making a different set of methods accessible but it preserves
the dynamic type an value components inside the interface value.

No matter the type that was asserted, if the operand is a nil interface value,
the type assertion fails. A type assertion to a less restricitve
interface type is rarely needed, as it behaves just like an assignment except
in nil case.

	w = rw  // io.ReadWriter is assignable to io.Writer
	w = rw.(io.Writer) // fails only if rw == nil

Often we are not sure of the dynamic type of an interface value, and we'd like
to test whether it is some particular type.
If the type assertion appears in an assignment in which two
results are expected, the operation does not panic on failure
but instead returns an additional second result, a boolean indicating
success.

	var w io.Writer = os.Stdout
	f, ok := w.(*os.File) // success: ok, f == os.Stdout
	b, ok := w.**bytes.Buffer) //failure: !ok, b == nil

The second result is conventionally assigned to a variable named ok.
The ok result is often immediately used to decide what to do next.

	if f, ok := w.(*os.File); ok {
		// ... use f ...
	}

When the operand of a type assertion is a variable, rather
than invent another name for a new local variable, you'll sometimes see the
original name reused, shadowing the original:

	if w,ok := w.(*os.File); ok {
		// ... use w ...
	}
*/

/*
Consider the set of errors returned by file operations in the os package.
I/O can fail for any number of reasons, but three kinds of failure must be
handled differently: file already exists, file not found, permission denied

A reliable approach to represent structured error values is by using
a dedicated type.
For example, the OS package defines:
PathError - to describe failures involving an operation on a file path.
LinkError - to descibe failres of operations involvind two fileath, like symlink
		and rename

*/
func main() {

}
