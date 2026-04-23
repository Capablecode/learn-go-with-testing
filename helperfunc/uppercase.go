package helperfunc

import "strings"

func ToUpperCase(word []string, index, n int) []string {
	// word = []string{"hello", "school"

	start := index - n
	if start < 0 {
		start = 0
	}
	for i := start; i < len(word); i++ {
		word[i] = strings.ToUpper(word[i])
	}
	return word
}

func ToLowerCase(word []string, index, n int) []string {
	start := index - n
	if start < 0 {
		start = 0
	}
	for i := start; i < len(word); i++ {
		word[i] = strings.ToLower(word[i])
	}
	return word
}
