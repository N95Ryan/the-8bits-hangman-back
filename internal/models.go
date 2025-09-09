package internal

type Game struct {
	ID        string   `json:"id"`
	Word      string   `json:"word"`
	Guesses   []string `json:"guesses"`
	Remaining int      `json:"remaining"`
	Status    string   `json:"status"` // "in_progress", "won", "lost"
}

type Player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type LeaderboardEntry struct {
	Player string `json:"player"`
	Score  int    `json:"score"`
}

type GuessRequest struct {
	Letter string `json:"letter" binding:"required,len=1"`
}

type CreateGameRequest struct {
	PlayerName string `json:"player_name" binding:"required,min=3,max=50"`
}
