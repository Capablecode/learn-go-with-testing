package helperfunc

import (
	"regexp"
	"strings"
)

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
	if index > len(word) {
		index = len(word)
	}
	for i := start; i < len(word); i++ {
		word[i] = strings.ToLower(word[i])
	}
	return word
}

func Cap(word string) string {
	// return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	return strings.ToUpper(string(word))
}

func Lower(word string) string {
	return strings.ToLower(string(word))
}

func FixQuotes(text string) string {
	trimBackTick := regexp.MustCompile(`'\s*([^']*?)\s*'`)
	result := trimBackTick.ReplaceAllString(text, "'$1'")
	return result
}

// func CalculateYears(years int) (result [3]int) {
//   // Write your solution here
//   if years == 1{
//     humanYears := 1
//     catYear := 15
//     dogYear := 15
//     result[0] = humanYears
//     result[1] = catYear
//     result[2] = dogYear
//   }else if(years == 2){
//     humanYears := 2
//     catYear := 15 + 9
//     dogYear := 15 + 9
//     result[0] = humanYears
//     result[1] = catYear
//     result[2] = dogYear
//   }else{
// 	humanYears := humanYears - 2
//     catYear := 24 + (3 * 4)
//     dogYear := 24 + (3 * 5)
//     result[1] := catYear
//     result[2] := dogYear
//   }
//   return result
// }
