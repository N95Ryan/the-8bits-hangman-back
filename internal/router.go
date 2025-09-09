package internal

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.POST("/api/games", CreateGameHandler)
	r.GET("/api/games/:id", GetGameHandler)
	r.POST("/api/games/:id/guess", SubmitGuessHandler)
	r.DELETE("/api/games/:id", AbandonGameHandler)
	r.GET("/api/leaderboard", GetLeaderboardHandler)
	r.POST("/api/leaderboard", SubmitScoreHandler)
}
