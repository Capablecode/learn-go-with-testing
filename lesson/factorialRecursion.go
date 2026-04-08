package lesson

import (
	"fmt"
)

func Factorial_Recursion(x int )(y int) {

	if x < 0{
		fmt.Println("Factorial is undefined for negative value")
	}
	if x > 0 {
		y = x * Factorial_Recursion(x-1)
	} else {
		y = 1
	}
	return
}


