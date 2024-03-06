// ---------------------------------------------------------
// EXERCISE: Declare with bits
//
//  1. Declare a few variables using the following types
//    int
//    int8
//    int16
//    int32
//    int64
//    float32
//    float64
//    complex64
//    complex128
//    bool
//    string
//    rune
//    byte
//
// 2. Observe their output
//
// 3. After you've done, check out the solution
//    and read the comments there
//
// EXPECTED OUTPUT
//  0 0 0 0 0 0 0 (0+0i) (0+0i) false 0 0
//  ""
// ---------------------------------------------------------

package exercise_6

import "fmt"

func WithBits() {
	var firstNumInt8 int8
	var firstNumInt16 int16
	var firstNumInt32 int32
	var firstNumInt64 int64
	var firstNumFloat32 float32
	var firstNumFloat64 float64
	var firstNumComplex64 complex64
	var firstNumComplex128 complex128
	var firstNumBool bool
	var firstNumString string
	var firstNumRune rune
	var firstNumByte byte

	fmt.Println("int8:", firstNumInt8)
	fmt.Println("int16:", firstNumInt16)
	fmt.Println("int32:", firstNumInt32)
	fmt.Println("int64:", firstNumInt64)
	fmt.Println("float32:", firstNumFloat32)
	fmt.Println("float64:", firstNumFloat64)
	fmt.Println("complex64:", firstNumComplex64)
	fmt.Println("complex128:", firstNumComplex128)
	fmt.Println("bool:", firstNumBool)
	fmt.Println("string:", firstNumString)
	fmt.Println("rune:", firstNumRune)
	fmt.Println("byte:", firstNumByte)

}
