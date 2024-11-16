package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	// Vérifier si la base est valide
	if !isValidBase(base) {
		PrintStr("NV")
		return
	}

	// Récupérer la taille de la base
	baseLen := len(base)

	// Gérer les nombres négatifs
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	// Convertir le nombre dans la base et stocker le résultat
	digits := []rune{}
	if nbr == 0 {
		digits = append(digits, rune(base[0]))
	} else {
		for nbr > 0 {
			digits = append([]rune{rune(base[nbr%baseLen])}, digits...)
			nbr /= baseLen
		}
	}

	// Afficher les chiffres
	for _, digit := range digits {
		z01.PrintRune(digit)
	}
}

// isValidBase vérifie si la base est valide
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' || seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

// PrintStr affiche une chaîne de caractères avec z01.PrintRune
func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
