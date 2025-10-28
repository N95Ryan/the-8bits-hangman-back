package game

import (
	"math/rand"
	"time"
)

// Liste des mots disponibles pour le jeu
var wordList = []string{
	"MARIO", "ZELDA", "LINK", "DONKEY", "PIKACHU",
	"SONIC", "TETRIS", "PACMAN", "ARCADE", "RETRO",
	"CONSOLE", "MANETTE", "PIXEL", "SPRITE", "JOYSTICK",
	"NINTENDO", "SEGA", "ATARI", "GAMEBOY", "PLAYSTATION",
	"MEGADRIVE", "POKEMON", "FINAL", "FANTASY", "METROID",
	"CASTLEVANIA", "KIRBY", "CONTRA", "MEGAMAN", "BOMBERMAN",
}

// Structure pour stocker un mot et son indice
type WordWithHint struct {
	Word string
	Hint string
}

// Catégories de mots pour des niveaux de difficulté avec indices
var wordCategories = map[string][]WordWithHint{
	"easy": {
		{Word: "MARIO", Hint: "Famous Italian plumber"},
		{Word: "SONIC", Hint: "Fast blue hedgehog"},
		{Word: "LINK", Hint: "Hero of the Triforce"},
		{Word: "TETRIS", Hint: "Falling blocks game"},
		{Word: "PACMAN", Hint: "Yellow character eating dots"},
		{Word: "PIXEL", Hint: "Smallest unit of a digital image"},
		{Word: "ARCADE", Hint: "Video game venue"},
		{Word: "RETRO", Hint: "Old-school nostalgic style"},
		{Word: "KIRBY", Hint: "Pink ball that inhales enemies"},
		{Word: "PONG", Hint: "One of the first video games (table tennis)"},
	},
	"medium": {
		{Word: "NINTENDO", Hint: "Japanese video game company"},
		{Word: "GAMEBOY", Hint: "Monochrome handheld console"},
		{Word: "POKEMON", Hint: "Creatures to catch and train"},
		{Word: "CONSOLE", Hint: "Device dedicated to gaming"},
		{Word: "CONTROLLER", Hint: "Gaming input device"},
		{Word: "JOYSTICK", Hint: "Directional control lever"},
		{Word: "PIKACHU", Hint: "Electric mouse"},
		{Word: "DONKEY", Hint: "Famous gorilla in video games"},
		{Word: "SPRITE", Hint: "2D image integrated in a scene"},
		{Word: "MEGAMAN", Hint: "Blue robot fighting other robots"},
		{Word: "ATARI", Hint: "Pioneer of gaming consoles"},
	},
	"hard": {
		{Word: "PLAYSTATION", Hint: "Sony's gaming console"},
		{Word: "CASTLEVANIA", Hint: "Vampire hunting game"},
		{Word: "MEGADRIVE", Hint: "Sega's 16-bit console"},
		{Word: "METROID", Hint: "Space adventure with Samus Aran"},
		{Word: "BOMBERMAN", Hint: "Game about placing bombs in a maze"},
		{Word: "FINALFANTASY", Hint: "Legendary Japanese RPG series"},
		{Word: "STREETSOFRAGE", Hint: "Sega's beat'em up game"},
		{Word: "MORTALKOMBAT", Hint: "Fighting game with fatalities"},
		{Word: "RESIDENTEVIL", Hint: "Horror game with zombies"},
		{Word: "METALSLUG", Hint: "Run and gun with vehicles"},
	},
}

// Initialiser le générateur de nombres aléatoires
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetRandomWord retourne un mot aléatoire de la liste
func GetRandomWord() string {
	return wordList[rand.Intn(len(wordList))]
}

// WordSelection contient un mot et son indice
type WordSelection struct {
	Word string `json:"word"`
	Hint string `json:"hint"`
}

// GetRandomWordByDifficulty retourne un mot aléatoire selon la difficulté
func GetRandomWordByDifficulty(difficulty string) WordSelection {
	words, exists := wordCategories[difficulty]
	if !exists {
		difficulty = "medium" // Difficulté par défaut
		words = wordCategories[difficulty]
	}

	selectedWord := words[rand.Intn(len(words))]
	return WordSelection(selectedWord)
}

// GetHint retourne l'indice pour un mot donné et une difficulté donnée
func GetHint(word string, difficulty string) string {
	words, exists := wordCategories[difficulty]
	if !exists {
		difficulty = "medium" // Difficulté par défaut
		words = wordCategories[difficulty]
	}

	for _, wordWithHint := range words {
		if wordWithHint.Word == word {
			return wordWithHint.Hint
		}
	}

	return "No hint available"
}
