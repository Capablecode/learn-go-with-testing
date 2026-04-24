package article

import "strings"

func Article(word string) string {
	words := strings.Fields(word)

	for i, word := range words {
		if word == "a" || word == "A" && strings.Contains("aeiouhAEIOUH", string(words[i+1][0])) && i+1 < len(words) {
			switch word {
			case "a":
				words[i] = "an"
			case "A":
				words[i] = "An"
			}
		}
	}
	return strings.Join(words, " ")
}
