package piscine

import (
	"unicode"
)

func Capitalize(s string) string {
	runes := []rune(s)
	inWord := false

	for i, char := range runes {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			if !inWord {
				runes[i] = unicode.ToUpper(char) /
				inWord = true
			} else {
				runes[i] = unicode.ToLower(char) 
			}
		} else {
			inWord = false 
		}
	}

	return string(runes)
}
