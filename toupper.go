package piscine

import "unicode"

func ToUpper(s string) string {
	result := ""
	for _, char := range s {
		result += string(unicode.ToUpper(char))
	}
	return result
}
