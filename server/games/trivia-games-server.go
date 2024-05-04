package games

import "math/rand"

const ROOM_ID_LENGTH = 6

type TriviaGamesServer struct {
	// TODO: make sure that using a struct pointer as a map key is fine
	playersToGames map[*Player]*TriviaGame
	create chan *Player
}

func NewTriviaGamesServer(gameId string, initialPlayer *Player) *TriviaGamesServer {
	return &TriviaGamesServer{}
}

func (server *TriviaGamesServer) Run() {
	for {
		select {
		case initialPlayer := <- server.create:
			server.registerNewGame(initialPlayer)
		}
	}
}

func (server *TriviaGamesServer) registerNewGame(player *Player) {
	// TODO: Fix the unlikely but possible scenario where the id clashes
	// 		 with an existing id
	gameId := generateGameId()
	triviaGame := NewTriviaGame(gameId, player)
	server.playersToGames[player] = triviaGame
}

func generateGameId() string {
	possibleCharacters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomString := make([]rune, ROOM_ID_LENGTH)
	for i := range randomString {
		randomString[i] = possibleCharacters[rand.Intn(len(possibleCharacters))]
	}
	return string(randomString)
}