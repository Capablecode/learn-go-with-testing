package lesson

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		wanted := "Hello, Chris"
		assertCorrectMessage(t, got, wanted)

		// if got != wanted {
		// 	t.Errorf("wanted %q got %q", wanted, got)
		// }
	})

	t.Run("Say 'hello, World', when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "spanish")
		wanted := "Hello, World"
		assertCorrectMessage(t, got, wanted)

		// if got != wanted {
		// 	t.Errorf("wanted %q got %q", wanted, got)
		// }
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		wanted := "Holla, Elodie"
		assertCorrectMessage(t, got, wanted)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		wanted := "Bonjour, Elodie"
		assertCorrectMessage(t, got, wanted)
	})
}

// helper function
func assertCorrectMessage(t testing.TB, got, wanted string) {
	t.Helper()
	if got != wanted {
		t.Errorf("got %q wanted %q", got, wanted)
	}
}

// func TestHello(t *testing.T) {
// 	got := Hello("Chris")
// 	wanted := "Hello, Chris"

// 	if wanted != got {
// 		t.Errorf("got %q wanted %q", got, wanted)
// 	}
// }
