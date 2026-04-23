package lesson

import "strings"

func ReverseAString(s string) string {
	lenOfTheString := len(s) - 1
	result := lenOfTheString

	// var output = ""
	var output strings.Builder
	for i := result; i >= 0; i-- {
		output.WriteString(string(s[i]))
	}
	return output.String()
}
