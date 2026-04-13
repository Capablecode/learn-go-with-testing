package sum

import (
	"slices"
	"testing"
)

// func TestSum(t *testing.T) {
// 	// numbers := [5]int{1, 2, 3, 4, 5}
// 	// got := Sum(numbers)
// 	// expected := 15

// 	// if got != expected {
// 	// 	t.Errorf("got '%d' but expected '%d', %v", expected, got, numbers)
// 	// }

// 	t.Run("Collection of 5 elements", func(t *testing.T) {
// 		numbers := []int{1, 2, 3, 4, 5}
// 		got := Sum(numbers)
// 		expected := 15

// 		if got != expected {
// 			t.Errorf("got '%d' but expected '%d', %v", got, expected, numbers)
// 		}
// 	})

// 	// t.Run("Collection of any size", func(t *testing.T) {
// 	// 	numbers := []int{1, 2, 3}
// 	// 	got := Sum(numbers)
// 	// 	expected := 6

// 	// 	if got != expected {
// 	// 		t.Errorf("got '%d' but expected '%d', %v", got, expected, numbers)
// 	// 	}
// 	// })

// }

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !slices.Equal(got, expected) {
		t.Errorf("got '%v' but expected '%v'", got, expected)
	}
}
