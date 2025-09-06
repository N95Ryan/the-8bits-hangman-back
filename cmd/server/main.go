package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Constantes pour le statut du jeu
const (
	StatusPlaying = "playing"
	StatusWon     = "won"
	StatusLost    = "lost"
)

// Game représente une partie de pendu
type Game struct {
	ID             string   `json:"id"`
	Word           string   `json:"word"` // Visible uniquement côté serveur
	LettersGuessed []string `json:"lettersGuessed"`
	AttemptsLeft   int      `json:"attemptsLeft"`
	Status         string   `json:"status"`
	MaskedWord     string   `json:"maskedWord"` // Le mot masqué avec des "_" pour les lettres non devinées
}

// Structure pour la requête de devinette
type GuessRequest struct {
	Letter string `json:"letter" binding:"required"`
}

// Map pour stocker les parties en cours
var games = make(map[string]*Game)

// Liste de mots pour le jeu
var wordList = []string{
	"mario", "zelda", "link", "donkey", "pikachu",
	"sonic", "tetris", "pacman", "arcade", "retro",
	"console", "manette", "cartridge", "pixel", "sprite",
	"joystick", "galaga", "metroid", "dragon", "quest",
	"chrono", "fantasy", "megaman", "street", "fighter",
}

// Génère un mot aléatoire depuis la liste
func getRandomWord() string {
	return wordList[rand.Intn(len(wordList))]
}

// Crée un masque pour le mot (ex: "hello" -> "_____")
func createMaskedWord(word string, guessedLetters []string) string {
	result := ""
	for _, char := range word {
		if containsLetter(guessedLetters, string(char)) {
			result += string(char)
		} else {
			result += "_"
		}
	}
	return result
}

// Vérifie si une lettre est dans la liste des lettres devinées
func containsLetter(letters []string, letter string) bool {
	for _, l := range letters {
		if l == letter {
			return true
		}
	}
	return false
}

// Vérifie si toutes les lettres du mot ont été devinées
func isWordGuessed(word string, guessedLetters []string) bool {
	for _, char := range word {
		if !containsLetter(guessedLetters, string(char)) {
			return false
		}
	}
	return true
}

// Génère un ID simple de maximum 6 chiffres
func generateSimpleID() string {
	// Génère un nombre aléatoire entre 1 et 999999 (6 chiffres max)
	var id string
	for {
		idNum := rand.Intn(999999) + 1
		id = fmt.Sprintf("%d", idNum)
		// Vérifie si cet ID existe déjà
		if _, exists := games[id]; !exists {
			break
		}
	}
	return id
}

// Crée une nouvelle partie
func createNewGame(c *gin.Context) {
	// Génère un ID simple
	gameID := generateSimpleID()

	// Sélectionne un mot aléatoire
	word := getRandomWord()

	// Crée une nouvelle partie
	game := &Game{
		ID:             gameID,
		Word:           word,
		LettersGuessed: []string{},
		AttemptsLeft:   6, // Nombre standard de tentatives pour le pendu
		Status:         StatusPlaying,
		MaskedWord:     createMaskedWord(word, []string{}),
	}

	// Stocke la partie dans la map
	games[gameID] = game

	// Renvoie les détails de la nouvelle partie
	c.JSON(http.StatusCreated, game)
}

// Récupère l'état actuel d'une partie
func getGame(c *gin.Context) {
	gameID := c.Param("id")

	// Vérifie si la partie existe
	game, exists := games[gameID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partie non trouvée"})
		return
	}

	// Renvoie l'état actuel de la partie
	c.JSON(http.StatusOK, game)
}

// Traite une tentative de devinette
func makeGuess(c *gin.Context) {
	gameID := c.Param("id")

	// Vérifie si la partie existe
	game, exists := games[gameID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partie non trouvée"})
		return
	}

	// Vérifie si la partie est toujours en cours
	if game.Status != StatusPlaying {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La partie est terminée"})
		return
	}

	// Parse la requête
	var request GuessRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format de requête invalide"})
		return
	}

	// Normalise la lettre (minuscule)
	letter := strings.ToLower(request.Letter)

	// Vérifie que c'est une seule lettre
	if len(letter) != 1 || !strings.ContainsAny(letter, "abcdefghijklmnopqrstuvwxyz") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Veuillez entrer une seule lettre valide"})
		return
	}

	// Vérifie si la lettre a déjà été devinée
	if containsLetter(game.LettersGuessed, letter) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cette lettre a déjà été proposée"})
		return
	}

	// Ajoute la lettre aux lettres devinées
	game.LettersGuessed = append(game.LettersGuessed, letter)

	// Vérifie si la lettre est dans le mot
	if !strings.Contains(game.Word, letter) {
		// Lettre incorrecte, diminue le nombre de tentatives restantes
		game.AttemptsLeft--

		// Vérifie si le joueur a perdu
		if game.AttemptsLeft <= 0 {
			game.Status = StatusLost
		}
	} else {
		// Vérifie si le joueur a gagné
		if isWordGuessed(game.Word, game.LettersGuessed) {
			game.Status = StatusWon
		}
	}

	// Met à jour le mot masqué
	game.MaskedWord = createMaskedWord(game.Word, game.LettersGuessed)

	// Renvoie l'état mis à jour de la partie
	c.JSON(http.StatusOK, game)
}

func main() {
	// Initialise le générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	// Crée un routeur Gin par défaut
	r := gin.Default()

	// Configuration CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// Route de test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Routes pour le jeu de pendu
	gameRoutes := r.Group("/game")
	{
		gameRoutes.POST("/new", createNewGame)
		gameRoutes.GET("/:id", getGame)
		gameRoutes.POST("/:id/guess", makeGuess)
	}

	// Démarrage du serveur sur le port 8080
	port := ":8080"
	fmt.Printf("Serveur démarré sur http://localhost%s\n", port)
	r.Run(port)
}
