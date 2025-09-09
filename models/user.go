package models

import (
	"errors"
	"sync"
	"time"

	"github.com/N95Ryan/8bit-hangman-back/utils"
	"golang.org/x/crypto/bcrypt"
)

// User représente un utilisateur du jeu
type User struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"-"` // Le mot de passe n'est jamais exposé en JSON
	GamesPlayed int       `json:"games_played"`
	GamesWon    int       `json:"games_won"`
	HighScore   int       `json:"high_score"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Stockage en mémoire des utilisateurs
var (
	users       = make(map[string]*User)
	usersByName = make(map[string]string) // map[username]userID
	usersMutex  sync.RWMutex
)

// Stockage des tokens d'authentification
var (
	authTokens  = make(map[string]string) // map[token]userID
	tokensMutex sync.RWMutex
)

// CreateUser crée un nouvel utilisateur
func CreateUser(username, password, email string) (*User, error) {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	// Vérifier si le nom d'utilisateur existe déjà
	if _, exists := usersByName[username]; exists {
		return nil, errors.New("username already exists")
	}

	// Hasher le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Créer l'utilisateur
	user := &User{
		ID:        utils.GenerateID(),
		Name:      username,
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Stocker l'utilisateur
	users[user.ID] = user
	usersByName[username] = user.ID

	return user, nil
}

// GetUser récupère un utilisateur par son ID
func GetUser(id string) (*User, bool) {
	usersMutex.RLock()
	defer usersMutex.RUnlock()

	user, exists := users[id]
	return user, exists
}

// UserExists vérifie si un utilisateur existe par son nom
func UserExists(username string) bool {
	usersMutex.RLock()
	defer usersMutex.RUnlock()

	_, exists := usersByName[username]
	return exists
}

// AuthenticateUser authentifie un utilisateur et génère un token
func AuthenticateUser(username, password string) (*User, string, error) {
	usersMutex.RLock()
	userID, exists := usersByName[username]
	usersMutex.RUnlock()

	if !exists {
		return nil, "", errors.New("user not found")
	}

	user := users[userID]

	// Vérifier le mot de passe
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("invalid password")
	}

	// Générer un token
	token := utils.GenerateID()

	// Stocker le token
	tokensMutex.Lock()
	authTokens[token] = user.ID
	tokensMutex.Unlock()

	return user, token, nil
}

// UpdateUser met à jour les informations d'un utilisateur
func UpdateUser(user *User) error {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	if _, exists := users[user.ID]; !exists {
		return errors.New("user not found")
	}

	user.UpdatedAt = time.Now()
	users[user.ID] = user

	return nil
}

// UpdateUserPassword met à jour le mot de passe d'un utilisateur
func UpdateUserPassword(userID, newPassword string) error {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	user, exists := users[userID]
	if !exists {
		return errors.New("user not found")
	}

	// Hasher le nouveau mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.UpdatedAt = time.Now()

	return nil
}

// UpdateUserStats met à jour les statistiques d'un utilisateur
func UpdateUserStats(userID string, won bool, score int) error {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	user, exists := users[userID]
	if !exists {
		return errors.New("user not found")
	}

	user.GamesPlayed++
	if won {
		user.GamesWon++
	}

	if score > user.HighScore {
		user.HighScore = score
	}

	user.UpdatedAt = time.Now()
	return nil
}

// ValidateToken vérifie si un token est valide et retourne l'ID utilisateur associé
func ValidateToken(token string) (string, bool) {
	tokensMutex.RLock()
	defer tokensMutex.RUnlock()

	userID, exists := authTokens[token]
	return userID, exists
}
