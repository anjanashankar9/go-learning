package main

import (
	"bytes"
	"fmt"
)

/*
Sets in Go are usually implemented as a map[T]bool,
where T is the element type. A set represented by a map
is very flexible but for certain problems, a specialized
representation may outperform it.

For example, in domain such as dataflow analysis where set
elements are small non-negative integers,
sets have many elements, and set operations like union and
intersection are common, a bit vector is ideal.

A bit vector uses slice of unsigned integer values.
Each bit of it represents a possible element of the set.
The set contains i if the i-th bit is set.
*/

type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

/*
Since each word has 64 bits, to locate the bit for x, we use
the quotient x/64 as the word index and the remainder x%64
as the bit index within that word.
The UnionWith operation uses the bitwise OR operator | to
compute the union 64 elements at a time.
*/

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer

	buf.WriteByte('{')

	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}

				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}

	buf.WriteByte('}')
	// bytes.Buffer is often used this way in String methods.
	return buf.String()
}

func main() {
	var x, y IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) //{1 9 144}

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) //{9 42}

	x.UnionWith(&y) //1 9 42 144}
	fmt.Println(x.String())

	fmt.Println(x.Has(9), x.Has(123)) //true false

	/*
		We declared the String and Has as methods of the pointer type
		*IntSet for consistency with the other two methods, which need
		a pointer receiver because they assign to s.words.

		As a consequence of this, IntSet value does not have a String method
		Thus, you see outputs like below
	*/
	fmt.Println(&x)         //{1 9 42 144}
	fmt.Println(x.String()) //{1 9 42 144}
	fmt.Println(x)          //{[4398046511618 0 65536]}}
	/*
		In the first case, we print an *IntSet pointer, which
		has the String method.

		In the second case, we call String() on an IntSet
		variable, the compiler inserts the implicit & operation
		giving us a pointer, which has a String method.

		In the third case, because the IntSet value does not have
		a String method, fmt.Println prints the representation of
		the struct instead.

		Making String a method of IntSet and not *IntSet, might be
		a good idea, but this is a case-by-case judgement.

	*/
}
