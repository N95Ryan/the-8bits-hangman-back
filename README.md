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
go run main.go
```

The API will be available at `http://localhost:8080` by default.

## Project Structure

- `main.go` - Main application entry point
- `game/` - Core game logic and word management
  - `game.go` - Game state and mechanics
  - `wordlist.go` - Word selection and categorization
  - `score.go` - Scoring system and leaderboard
- `handlers/` - HTTP request handlers
  - `gameHandler.go` - Game-related API endpoints
  - `userHandler.go` - User authentication and management
- `models/` - Data structures and business logic
  - `user.go` - User model and authentication
- `utils/` - Helper functions and utilities
  - `helpers.go` - Common utility functions

## API Endpoints

### Game Management

- `POST /api/games` - Create a new game session
- `GET /api/games/:id` - Retrieve current game state
- `POST /api/games/:id/guess` - Submit a letter guess
- `DELETE /api/games/:id` - Abandon a game

### User Management

- `POST /api/users/register` - Register a new user
- `POST /api/users/login` - Authenticate a user

### Leaderboard

- `GET /api/leaderboard` - Get top scores
- `POST /api/leaderboard` - Submit a score

### ID Format

All IDs in the system (games, users, tokens) follow a standardized format:

- 3 digits followed by 3 uppercase letters (e.g., `123ABC`, `789XYZ`)
- This format provides ~17.5 million unique combinations
- Easy to read, communicate, and remember

### Example Response

```json
{
  "id": "755XRK",
  "status": "in_progress",
  "remaining": 8,
  "word": "_______",
  "guesses": ["A", "E"],
  "score": 0
}
```

## Development

### Running Tests

```bash
go test ./... -v
```

### API Testing

You can use Postman or any other API testing tool to interact with the endpoints.
For example, to create a new game:

```
POST http://localhost:8080/api/games
Content-Type: application/json

{
  "player_name": "Player1"
}
```
