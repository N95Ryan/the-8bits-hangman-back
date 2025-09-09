package game

import (
	"math/rand"
	"time"
)

// Liste des mots disponibles pour le jeu
var wordList = []string{
	"PROGRAMMATION",
	"ORDINATEUR",
	"ALGORITHME",
	"DEVELOPPEUR",
	"LOGICIEL",
	"INTERFACE",
	"VARIABLE",
	"FONCTION",
	"SERVEUR",
	"INTERNET",
	"NAVIGATEUR",
	"JAVASCRIPT",
	"DATABASE",
	"FRAMEWORK",
	"ARCHITECTURE",
	"MICROSERVICE",
	"CONTAINER",
	"CLOUD",
	"SECURITE",
	"RESEAU",
	"PROTOCOL",
	"BACKEND",
	"FRONTEND",
	"FULLSTACK",
	"COMPILATION",
	"RECURSION",
	"ITERATION",
	"OBJET",
	"CLASSE",
	"HERITAGE",
	"POLYMORPHISME",
	"ENCAPSULATION",
	"ABSTRACTION",
	"INTERFACE",
	"MODULE",
	"PACKAGE",
	"LIBRAIRIE",
	"DEPENDANCE",
	"INJECTION",
	"MIDDLEWARE",
	"ROUTEUR",
	"ENDPOINT",
	"REQUETE",
	"REPONSE",
	"COOKIE",
	"SESSION",
	"AUTHENTIFICATION",
	"AUTORISATION",
	"CRYPTOGRAPHIE",
}

// Catégories de mots pour des niveaux de difficulté
var wordCategories = map[string][]string{
	"easy": {
		"CHAT", "CHIEN", "MAISON", "LIVRE", "TABLE",
		"PORTE", "ARBRE", "FLEUR", "SOLEIL", "LUNE",
	},
	"medium": {
		"ORDINATEUR", "INTERNET", "LOGICIEL", "SERVEUR",
		"VARIABLE", "FONCTION", "RESEAU", "SECURITE",
	},
	"hard": {
		"ARCHITECTURE", "MICROSERVICE", "POLYMORPHISME",
		"ENCAPSULATION", "RECURSION", "CRYPTOGRAPHIE",
		"AUTHENTIFICATION", "MIDDLEWARE", "INJECTION",
	},
}

// Initialiser le générateur de nombres aléatoires
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetRandomWord retourne un mot aléatoire de la liste
func GetRandomWord() string {
	return wordList[rand.Intn(len(wordList))]
}

// GetRandomWordByDifficulty retourne un mot aléatoire selon la difficulté
func GetRandomWordByDifficulty(difficulty string) string {
	words, exists := wordCategories[difficulty]
	if !exists {
		difficulty = "medium" // Difficulté par défaut
		words = wordCategories[difficulty]
	}

	return words[rand.Intn(len(words))]
}
