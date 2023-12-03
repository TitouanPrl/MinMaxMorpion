package Game

/* minMax algorithm
* We don't use the depth argument because in the case of the Morpion we will test all the cases with no limit
 */
func minMax(board *Board, evalMax bool) int {
	score := evaluation(board)

	/* Stop case */
	if score != 0 {
		return score
	}

	if evalMax {
		bestValue := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == Empty {
					board[i][j] = PlayerO
					moveValue := minMax(board, false)
					board[i][j] = Empty
					bestValue = maxInt(bestValue, moveValue)
				}
			}
		}
		return bestValue
	} else {
		bestValue := 1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == Empty {
					board[i][j] = PlayerX
					moveValue := minMax(board, true)
					board[i][j] = Empty
					bestValue = minInt(bestValue, moveValue)
				}
			}
		}
		return bestValue
	}
}

/* evaluation returns the heuristic for a node */
func evaluation(board *Board) int {
	if IsVictory(board, PlayerX) {
		return -100
	} else if IsVictory(board, PlayerO) {
		return 100
	} else if IsDraw(board) {
		return 50
	}

	return 0
}

/* IsVictory checks if the actual node is a victory one */
func IsVictory(board *Board, player rune) bool {
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

/* IsDraw checks if the actual node is a draw one (must be called after IsVictory) */
func IsDraw(board *Board) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == Empty {
				return false
			}
		}
	}
	return true
}
