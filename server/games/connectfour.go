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

func (game *connectFourGame) CheckWin() uint8 {
	checkVertical := func(colour uint8) bool {
		numRows := len(game.board)
		numCols := len(game.board[0])

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
		numRows := len(game.board)
		numCols := len(game.board[0])

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

	checkForwardDiagonal := func(colour uint8) bool {
		numRows := len(game.board)
		numCols := len(game.board[0])

		for j := game.winningSequenceLength - 1; j < numCols; j++ {
			i := 0
			currentJ := j
			count := 0
			for i < numRows && currentJ >= 0 {
				if game.board[i][j] != colour {
					count = 0
				} else {
					count++
					if count == game.winningSequenceLength {
						return true
					}
				}
				i++
				currentJ--
			}
		}
		return false
	}

	checkBackwardDiagonal := func(colour uint8) bool {
		numRows := len(game.board)
		numCols := len(game.board[0])
		for j := numCols - game.winningSequenceLength; j >= 0; j-- {
			i := 0
			currentJ := j
			count := 0
			for i < numRows && currentJ >= 0 {
				if game.board[i][j] != colour {
					count = 0
				} else {
					count++
					if count == game.winningSequenceLength {
						return true
					}
				}
				i++
				currentJ++
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
	var row int
	for i := 0; i < numRows; i++ {
		if game.board[i][col] != 0 {
			row = i - 1
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
