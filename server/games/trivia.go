package games

type TriviaGame struct {
	id string
	// TODO: make sure that using a struct pointer as a map key is fine
	players map[*Player]bool
	join chan *Player
	leave chan *Player
	play chan *Player
}

// type TriviaGameBroadcastMessage struct {
// 	messageType int `json:"messageType"`
// 	roundResults map[string]int
// 	roundNumber
// }

// type RoundResults map[string]int

// type TriviaQuestion struct {
// 	question string
// 	options map[string]string
// 	answer string
// }

func NewTriviaGame(gameId string, initialPlayer *Player) *TriviaGame {
	players := make(map[*Player]bool)
	players[initialPlayer] = true
	return &TriviaGame{id: gameId, players: players}
}

// func (game *TriviaGame) Run() {
// 	/*
// 	* listen on channels for the following messages
// 	* user attempting to join a trivia game
// 	* user left a trivia game
// 	* trivia answer submission
// 	*/
// 	for {
// 		select {
// 		case player := <- game.join:
// 			game.addPlayer(player)
// 		case player := <- game.leave:
// 			game.removePlayer(player)
// 			if (game.)
// 		case answer := <- game.play:
// 			game.playTurn(answer)
// 		}
// 	}
// }

// func (game *TriviaGame) addPlayer(player *Player) {}
// func (game *TriviaGame) removePlayer(player *Player) {}
// func (game *TriviaGame) playTurn (player *Player) {}
