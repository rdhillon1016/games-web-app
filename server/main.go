package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rdhillon1016/games-web-app/server/games"
	"github.com/rdhillon1016/games-web-app/server/handlers"
)

func main() {
	triviaGamesServer := games.NewTriviaGamesServer()
	go triviaGamesServer.Run()

	router := gin.Default()
	// router.GET("ws/trivia/:roomId", handlers.HandleTriviaJoin)
	router.GET("ws/trivia/create", func(c *gin.Context) {
		handlers.HandleTriviaCreate(c, triviaGamesServer)
	})
	router.Run("localhost:8080")
}
