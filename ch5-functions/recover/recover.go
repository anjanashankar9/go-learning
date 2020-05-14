package main

/*
In most cases, terminating the program in case of panic is the
right thing to do. Recover, provides you a mechanism to clean
the mess before quitting.

For example, closing the client connection.

If the built-in recover function is called with a deferred function
and the function containing the defer statement is panicking,
recover ends the current state of panic and returns the panic value.

The function that was panicking does not continue where it left off
but returns normally.
If recover is called at any other time, it has no effect and returns nil.

Example: consider the development of a parser for a language. Even when it
appears to be working well, given the complexity of its job, bugs may still
lurk in obscure corner cases. We might prefer that, instead of crashing, the
parser turns these panics into ordinary parse errors, perhaps with an extra
message exhorting the user to file a bug report.

    func Parse(input string) (s *Syntax, err error) {
        defer func() {
	    	if p := recover(); p != nil {
    	    	err = fmt.Errorf("internal error: %v", p)
			}
		}()
	}
// ...parser...

The deferred function in Parse recovers from a panic, using the panic value to
construct an error message; a fancier version might include the entire call
stack using runtime.Stack. The deferred function then assigns to the err
result, which is returned to the caller.


*/
func main() {

}
