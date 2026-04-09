package integer

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(5, 4)
	fmt.Println(sum)

}

func TestIntger(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Wanted '%d' but got '%d'", expected, sum)
	}

}
