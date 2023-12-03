package Game

import (
	"fmt"
)

/* PlayerMove ask which move the player want to do */
func PlayerMove(board *Board) {
	var row, col int

	/* Asking the move until it's correct */
	for {
		fmt.Print("Entrez les coordonnées (ligne colonne) pour votre coup : ")
		_, err := fmt.Scanf("%d %d", &row, &col)
		if err != nil || row < 0 || row >= 3 || col < 0 || col >= 3 || board[row][col] != Empty {
			fmt.Println("Les coordonnées rentrées ne sont pas valides (déjà pris ou hors tableau), veuillez réessayer.")
			continue
		}
		break
	}

	/* Making the move */
	board[row][col] = PlayerX
}

/* AiMove decide which move the AI will do */
func AiMove(board *Board, method int) {
	bestValue := -1000
	var bestMove []int
	var moveValue int
	alpha := -1000
	beta := 1000

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == Empty {
				board[i][j] = PlayerO
				/* Choosing method */
				if method == 1 {
					moveValue = minMax(board, false, 0, 0)
				} else {
					moveValue = minMax(board, false, alpha, beta)
				}

				board[i][j] = Empty

				if moveValue > bestValue {
					bestMove = []int{i, j}
					bestValue = moveValue
				}

				/* Redefining alpha if AB method */
				if method == 2 {
					alpha = maxInt(alpha, moveValue)
				}
			}
		}
	}

	board[bestMove[0]][bestMove[1]] = PlayerO
}
