package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Vérifier l'authentification
		token := c.GetHeader("Authorization")

		// Pour l'instant, on accepte tout le monde
		// À remplacer par une logique d'authentification réelle
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			return
		}

		// Simuler une validation de token
		if token != "Bearer valid-token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}
