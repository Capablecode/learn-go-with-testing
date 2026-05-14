package maplesson

import (
	"testing"
)

func TestSearch(t *testing.T){
	dictionary := map[string]string{"testing":"This is just a test"}
	 t.Run("Known Word", func(t *testing.T){
		got := dictionary.Search("test")
		want := "This is just a test"
		assertStrings(t, got, want)
	 })

	 t.Run("UnKnown Words", func(t *testing.T){
		_, err := dictionary.Search("unknown")
		want := "Couldn't find the word you were looking for"
		if err != nil{
			t.Fatal("Expected to get an error")
		}
		assertStrings(t, err.Error(), want)
	 })
	
}

func assertStrings(t *testing.TB, got, want string){
	t.Helper()
	if got != want{
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

	
