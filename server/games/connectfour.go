package games

import "errors"

const (
	EMPTY     uint8 = 0
	RED       uint8 = 1
	BLUE      uint8 = 2
	NO_WINNER uint8 = 0
)

type ConnectFourGame interface {
	PlayTurn(int) error
	CheckWin() uint8
	GetWhoseTurn() uint8
}

type connectFourGame struct {
	turn                  uint8
	board                 [][]uint8
	numFilledBoxes        int
	winningSequenceLength int
}

/* TODO: improve struct
type connectFourGame struct {
	turn *Player
	board [][]*Player
	numFilledBoxes int
	winningSequenceLength int
}
*/

func (game *connectFourGame) CheckWin() uint8 {
	numRows := len(game.board)
	numCols := len(game.board[0])
	checkVertical := func(colour uint8) bool {
		for j := 0; j < numCols; j++ {
			count := 0
			for i := 0; i < numRows; i++ {
				if game.board[i][j] != colour {
					count = 0
				} else {
					count++
					if count == game.winningSequenceLength {
						return true
					}
				}
			}
		}
		return false
	}

	checkHorizontal := func(colour uint8) bool {
		for i := 0; i < numRows; i++ {
			count := 0
			for j := 0; j < numCols; j++ {
				if game.board[i][j] != colour {
					count = 0
				} else {
					count++
					if count == game.winningSequenceLength {
						return true
					}
				}
			}
		}
		return false
	}

	checkForwardDiagonalStartingFrom := func(colour uint8, startingRow int, startingCol int) bool {
		currentDiagonalI := startingRow
		currentDiagonalJ := startingCol
		count := 0
		for currentDiagonalI < numRows && currentDiagonalJ >= 0 {
			if game.board[currentDiagonalI][currentDiagonalJ] != colour {
				count = 0
			} else {
				count++
				if count == game.winningSequenceLength {
					return true
				}
			}
			currentDiagonalI++
			currentDiagonalJ--
		}
		return false
	}

	checkBackwardDiagonalStartingFrom := func(colour uint8, startingRow int, startingCol int) bool {
		currentDiagonalI := startingRow
		currentDiagonalJ := startingCol
		count := 0
		for currentDiagonalI < numRows && currentDiagonalJ < numCols {
			if game.board[currentDiagonalI][currentDiagonalJ] != colour {
				count = 0
			} else {
				count++
				if count == game.winningSequenceLength {
					return true
				}
			}
			currentDiagonalI++
			currentDiagonalJ++
		}
		return false
	}

	checkForwardDiagonal := func(colour uint8) bool {
		/*
			Covers the case that the winning sequence lies on a diagonal including a slot
			from the topmost row
			e.g, for a 3x6 board with a winning sequence length of 2,
			     these would be the diagonals investigated
			| - | x | x |
			| x | x | - |
			| x | - | - |
			| - | - | - |
			| - | - | - |
			| - | - | - |
		*/
		for j := game.winningSequenceLength - 1; j < numCols; j++ {
			if checkForwardDiagonalStartingFrom(colour, 0, j) {
				return true
			}
		}

		/*
			Covers the case that the winning sequence lies on a diagonal including a slot
			from the rightmost column (excluding the first row -- we covered that in previous loop)
			e.g, for a 3x6 board with a winning sequence length of 2,
			     these would be the diagonals investigated
			| - | - | - |
			| - | - | x |
			| - | x | x |
			| x | x | x |
			| x | x | x |
			| x | x | - |
		*/
		for i := 1; i <= numRows-game.winningSequenceLength; i++ {
			if checkForwardDiagonalStartingFrom(colour, i, numCols-1) {
				return true
			}
		}

		return false
	}

	checkBackwardDiagonal := func(colour uint8) bool {
		/*
			Covers the case that the winning sequence lies on a diagonal including a slot
			from the topmost row
			e.g, for a 3x6 board with a winning sequence length of 2,
			     these would be the diagonals investigated
			| x | x | - |
			| - | x | x |
			| - | - | x |
			| - | - | - |
			| - | - | - |
			| - | - | - |
		*/
		for j := 0; j <= numCols-game.winningSequenceLength; j++ {
			if checkBackwardDiagonalStartingFrom(colour, 0, j) {
				return true
			}
		}

		/*
			Covers the case that the winning sequence lies on a diagonal including a slot
			from the rightmost column (excluding the first row -- we covered that in previous loop)
			e.g, for a 3x6 board with a winning sequence length of 2,
			     these would be the diagonals investigated
			| - | - | - |
			| x | - | - |
			| x | x | - |
			| x | x | x |
			| x | x | x |
			| - | x | x |
		*/
		for i := 1; i <= numRows-game.winningSequenceLength; i++ {
			if checkBackwardDiagonalStartingFrom(colour, i, 0) {
				return true
			}
		}

		return false
	}

	if checkVertical(RED) || checkHorizontal(RED) || checkForwardDiagonal(RED) || checkBackwardDiagonal(RED) {
		return RED
	}
	if checkVertical(BLUE) || checkHorizontal(BLUE) || checkForwardDiagonal(BLUE) || checkBackwardDiagonal(BLUE) {
		return BLUE
	}
	return NO_WINNER
}

func MakeConnectFourGame(numColumns int, numRows int, winningSequenceLength int) (ConnectFourGame, error) {
	if numColumns < 3 || numRows < 3 {
		return nil, errors.New("number of columns and rows must be greater than 3")
	}
	if winningSequenceLength > numColumns || winningSequenceLength > numRows {
		return nil, errors.New("winning sequence length must not be greater than number of rows or columns")
	}
	board := make([][]uint8, numRows)
	for i := 0; i < numRows; i++ {
		board[i] = make([]uint8, numColumns)
	}
	return &connectFourGame{RED, board, 0, winningSequenceLength}, nil
}

func (game *connectFourGame) PlayTurn(col int) error {
	numRows := len(game.board)
	numCols := len(game.board[0])
	if col >= numCols || col < 0 {
		return errors.New("index out of bounds")
	}
	row := -1
	for i := 0; i < numRows; i++ {
		if game.board[i][col] == 0 {
			row = i
		} else {
			break
		}
	}
	if row < 0 {
		return errors.New("column is full already")
	}
	if row+1 == numRows || game.board[row+1][col] != 0 {
		game.board[row][col] = game.turn
		if game.turn == RED {
			game.turn = BLUE
		} else {
			game.turn = RED
		}
		game.numFilledBoxes++
	}
	return nil
}

func (game *connectFourGame) GetWhoseTurn() uint8 {
	return game.turn
}
