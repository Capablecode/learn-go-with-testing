package sum

// func Sum(numbers [5]int) int {
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += numbers[i]
// 	}
// 	return sum
// }

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// func Sum2(number []int) int {
// 	sum := 0
// 	for _, num := range number {
// 		sum += num
// 	}
// 	return sum
// }

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNUmbers := len(numbersToSum)
	// sum := make([]int, lengthOfNUmbers)
	// for i, numbers := range numbersToSum {
	// 	sum[i] = Sum(numbers)
	// }
	// return sum
	var sums []int
	for _, number := range numbersToSum {
		sums = append(sums, Sum(number))
	}
	return sums
}
