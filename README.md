# 8Bit Hangman - Backend 🖥️

API backend for the 8Bit Hangman game, built with Go and Gin.

## Goal
Provide a REST API for the Hangman game logic to be consumed by the frontend.

## Tech Stack
- Go
- Gin (HTTP server)
- In-memory storage for sessions (or optional DB for leaderboard)
- Unit tests with Go `testing` package

## Getting Started
```bash
# Initialize Go modules
go mod tidy

# Run the server
go run cmd/server/main.go
```

The API will run on `http://localhost:8080` by default.

---

## Project Structure
- `internal/hangman` → game logic (core of the game)
- `internal/api` → Gin HTTP handlers
- `internal/storage` → session management (in-memory or DB)
- `cmd/server` → main entrypoint for the server
- `tests` → unit tests for game logic and API endpoints

---

## API Endpoints (examples)
- `POST /game` → create a new game, returns game ID and initial state
- `POST /game/:id/guess` → submit a letter guess, returns updated game state
- `GET /game/:id` → get current state of a game

### Example response
```json
{
  "id": "abc123",
  "state": "goodGuess",
  "turnsLeft": 5,
  "foundLetters": ["H", "_", "N", "G", "M", "A", "N"],
  "usedLetters": ["H", "A"],
  "wordLength": 7
}
```

---

## TODO
- Implement all API endpoints
- Connect frontend to API
- Add unit tests for game logic
- Optionally, add persistent storage for scores
