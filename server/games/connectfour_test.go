package games_test

import (
	"testing"

	"github.com/rdhillon1016/games-web-app/server/games"
)

func TestMakeWithInvalidSequenceNumber(t *testing.T) {
	const numRows = 3
	const numCols = 3
	const winningSequence = 4
	_, err := games.MakeConnectFourGame(numRows, numCols, winningSequence)
	if err == nil {
		t.Errorf("Expected an error since the winning sequence %d is greater the num of cols %d or num of rows %d", winningSequence, numCols, numRows)
	}
}

func TestHorizontalOutOfBounds(t *testing.T) {
	game, _ := games.MakeConnectFourGame(3, 3, 3)
	err := game.PlayTurn(4)
	if err == nil {
		t.Error("Expected error for out of bounds play")
	}
}

func TestColumnAlreadyFull(t *testing.T) {
	game, _ := games.MakeConnectFourGame(3, 3, 3)
	game.PlayTurn(1)
	game.PlayTurn(1)
	game.PlayTurn(1)
	err := game.PlayTurn(1)
	if err == nil {
		t.Error("Expected error for placement on already-full column")
	}
}

func TestVerticalWin(t *testing.T) {
	game, _ := games.MakeConnectFourGame(3, 3, 3)
	expectedWinningPlayer := game.GetWhoseTurn()

	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(0)
	result := game.CheckWin()
	if result != expectedWinningPlayer {
		t.Errorf("Expected %d to win", expectedWinningPlayer)
	}
}

func TestHorizontalWin(t *testing.T) {
	game, _ := games.MakeConnectFourGame(3, 3, 3)
	expectedWinningPlayer := game.GetWhoseTurn()

	game.PlayTurn(0)
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(1)
	game.PlayTurn(2)
	result := game.CheckWin()
	if result != expectedWinningPlayer {
		t.Errorf("Expected %d to win", expectedWinningPlayer)
	}
}

func TestForwardDiagonalWin(t *testing.T) {
	game, _ := games.MakeConnectFourGame(3, 3, 2)
	expectedWinningPlayer := game.GetWhoseTurn()

	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(1)
	result := game.CheckWin()
	if result != expectedWinningPlayer {
		t.Errorf("Expected %d to win", expectedWinningPlayer)
	}
}

func TestBackwardDiagonalWin(t *testing.T) {
	game, _ := games.MakeConnectFourGame(3, 3, 2)
	expectedWinningPlayer := game.GetWhoseTurn()

	game.PlayTurn(1)
	game.PlayTurn(0)
	game.PlayTurn(0)
	result := game.CheckWin()
	if result != expectedWinningPlayer {
		t.Errorf("Expected %d to win", expectedWinningPlayer)
	}
}

func TestDraw(t *testing.T) {
	game, _ := games.MakeConnectFourGame(4, 4, 3)
	expectedWinningPlayer := game.GetWhoseTurn()

	// Fill up first row
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(2)
	game.PlayTurn(3)

	// Fill up second row
	game.PlayTurn(0)
	game.PlayTurn(1)
	game.PlayTurn(2)
	game.PlayTurn(3)

	// Fill up third row
	game.PlayTurn(1)
	game.PlayTurn(0)
	game.PlayTurn(3)
	game.PlayTurn(2)

	// Fill up fourth row
	game.PlayTurn(1)
	game.PlayTurn(0)
	game.PlayTurn(3)
	game.PlayTurn(2)

	result := game.CheckWin()
	if result != games.NO_WINNER {
		t.Errorf("Expected %d to win", expectedWinningPlayer)
	}
}
