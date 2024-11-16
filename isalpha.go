package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return true
	} else {
		for _, char := range s {
			if !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z') && !(char >= '0' && char < '9') {
				return false
			}
		}
		return true
	}

}
