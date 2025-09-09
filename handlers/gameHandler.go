package handlers

import (
	"net/http"

	"github.com/N95Ryan/8bit-hangman-back/game"
	"github.com/N95Ryan/8bit-hangman-back/models"
	"github.com/gin-gonic/gin"
)

// Structures pour les requêtes
type CreateGameRequest struct {
	PlayerName string `json:"player_name" binding:"required,min=3,max=50"`
}

type GuessRequest struct {
	Letter string `json:"letter" binding:"required,len=1"`
}

// CreateGame crée une nouvelle partie
func CreateGame(c *gin.Context) {
	var req CreateGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Créer une nouvelle partie
	newGame := game.NewGame()

	c.JSON(http.StatusCreated, gin.H{
		"id":        newGame.ID,
		"word":      newGame.GetMaskedWord(),
		"remaining": newGame.Remaining,
		"status":    newGame.Status,
	})
}

// GetGame récupère l'état d'une partie
func GetGame(c *gin.Context) {
	id := c.Param("id")

	gameInstance, exists := game.GetGame(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        gameInstance.ID,
		"word":      gameInstance.GetMaskedWord(),
		"guesses":   gameInstance.Guesses,
		"remaining": gameInstance.Remaining,
		"status":    gameInstance.Status,
		"score":     gameInstance.Score,
	})
}

// SubmitGuess soumet une lettre pour une partie
func SubmitGuess(c *gin.Context) {
	id := c.Param("id")

	gameInstance, exists := game.GetGame(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	if gameInstance.Status != "in_progress" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Game is already completed"})
		return
	}

	var req GuessRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success := gameInstance.MakeGuess(req.Letter)

	c.JSON(http.StatusOK, gin.H{
		"success":   success,
		"word":      gameInstance.GetMaskedWord(),
		"guesses":   gameInstance.Guesses,
		"remaining": gameInstance.Remaining,
		"status":    gameInstance.Status,
		"score":     gameInstance.Score,
	})
}

// AbandonGame abandonne une partie
func AbandonGame(c *gin.Context) {
	id := c.Param("id")

	if success := game.DeleteGame(id); !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetLeaderboard récupère le classement
func GetLeaderboard(c *gin.Context) {
	leaderboard := game.GetLeaderboard(10) // Limiter à 10 entrées
	c.JSON(http.StatusOK, leaderboard)
}

// SubmitScore soumet un score au classement
func SubmitScore(c *gin.Context) {
	var req struct {
		GameID string `json:"game_id" binding:"required"`
		UserID string `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gameInstance, exists := game.GetGame(req.GameID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	user, exists := models.GetUser(req.UserID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Ajouter le score au classement
	game.AddToLeaderboard(
		user.ID,
		user.Name,
		gameInstance.Score,
		len(gameInstance.Word),
		gameInstance.Remaining,
	)

	c.Status(http.StatusCreated)
}
