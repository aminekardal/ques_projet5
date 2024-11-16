package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	foundNumber := false

	for _, char := range s {
		if char >= '0' && char <= '9' {
			foundNumber = true
			result = result*10 + int(char-'0')
		} else if char == '-' && !foundNumber {
			sign = -1
		}
	}

	return result * sign
}
