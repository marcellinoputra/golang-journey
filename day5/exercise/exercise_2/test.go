package exercise_2

import "fmt"

// EXERCISE: Print hexes
//
//  1. Print 0 to 9 in hexadecimal
//
//  2. Print 10 to 15 in hexadecimal
//
//  3. Print 17 in hexadecimal
//
//  4. Print 25 in hexadecimal
//
//  5. Print 50 in hexadecimal
//
//  6. Print 100 in hexadecimal

func ZeroToNine() {
	for num := 0x0; num <= 0x9; num++ {
		fmt.Println(num)
	}
}

func TenToFifTeen() {
	for num := 0xa; num <= 0xf; num++ {
		fmt.Println(num)
	}
}

func SevenTeen() {
	fmt.Println(0x11)
}

func TwentyFive() {
	fmt.Println(0x19)
}

func Fifty(){
	fmt.Println(0x32)
}

func Hundred() {
	fmt.Println(0x64)

}
