package reverse

func Reverse(s string) string {
	stringTorune := []rune(s)
	var reversedSlicerune []rune

	for i := len(stringTorune) - 1; i >= 0; i-- {
		reversedSlicerune = append(reversedSlicerune, stringTorune[i])
	}
	return string(reversedSlicerune)
}
