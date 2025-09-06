# The 8Bits Hangman - Backend 🖥️

API backend for The 8Bits Hangman game, built with Go and Gin framework.

## Overview

This backend provides a robust REST API that powers the Hangman game logic, designed to be consumed by the frontend application. It handles game state management, word selection, guess validation, and scoring.

## Tech Stack

- **Go** - Fast and efficient programming language
- **Gin** - High-performance HTTP web framework
- **PostgreSQL** - Optional database for leaderboard and persistent storage
- **Go Testing** - Comprehensive test suite with the standard Go `testing` package

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Git

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/the-8bits-hangman-back.git
cd the-8bits-hangman-back

# Initialize Go modules
go mod tidy

# Run the server
go run cmd/server/main.go
```

The API will be available at `http://localhost:8080` by default.

## Project Structure

- `cmd/server/` - Main application entry point
- `internal/hangman/` - Core game logic and business rules
- `internal/api/` - HTTP handlers and routing with Gin
- `internal/storage/` - Data persistence layer (in-memory or PostgreSQL)
- `tests/` - Unit and integration tests
- `ai-agents/` - AI agent configurations for development assistance

## API Endpoints

### Game Management

- `POST /api/games` - Create a new game session
- `GET /api/games/:id` - Retrieve current game state
- `POST /api/games/:id/guess` - Submit a letter guess
- `DELETE /api/games/:id` - Abandon a game

### Leaderboard (Optional)

- `GET /api/leaderboard` - Get top scores
- `POST /api/leaderboard` - Submit a score

### Example Response

```json
{
  "id": "game_abc123",
  "status": "in_progress",
  "turnsLeft": 5,
  "word": ["H", "_", "N", "G", "M", "_", "N"],
  "usedLetters": ["H", "N", "G", "M"],
  "difficulty": "medium",
  "score": 120
}
```

## Development

### Running Tests

```bash
go test ./tests/... -v
```
