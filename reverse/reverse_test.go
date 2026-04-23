package reverse

import (
	"testing"
)

func TestReverse(t *testing.T) {
	t.Run("Reverse string", func(t *testing.T) {
		got := Reverse("hello")
		expected := "olleh"

		if got != expected {
			t.Errorf("got %s but expected %s", got, expected)
		}
	})
}
