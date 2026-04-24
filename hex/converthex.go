package hex

import "strconv"

func ConvertHexidecimal(word string) string {
	val, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(val, 10)
}

func ConvertBinary(word string) string {
	val, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(val, 10)
}
