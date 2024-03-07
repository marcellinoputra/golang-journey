// ---------------------------------------------------------
// EXERCISE: Short With Expression
//
// 	1. Short declare a variable named `sum`
//
//  2. Initialize it with an expression by adding 27 and 3.5
//
// EXPECTED OUTPUT
//  30.5
// ---------------------------------------------------------

package exercise_12

import "fmt"

func Sum() {
	num1 := 27
	num2 := 3.5

	sum := float32(num1) + float32(num2)

	fmt.Println(sum)
}
