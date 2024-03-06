package exercise_3

import "fmt"

func Syntax() {
	var speed int
	fmt.Println(speed)
}

func NamingRules() {
	// CORRECT DECLARATIONS
	var speed int
	var SpeeD int

	// underscore is allowed but not recommended
	var _speed int

	// Unicode letters are allowed
	var 速度 int

	// keep the compiler happy
	_, _, _, _ = speed, SpeeD, _speed, 速度

	fmt.Println(speed)
	fmt.Println(SpeeD)
	fmt.Println(_speed)
	fmt.Println(速度)
	
}
