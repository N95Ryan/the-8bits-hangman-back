package utils

import (
	"math/rand"
	"time"
)

// Initialiser le générateur de nombres aléatoires
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateID génère un identifiant unique court (3 chiffres + 3 lettres)
func GenerateID() string {
	const digits = "0123456789"
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Générer 3 chiffres
	numPart := make([]byte, 3)
	for i := range numPart {
		numPart[i] = digits[rand.Intn(len(digits))]
	}

	// Générer 3 lettres
	letterPart := make([]byte, 3)
	for i := range letterPart {
		letterPart[i] = letters[rand.Intn(len(letters))]
	}

	// Concaténer les deux parties
	return string(numPart) + string(letterPart)
}

// Contains vérifie si un élément est présent dans une slice de strings
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// ShuffleStrings mélange une slice de strings
func ShuffleStrings(slice []string) []string {
	result := make([]string, len(slice))
	copy(result, slice)

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result
}

// RandomElement retourne un élément aléatoire d'une slice
func RandomElement(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

// TruncateString tronque une chaîne à la longueur spécifiée
func TruncateString(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}

	return s[:maxLength-3] + "..."
}

// SanitizeString nettoie une chaîne (supprime les caractères spéciaux)
func SanitizeString(s string) string {
	// À implémenter selon les besoins
	return s
}
