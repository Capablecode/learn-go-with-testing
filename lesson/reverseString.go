package lesson

func ReverseAString(s string) string {
	lenOfTheString := len(s) - 1
	result := lenOfTheString

	var output = ""
	for i := result; i >= 0; i-- {
		output += string(s[i])
	}
	return output

}
