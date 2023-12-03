package Morpion

/* minMax algorithm */
func minMax(node *Board, depth int, evalMax bool) int {
	if depth == 0 || isVictory(node, PlayerX) || isVictory(node, PlayerO) || isDraw(node) {
		return evaluate(node)
	}

	if evalMax {
		bestValue := -1 << 31
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if node[i][j] == Empty {
					node[i][j] = PlayerX
					value := minMax(node, depth-1, false)
					node[i][j] = Empty
					bestValue = max(bestValue, value)
				}
			}
		}
		return bestValue
	} else {
		bestValue := 1 << 31
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if node[i][j] == Empty {
					node[i][j] = PlayerO
					value := minMax(node, depth-1, true)
					node[i][j] = Empty
					bestValue = min(bestValue, value)
				}
			}
		}
		return bestValue
	}
}

/* checkFinalCondition checks if the board is in a win situation and returning the player who won if it exists */
func checkFinalCondition(board *Board) int {
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

/* evaluation returns the heuristic for a node */
func evaluation(board *Board) int {
	var res int

	/* Checking situation */
	winPos := checkFinalCondition(board)

	/* Attributing points regarding the situation */
	if winPos == 1 {
		res = 100
	} else if winPos == -1 {
		res = -100
	} else if isFull(board) {
		res = 50
	} else {
		res = 0
	}

	return res
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

/* isDraw checks if the actual node is a draw one */
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
