// ---------------------------------------------------------
// EXERCISE: Short Declare
//
//  Declare and then print four variables using
//  the short declaration statement.
//
// EXPECTED OUTPUT
//  i: 314 f: 3.14 s: Hello b: true
// ---------------------------------------------------------

package exercise_9

import "fmt"

func Unused() {
	i := 314
	f := 3.14
	s := "Hello"
	b := true

	fmt.Println("i: ", i)
	fmt.Println("f: ", f)
	fmt.Println("s: ", s)
	fmt.Println("b: ", b)
}
