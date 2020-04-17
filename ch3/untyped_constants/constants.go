package main

/*
The compiler represents these uncommitted constants with much greater numeric precision than values of basic types, and arithmetic on them is more precise than machine arithmetic; you may assume at least 256 bits of precision. There are six flavors of these uncommitted constants, called untyped boolean, untyped integer, untyped rune, untyped floating-point, untyped complex, and untyped string.
By deferring this commitment, untyped constants not only retain their higher precision until later, but they can participate in many more expressions than committed constants without requiring conversions. For example, the values ZiB and YiB in the example above are too big to store in any integer variable, but they are legitimate constants that may be used in expres- sions like this one:
     fmt.Println(YiB/ZiB) // "1024"
As another example, the floating-point constant math.
Pi may be used wherever any floating- point or complex value is needed:


     var x float32 = math.Pi
     var y float64 = math.Pi
     var z complex128 = math.Pi

If math.Pi had been committed to a specific type such as float64, the result would not be as precise, and type conversions would be required to use it when a float32 or complex128 value is wanted:
*/
