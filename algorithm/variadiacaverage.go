package algorithm

import (

)

//	func reverseCipher(word string, step int) string {
//		cipher := ""
//		for _, val := range word {
//			if unicode.IsLetter(val) && unicode.IsLower(val) {
//				cipher += string('a' + (val-'a'-rune(step)+26)%26)
//			}
//			if unicode.IsLetter(val) && unicode.IsUpper(val) {
//				cipher += string('A' + (val-'A'-rune(step)+26)%26)
//			}
//			// else {
//			// 	cipher += string(val)
//			// }
//		}
//		return cipher
//	}
func averageScore(dropLowest bool, scores ...int) int {
	min := scores[0]
	sum := 0
	for _, val := range scores {
		sum += val
		if val < min {
			min = val
		}
	}
	avg := 0
	lenOfScore := len(scores)
	if dropLowest {
		if len(scores) == 1 {
			return 0
		}
		avg = (sum - min) / lenOfScore
	} else {
		avg = sum / lenOfScore
	}
	return avg
}
// func main() {
// 	// fmt.Println(reverseCipher("hello", 3))
// 	fmt.Println(averageScore(false, 70, 20, 30, 10))
// }
