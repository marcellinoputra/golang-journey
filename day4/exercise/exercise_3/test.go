package exercise_3

import (
	"fmt"
	"runtime"
)

func CheckGoVersion() {
	version := runtime.Version()

	fmt.Println(version)

}
