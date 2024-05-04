package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rdhillon1016/games-web-app/server/games"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

// func HandleTriviaJoin(c *gin.Context) {
// 	roomId := c.Param("name")
// }

func HandleTriviaCreate(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	player := games.NewPlayer(conn, "testusername", false)

}

// func HandleConnectFourJoin(c *gin.Context) {
// 	roomId := generateRoomId()
// }

// func HandleConnectFourCreate(c *gin.Context) {
// 	roomId := generateRoomId()
// }
