package punctuation

import "strings"

func FixPunctuation(text string) string {
	punc := ".,:;?!"
	var result strings.Builder

	//fix space around and before punctuation
	for i := 0; i < len(text); i++ {
		char := string(text[i])

		if strings.Contains(punc, char) && result.Len() > 0 && result.String()[result.Len()-1] == ' ' {
			res := result.String()
			result.Reset()

			result.WriteString(res[:len(res)-1])
		}
		result.WriteString(char)

		//Add space only if it is a text, number, and only if we are not at the end of the text
		if strings.Contains(punc, char) && i+1 < len(text) {
			next := text[i+1]
			if next >= 'A' && next <= 'Z' || next >= 'a' && next <= 'z' || next >= '0' && next <= '9' {
				result.WriteString(" ")
			}
		}
	}
	return strings.Join(strings.Fields(result.String()), " ")
}

