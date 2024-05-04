package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rdhillon1016/games-web-app/server/handlers"
)

func main() {
	router := gin.Default()
	// router.GET("ws/trivia/:roomId", handlers.HandleTriviaJoin)
	router.GET("ws/trivia/create", handlers.HandleTriviaCreate)
	// router.GET("ws/connectfour/:roomId", handlers.HandleConnectFourJoin)
	// router.GET("ws/connectfour/create", handlers.HandleConnectFourCreate)
	router.Run("localhost:8080")
}
