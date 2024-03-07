// ---------------------------------------------------------
// EXERCISE: Redeclare
//
// 	1. Short declare two int variables: age and yourAge
//     (use multiple short declaration syntax)
//
//  2. Short declare a new float variable: ratio
//     And, change the 'age' variable to 42
//
//     (! You should use redeclaration)
//
//  4. Print all the variables
//
// EXPECTED OUTPUT
//  42, 20, 3.14
// ---------------------------------------------------------

package exercise_14

import "fmt"

func Redeclare() {
	var (
		age     int
		yourAge int
	)

	var ratio float32

	age = 42
	yourAge = 20
	ratio = 3.14

	fmt.Println(age, yourAge, ratio)
}
