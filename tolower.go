package piscine

import "unicode"

func ToLower(s string) string {
	result := ""
	for _, char := range s {
		result += string(unicode.ToLower(char))
	}
	return result
}
