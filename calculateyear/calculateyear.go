// package calculateyear

package main

import (
	"fmt"
)

func CalculateYears(years int) (result [3]int) {
	// Write your solution here
	if years == 1 {
		humanYears := years
		catYear := 15
		dogYear := 15
		result[0] = humanYears
		result[1] = catYear
		result[2] = dogYear
	}
	if years == 2 {
		humanYears := years
		catYear := 15 + 9
		dogYear := 15 + 9
		result[0] = humanYears
		result[1] = catYear
		result[2] = dogYear
	}
	if years > 2 {
		humanYears := years
		catYear := 24 + (years-2)*4
		dogYear := 24 + (years-2)*5
		result[0] = humanYears
		result[1] = catYear
		result[2] = dogYear
	}

	return result
}

func main() {
	fmt.Println(CalculateYears(5))
}
