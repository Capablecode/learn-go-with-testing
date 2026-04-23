package algorithm

func PrintMaxNum(numSlice []int) (int, int) {
	maxValue := numSlice[0]
	minValue := numSlice[0]
	for i := 0; i < len(numSlice); i++ {
		if numSlice[i] > maxValue {
			maxValue = numSlice[i]
		}
		if numSlice[i] < minValue {
			minValue = numSlice[i]
		}
	}
	return maxValue, minValue
}
