package handlers

import (
	"net/http"

	"github.com/N95Ryan/8bit-hangman-back/models"
	"github.com/gin-gonic/gin"
)

// Structures pour les requêtes
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterUser enregistre un nouvel utilisateur
func RegisterUser(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Vérifier si l'utilisateur existe déjà
	if models.UserExists(req.Username) {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// Créer un nouvel utilisateur
	user, err := models.CreateUser(req.Username, req.Password, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retourner l'utilisateur créé (sans le mot de passe)
	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"username": user.Name,
		"email":    user.Email,
	})
}

// LoginUser connecte un utilisateur
func LoginUser(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Vérifier les identifiants
	user, token, err := models.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Retourner le token et les informations utilisateur
	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"id":       user.ID,
		"username": user.Name,
	})
}

// GetUserProfile récupère le profil d'un utilisateur
func GetUserProfile(c *gin.Context) {
	userID := c.Param("id")

	user, exists := models.GetUser(userID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Name,
		"email":    user.Email,
		"stats": gin.H{
			"games_played": user.GamesPlayed,
			"games_won":    user.GamesWon,
			"high_score":   user.HighScore,
		},
	})
}

// UpdateUserProfile met à jour le profil d'un utilisateur
func UpdateUserProfile(c *gin.Context) {
	userID := c.Param("id")

	user, exists := models.GetUser(userID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mettre à jour les champs si fournis
	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Password != "" {
		if err := models.UpdateUserPassword(userID, req.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if err := models.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Name,
		"email":    user.Email,
	})
}
