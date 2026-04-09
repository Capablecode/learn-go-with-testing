package iteration

import (
	"testing"
)

func TestIteration(t *testing.T) {
	repeated := Repeated("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

