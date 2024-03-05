package exercise_3

import (
	"fmt"
	f "fmt"
	fm "fmt"
)

func Print() {

	test1 := "Hello This Is Using fmt"
	test2 := "Hello This Is Using fm as fmt"
	test3 := "Hello This Is Using f as fmt"

	fmt.Println(test1)
	fm.Println(test2)
	f.Println(test3)
}
