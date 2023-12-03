package Morpion

/* Setting the symbols representing the players */

const (
	Empty   = ' '
	PlayerX = 'X'
	PlayerO = 'O'
)

/* Board is the type for the board game */
type Board [3][3]rune

/* checkFinalCondition checks if the board is in a win situation and returning the player who won if it exists */
func checkFinalCondition(board Board) int {
	for i := 0; i < 3; i++ {
		/* Checking lines */
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			if board[i][0] == PlayerX {
				return 1
			} else if board[i][0] == PlayerO {
				return -1
			}
		}

		/* Checking columns */
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			if board[0][i] == PlayerX {
				return 1
			} else if board[0][i] == PlayerO {
				return -1
			}
		}
	}

	/* Checking diagonals */
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		if board[0][0] == PlayerX {
			return 1
		} else if board[0][0] == PlayerO {
			return -1
		}
	}

	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		if board[0][2] == PlayerX {
			return 1
		} else if board[0][2] == PlayerO {
			return -1
		}
	}

	return 0
}

/* isFull checks if the board is full */
func isFull(board Board) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == Empty {
				return false
			}
		}
	}
	return true
}
