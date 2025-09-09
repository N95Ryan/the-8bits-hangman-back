package internal

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	games      = make(map[string]*Game)
	gamesMutex sync.RWMutex
)

func CreateGameHandler(c *gin.Context) {
	var req CreateGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game := NewGame(RandomWord())
	game.ID = uuid.New().String()

	gamesMutex.Lock()
	games[game.ID] = game
	gamesMutex.Unlock()

	c.JSON(http.StatusCreated, gin.H{
		"id":        game.ID,
		"word":      game.GetMaskedWord(),
		"remaining": game.Remaining,
		"status":    game.Status,
	})
}

func GetGameHandler(c *gin.Context) {
	id := c.Param("id")

	gamesMutex.RLock()
	game, exists := games[id]
	gamesMutex.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        game.ID,
		"word":      game.GetMaskedWord(),
		"guesses":   game.Guesses,
		"remaining": game.Remaining,
		"status":    game.Status,
	})
}

func SubmitGuessHandler(c *gin.Context) {
	id := c.Param("id")

	gamesMutex.RLock()
	game, exists := games[id]
	gamesMutex.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	if game.Status != "in_progress" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Game is already completed"})
		return
	}

	var req GuessRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success := game.Guess(req.Letter)

	c.JSON(http.StatusOK, gin.H{
		"success":   success,
		"word":      game.GetMaskedWord(),
		"guesses":   game.Guesses,
		"remaining": game.Remaining,
		"status":    game.Status,
	})
}

func AbandonGameHandler(c *gin.Context) {
	id := c.Param("id")

	gamesMutex.Lock()
	defer gamesMutex.Unlock()

	if _, exists := games[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	delete(games, id)
	c.Status(http.StatusNoContent)
}

func GetLeaderboardHandler(c *gin.Context) {
	// Pour l'instant, on retourne un classement vide
	c.JSON(http.StatusOK, []LeaderboardEntry{})
}

func SubmitScoreHandler(c *gin.Context) {
	// À implémenter plus tard
	c.Status(http.StatusCreated)
}
