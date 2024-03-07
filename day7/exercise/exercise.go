package day7

import "fmt"

func Test() {
	// Exercise1()
	// Exercise2()
	// Exercise3()
	// Exercise4()
	// Exercise5()
	// Exercise6()
	// Exercise7()
	// Exercise8()
	Exersice9()
}

func Exercise1() {

	// ---------------------------------------------------------
	// EXERCISE: Make It Blue
	//
	//  1. Change `color` variable's value to "blue"
	//
	//  2. Print it
	//
	// EXPECTED OUTPUT
	//  blue
	// ---------------------------------------------------------
	color := "green"
	newColor := "blue"

	color = newColor

	fmt.Println(color)

}

func Exercise2() {
	// ---------------------------------------------------------
	// EXERCISE: Variables To Variables
	//
	//  1. Change the value of `color` variable to "dark green"
	//
	//  2. Do not assign "dark green" to `color` directly
	//
	//     Instead, use the `color` variable again
	//     while assigning to `color`
	//
	//  3. Print it
	//
	// RESTRICTIONS
	//  WRONG ANSWER, DO NOT DO THIS:
	//  `color = "dark green"`
	//
	// HINT
	//  + operator can concatenate string values
	//
	//  Example:
	//
	//  "h" + "e" + "y" returns "hey"
	//
	// EXPECTED OUTPUT
	//  dark green
	// ---------------------------------------------------------

	color := "green"
	color = "dark" + color

	fmt.Println(color)

}

func Exercise3() {
	// ---------------------------------------------------------
	// EXERCISE: Assign With Expressions
	//
	//  1. Multiply 3.14 with 2 and assign it to `n` variable
	//
	//  2. Print the `n` variable
	//
	// HINT
	//  Example: 3 * 2 = 6
	//  * is the product operator (it multiplies numbers)
	//
	// EXPECTED OUTPUT
	//  6.28
	// ---------------------------------------------------------

	// DON'T TOUCH THIS

	// Declares a new float64 variable
	// 0. means 0.0

	n := float64(0.)

	// ADD YOUR CODE BELOW

	num1 := float64(3.14)
	num2 := float64(2)

	sum := num1 * num2

	fmt.Println(sum, n)
}

func Exercise4() {
	// ---------------------------------------------------------
	// EXERCISE: Find the Rectangle's Perimeter
	//
	//  1. Find the perimeter of a rectangle
	//     Its width is  5
	//     Its height is 6
	//
	//  2. Assign the result to the `perimeter` variable
	//
	//  3. Print the `perimeter` variable
	//
	// HINT
	//  Formula = 2 times the width and height
	//
	// EXPECTED OUTPUT
	//  22
	//
	// BONUS
	//  Find more formulas here and calculate them in new programs
	//  https://www.mathsisfun.com/area.html
	// ---------------------------------------------------------

	// UNCOMMENT THE CODE BELOW:

	var (
		perimeter     int
		width, height = 5, 6
	)

	perimeter = width*2 + height*2

	// USE THE VARIABLES ABOVE WHEN CALCULATING YOUR RESULT

	// ADD YOUR CODE BELOW

	fmt.Println(perimeter)
}

func Exercise5() {
	// ---------------------------------------------------------
	// EXERCISE: Multi Assign
	//
	//  1. Assign "go" to `lang` variable
	//     and assign 2 to `version` variable
	//     using a multiple assignment statement
	//
	//  2. Print the variables
	//
	// EXPECTED OUTPUT
	//  go version 2
	// ---------------------------------------------------------

	var (
		lang    string
		version int
	)

	lang = "go"
	version = 2

	fmt.Println(lang, "version", version)
}

func Exercise6() {
	// ---------------------------------------------------------
	// EXERCISE: Multi Assign #2
	//
	//  1. Assign the correct values to the variables
	//     to match to the EXPECTED OUTPUT below
	//
	//  2. Print the variables
	//
	// HINT
	//  Use multiple Println calls to print each sentence.
	//
	// EXPECTED OUTPUT
	//  Air is good on Mars
	//  It's true
	//  It is 19.5 degrees
	// ---------------------------------------------------------

	var (
		planet string
		isBool bool
		temp   float64
	)

	planet = "Mars"
	isBool = true
	temp = 19.5

	fmt.Println("Air is Good on", planet)
	fmt.Println("It's", isBool)
	fmt.Println("It is", temp)
}

func Exercise7() {
	// ---------------------------------------------------------
	// EXERCISE: Multi Short Func
	//
	//  1. Declare two variables using multiple short declaration syntax
	//
	//  2. Initialize the variables using `multi` function below
	//
	//  3. Discard the 1st variable's value in the declaration
	//
	//  4. Print only the 2nd variable
	//
	// NOTE
	//
	//	You should use `multi` function
	//	while initializing the variables
	//
	// EXPECTED OUTPUT
	//
	//	4
	//
	// ---------------------------------------------------------

	a, _ := multi()

	fmt.Println(a)

}

func multi() (int, int) {
	return 4, 5
}

func Exercise8() {
	// ---------------------------------------------------------
	// EXERCISE: Swapper
	//
	//  1. Change `color` to "orange"
	//     and `color2` to "green" at the same time
	//
	//     (use multiple-assignment)
	//
	//  2. Print the variables
	//
	// EXPECTED OUTPUT
	//  orange green
	// ---------------------------------------------------------

	color, color2 := "red", "blue"

	color, color2 = "orange", "green"

	fmt.Println(color, color2)
}

func Exersice9() {

	// ---------------------------------------------------------
	// EXERCISE: Swapper #2
	//
	//  1. Swap the values of `red` and `blue` variables
	//
	//  2. Print them
	//
	// EXPECTED OUTPUT
	//  blue red
	// ---------------------------------------------------------

	red, blue := "red", "blue"

	red, blue = blue, red

	fmt.Println(red, blue)
}
