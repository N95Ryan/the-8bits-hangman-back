package internal

import (
	"strings"
)

func NewGame(word string) *Game {
	return &Game{
		Word:      strings.ToUpper(word),
		Guesses:   []string{},
		Remaining: 8,
		Status:    "in_progress",
	}
}

func (g *Game) Guess(letter string) bool {
	letter = sanitizeLetter(letter)

	// Vérifier si la lettre a déjà été essayée
	for _, l := range g.Guesses {
		if l == letter {
			return false
		}
	}

	g.Guesses = append(g.Guesses, letter)

	// Vérifier si la lettre est dans le mot
	if !strings.Contains(g.Word, letter) {
		g.Remaining--
		if g.Remaining <= 0 {
			g.Status = "lost"
		}
		return false
	}

	// Vérifier si le joueur a gagné
	if g.isWon() {
		g.Status = "won"
	}

	return true
}

func (g *Game) isWon() bool {
	for _, char := range g.Word {
		if char != ' ' && !contains(g.Guesses, string(char)) {
			return false
		}
	}
	return true
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func sanitizeLetter(letter string) string {
	if len(letter) == 0 {
		return ""
	}
	return strings.ToUpper(string(letter[0]))
}

func (g *Game) GetMaskedWord() string {
	masked := ""
	for _, char := range g.Word {
		if char == ' ' {
			masked += " "
		} else if contains(g.Guesses, string(char)) {
			masked += string(char)
		} else {
			masked += "_"
		}
	}
	return masked
}
