package game

import (
	"strings"
	"sync"

	"github.com/N95Ryan/8bit-hangman-back/utils"
)

// Game représente l'état d'une partie de pendu
type Game struct {
	ID        string   `json:"id"`
	Word      string   `json:"word"`
	Guesses   []string `json:"guesses"`
	Remaining int      `json:"remaining"`
	Status    string   `json:"status"` // "in_progress", "won", "lost"
	Score     int      `json:"score"`
}

// Stockage en mémoire des parties
var (
	games      = make(map[string]*Game)
	gamesMutex sync.RWMutex
)

// NewGame crée une nouvelle partie avec un mot aléatoire
func NewGame() *Game {
	word := GetRandomWord()
	game := &Game{
		ID:        utils.GenerateID(),
		Word:      strings.ToUpper(word),
		Guesses:   []string{},
		Remaining: 5,
		Status:    "in_progress",
		Score:     0,
	}

	// Stocker la partie en mémoire
	gamesMutex.Lock()
	games[game.ID] = game
	gamesMutex.Unlock()

	return game
}

// GetGame récupère une partie par son ID
func GetGame(id string) (*Game, bool) {
	gamesMutex.RLock()
	defer gamesMutex.RUnlock()

	game, exists := games[id]
	return game, exists
}

// DeleteGame supprime une partie
func DeleteGame(id string) bool {
	gamesMutex.Lock()
	defer gamesMutex.Unlock()

	if _, exists := games[id]; !exists {
		return false
	}

	delete(games, id)
	return true
}

// MakeGuess traite une tentative de lettre
func (g *Game) MakeGuess(letter string) bool {
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

	// Calculer le score pour cette lettre
	g.Score += CalculateScore(g.Word, letter)

	// Vérifier si le joueur a gagné
	if g.IsWon() {
		g.Status = "won"
		g.Score += CalculateBonusScore(g.Remaining)
	}

	return true
}

// IsWon vérifie si toutes les lettres du mot ont été trouvées
func (g *Game) IsWon() bool {
	for _, char := range g.Word {
		if char != ' ' && !utils.Contains(g.Guesses, string(char)) {
			return false
		}
	}
	return true
}

// GetMaskedWord retourne le mot avec les lettres non devinées masquées
func (g *Game) GetMaskedWord() string {
	masked := ""
	for _, char := range g.Word {
		if char == ' ' {
			masked += " "
		} else if utils.Contains(g.Guesses, string(char)) {
			masked += string(char)
		} else {
			masked += "_"
		}
	}
	return masked
}

// sanitizeLetter normalise une lettre (majuscule)
func sanitizeLetter(letter string) string {
	if len(letter) == 0 {
		return ""
	}
	return strings.ToUpper(string(letter[0]))
}
