package games

import "github.com/gorilla/websocket"

type Player struct {
	conn *websocket.Conn
	username string
	verified bool
}

func NewPlayer(conn *websocket.Conn, username string, verified bool) *Player {
	return &Player{conn, username, verified}
}

func (player *Player) Write() {

}

func (player *Player) Read() {
	
}