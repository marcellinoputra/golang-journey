// ---------------------------------------------------------
// EXERCISE: Declare string
//
//  1. Declare a string variable
//
//  2. Print that variable
//
// EXPECTED OUTPUT
//  ""
// ---------------------------------------------------------

package exercise_4

import "fmt"

func DeclareString() {
	var name string

	fmt.Printf("s (%T): %q \n", name, name)
}
