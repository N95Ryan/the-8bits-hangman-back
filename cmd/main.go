package main

import (
	"github.com/N95Ryan/8bit-hangman-back/config"
	"github.com/N95Ryan/8bit-hangman-back/internal"
	"github.com/N95Ryan/8bit-hangman-back/middleware"
	"github.com/N95Ryan/8bit-hangman-back/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialisation des utilitaires
	utils.InitRandom()

	// Chargement de la configuration
	cfg := config.LoadConfig()

	// Création du routeur Gin
	r := gin.Default()

	// Middlewares
	r.Use(middleware.RateLimitMiddleware())
	r.Use(middleware.AuthMiddleware())

	// Configuration des routes
	internal.SetupRoutes(r)

	// Démarrage du serveur
	r.Run(":" + cfg.Port)
}
