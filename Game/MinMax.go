package Game

import "Morpion"

/* minMax algorithm */
func minMax(board *Board, depth int, evalMax bool) int {
	score := evaluation(board)

	if score != 0 {
		return score
	}

	if evalMax {
		bestValue := -1 << 31
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == Empty {
					board[i][j] = PlayerX
					moveValue := minMax(board, depth+1, false)
					board[i][j] = Empty
					bestValue = Morpion.maxInt(bestValue, moveValue)
				}
			}
		}
		return bestValue
	} else {
		bestValue := 1 << 31
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == Empty {
					board[i][j] = PlayerO
					moveValue := minMax(board, depth+1, true)
					board[i][j] = Empty
					bestValue = Morpion.minInt(bestValue, moveValue)
				}
			}
		}
		return bestValue
	}
}

/* evaluation returns the heuristic for a node */
func evaluation(board *Board) int {
	if isVictory(board, PlayerX) {
		return 100
	} else if isVictory(board, PlayerO) {
		return -100
	} else if isDraw(board) {
		return 50
	}

	return 0
}

/* isVictory checks if the actual node is a victory one */
func isVictory(board *Board, player rune) bool {
	/* Checking lines and columns */
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}

	/* Checking diagonals */
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}

/* isDraw checks if the actual node is a draw one (must be called after isVictory) */
func isDraw(board *Board) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == Empty {
				return false
			}
		}
	}
	return true
}
