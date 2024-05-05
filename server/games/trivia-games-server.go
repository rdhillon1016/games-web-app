package games

import "math/rand"

const roomIdLength = 6

type TriviaGamesServer struct {
	// TODO: make sure that using a struct pointer as a map key is fine
	playersToGames map[*Player]*TriviaGame
	Create         chan *Player
	AdvanceRound   chan string
	EndGame        chan string
	AddPlayer      chan *Player
	RemovePlayer   chan *Player
}

func NewTriviaGamesServer() *TriviaGamesServer {
	return &TriviaGamesServer{
		playersToGames: make(map[*Player]*TriviaGame),
		Create:         make(chan *Player),
	}
}

func (server *TriviaGamesServer) Run() {
	for {
		select {
		case initialPlayer := <-server.Create:
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
	// TODO: send the player a "created" message
}

func generateGameId() string {
	possibleCharacters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomString := make([]rune, roomIdLength)
	for i := range randomString {
		randomString[i] = possibleCharacters[rand.Intn(len(possibleCharacters))]
	}
	return string(randomString)
}
