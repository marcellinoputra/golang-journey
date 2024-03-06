package exercise_1

import "fmt"

func PrintSomething() {
	// EXERCISE: Print the literals
	//
	//  1. Print a few integer literals
	//
	//  2. Print a few float literals
	//
	//  3. Print true and false bool constants
	//
	//  4. Print your name using a string literal
	//
	//  5. Print a non-english sentence using a string literal

	const firstNum = 10
	const secondNum = 20

	const firstDeci = 6.11
	const secondDeci = 22.5

	const ok = true
	const bad = false

	const myName = "Marcellino Putra Nugraha"
	const secondName = "Hyerin"

	fmt.Println(firstNum)
	fmt.Println(secondNum)
	fmt.Println(firstDeci)
	fmt.Println(secondDeci)
	fmt.Println(ok)
	fmt.Println(bad)
	fmt.Println(myName)
	fmt.Println(secondName)

}
