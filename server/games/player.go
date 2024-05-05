package games

import (
	"log"

	"github.com/gorilla/websocket"
)

type Player struct {
	conn     *websocket.Conn
	username string
	verified bool
	messageToSend chan string
}

type WriteMessage struct {
	messageType uint

}

type GameStartingInMessage struct {
	secondsLeft uint
}

type RoundResultsMessage struct {
	
}

type GameEndingMessage struct {

}

type ReceivedMessage struct {
	answer string
	round  uint
}

func NewPlayer(conn *websocket.Conn, username string, verified bool) *Player {
	return &Player{
		conn: conn, 
		username: username, 
		verified: verified,
		messageToSend: make(chan string),
	}
}

func (player *Player) Write() {
	for {
		message := <- player.messageToSend
		if err := player.conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println(err)
			return
		}
	}
}

func (player *Player) Read() {
	for {
		var receivedMessage ReceivedMessage
		err := player.conn.ReadJSON(&receivedMessage)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
