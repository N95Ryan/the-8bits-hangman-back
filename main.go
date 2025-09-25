package main

import (
	"os"

	"github.com/N95Ryan/8bit-hangman-back/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Configuration du port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialisation du routeur Gin
	r := gin.Default()

	// Routes pour les jeux
	r.POST("/api/games", handlers.CreateGame)
	r.GET("/api/games/:id", handlers.GetGame)
	r.POST("/api/games/:id/guess", handlers.SubmitGuess)
	r.GET("/api/games/:id/hint", handlers.GetHint)
	r.DELETE("/api/games/:id", handlers.AbandonGame)

	// Routes pour les utilisateurs
	r.POST("/api/users/register", handlers.RegisterUser)
	r.POST("/api/users/login", handlers.LoginUser)

	// Routes pour les scores
	r.GET("/api/leaderboard", handlers.GetLeaderboard)
	r.POST("/api/leaderboard", handlers.SubmitScore)

	// Démarrage du serveur
	r.Run(":" + port)
}
