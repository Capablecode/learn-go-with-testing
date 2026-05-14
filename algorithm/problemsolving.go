package algorithm

import (
	"fmt"
)

//	func reverse(s string) string {
//		res := ""
//		for i := len(s) - 1; i >= 0; i-- {
//			res += string(s[i])
//		}
//		return res
//	}
// func isPunctuation(s string) bool {
// 	for _, val := range s {
// 		if strings.Trim(string(val), ":;.,?!()") == "" {
// 			return true
// 		}
// 	}
// 	return false
// }

// func formatHex(s string) string {
// 	word := strings.Fields(s)

// 	for i := 0; i < len(word); i++ {
// 		if word[i] == "(hex)" && i > 0 {
// 			decimal := Hex(word[i-1])
// 			word[i-1] = decimal
// 			word[i-1] = strconv.FormatInt(word[i-1], 10)

//				word = append(word[:i], word[i+1:]...)
//				i--
//			}
//		}
//		return strings.Join(word, " ")
//	}
//
//	func Hex(t string) (error, int64) {
//		val, err := strconv.ParseInt(t, 16, 64)
//		if err != nil {
//			return 0, err
//		}
//	}
//
//	func aOrAn(nextWord string) string {
//		word := strings.Fields(nextWord)
//		for _, val := range word {
//			if strings.Contains("aeiouhAEIOUH", string(val[0])) {
//				return "an"
//			}
//		}
//		return "a"
//	}
// func uppercaseLastN(words []string, n int) []string {
// 	for i := len(words) - 1; i >= len(words)-n; i-- {
// 		words[i] = strings.ToUpper(string(words[i]))
// 	}
// 	return words
// }

// func repeatStr(s string, n int) string {
// 	if n == 0 {
// 		return ""
// 	}
// 	if n < 0 {
// 		return ""
// 	}
// 	result := ""
// 	for i := 0; i < n; i++ {
// 		result += s
// 	}
// 	return result
// }

// func caesarCipher(s string, shift int) string {
// 	cipher := ""
// 	for _, val := range s {
// 		if unicode.IsLower(val) {
// 			cipher += string('a' + (val-'a'+rune(shift)+26)%26)
// 		} else if unicode.IsUpper(val) {
// 			cipher += string('A' + (val-'A'+rune(shift))%26)
// 		} else {
// 			cipher += string(val)
// 		}
// 	}
// 	return cipher
// }

// func isPalindrome(s string) bool {
// 	word := strings.ToLower(s)
// 	left := 0
// 	right := len(word) - 1

//		for left < right {
//			if word[left] != word[right] {
//				return false
//			}
//			left++
//			right--
//		}
//		return true
//	}
// func maxInSlice(nums []int) (int, error) {
// 	maxValue := nums[0]
// 	minValue := nums[0]
// 	if nums == "" {
// 		return 0, nil
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] > maxValue {
// 			maxValue = nums[i]
// 		}
// 		if nums[i] < minValue {
// 			minValue = nums[i]
// 		}
// 	}
// 	return maxValue, nil
// }

func main() {
	// fmt.Println(maxInSlice(([]int{3, 1, 4, 1, 5, 7})))
	// fmt.Println(isPalindrome("racecar"))
	// fmt.Println(caesarCipher("Khoor", 3))
	// fmt.Println(repeatStr("W's", 6700))
	// fmt.Println(uppercaseLastN([]string{"this", "is", "so", "exciting"}, 4))
	// fmt.Println(aOrAn("ELEPHANT"))
	// (formatHex("1E (hex) is great"))
	// fmt.Println(reverse("GRILLED bALLs"))
	// fmt.Print(isPunctuation("?"))
	// fmt.Println()
	fmt.Println(" step : 1 I will first loop through the string")
	fmt.Println(" step : 2 To see how i can convert the word (bin) to it decimal equivalent")
	fmt.Println(" step : 3 I will do this by writing a function that convert the number after the word (bin) to its decimal equivalent")
	fmt.Println(" step : 4 print the text")
	fmt.Println(" step : 5 It has been 2 years")
}
