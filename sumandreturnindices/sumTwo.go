package sumandreturnindices

func SumAndReturnIndices(num []int, target int) []int {
	count := make(map[int]int)
	
	for i, val := range num {
		needed := target - val

		if idx, found := count[needed]; found {
			return []int{idx, i}
		}
		count[val] = i
	}
	return nil
}