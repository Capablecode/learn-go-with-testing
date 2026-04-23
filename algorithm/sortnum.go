package algorithm

// Sort a numbers of slice to print num in ascending order
func SortNumInAscendingOrder(num []int) []int {
	for i := 0; i < len(num); i++ {
		minIndex := i

		for j := i + 1; j < len(num); j++ {
			if num[j] < num[minIndex] {
				minIndex = j
			}
		}
		num[i], num[minIndex] = num[minIndex], num[i]
	}
	return num
}

// Sort a numbers of slice to print num in descending order
func SortNumInDescendingOrder(num []int) []int {
	for k := 0; k < len(num); k++ {
		maxIndex := k

		for m := k + 1; m < len(num); m++ {
			if num[m] > num[maxIndex] {
				maxIndex = m
			}
		}
		num[k], num[maxIndex] = num[maxIndex], num[k]
	}
	return num
}
