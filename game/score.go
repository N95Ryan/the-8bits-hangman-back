package game

import (
	"strings"
	"sync"
)

// Structure pour stocker les scores
type LeaderboardEntry struct {
	PlayerID          string `json:"player_id"`
	PlayerName        string `json:"player_name"`
	Score             int    `json:"score"`
	WordLength        int    `json:"word_length"`
	RemainingAttempts int    `json:"remaining_attempts"`
}

// Stockage en mémoire des scores
var (
	leaderboard      = []LeaderboardEntry{}
	leaderboardMutex sync.RWMutex
)

// CalculateScore calcule le score pour une lettre correcte
func CalculateScore(word string, letter string) int {
	// Points de base pour chaque occurrence de la lettre
	basePoints := 10
	occurrences := strings.Count(word, letter)

	// Bonus pour les lettres rares
	letterFrequency := map[string]int{
		"E": 1, "A": 1, "I": 1, "N": 1, "O": 1, "R": 1, "S": 1, "T": 1,
		"U": 2, "L": 2, "D": 2, "M": 2,
		"G": 3, "B": 3, "C": 3, "P": 3,
		"F": 4, "H": 4, "V": 4,
		"J": 5, "Q": 5, "K": 5, "W": 5, "X": 5, "Y": 5, "Z": 5,
	}

	rarityMultiplier := letterFrequency[letter]
	if rarityMultiplier == 0 {
		rarityMultiplier = 1
	}

	return basePoints * occurrences * rarityMultiplier
}

// CalculateBonusScore calcule le bonus de fin de partie
func CalculateBonusScore(remainingAttempts int) int {
	// Plus il reste de tentatives, plus le bonus est élevé
	return remainingAttempts * 50
}

// AddToLeaderboard ajoute un score au classement
func AddToLeaderboard(playerID string, playerName string, score int, wordLength int, remainingAttempts int) {
	entry := LeaderboardEntry{
		PlayerID:          playerID,
		PlayerName:        playerName,
		Score:             score,
		WordLength:        wordLength,
		RemainingAttempts: remainingAttempts,
	}

	leaderboardMutex.Lock()
	defer leaderboardMutex.Unlock()

	leaderboard = append(leaderboard, entry)
}

// GetLeaderboard retourne le classement des meilleurs scores
func GetLeaderboard(limit int) []LeaderboardEntry {
	leaderboardMutex.RLock()
	defer leaderboardMutex.RUnlock()

	// Copier le classement pour éviter les modifications concurrentes
	result := make([]LeaderboardEntry, len(leaderboard))
	copy(result, leaderboard)

	// Trier par score (à implémenter avec sort.Slice si nécessaire)

	// Limiter le nombre de résultats
	if limit > 0 && limit < len(result) {
		return result[:limit]
	}

	return result
}
