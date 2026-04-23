package sum

import (
	"reflect"
	// "slices"
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

// func TestSumAll(t *testing.T) {
// 	got := SumAll([]int{1, 2}, []int{0, 9})
// 	expected := []int{3, 9}

// 	if !slices.Equal(got, expected) {
// 		t.Errorf("got '%v' but expected '%v'", got, expected)
// 	}
// }

func TestAllTails(t *testing.T) {

	checkSum := func(t *testing.T, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v but expected %v", got, expected)
		}
	}
	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checkSum(t, got, expected)

	})

	t.Run("Safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		expected := []int{0, 9}
		checkSum(t, got, expected)
	})
}
