package Morpion

import "fmt"

/* playerMove ask which move the player want to do */
func playerMove(board *Board) {
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

/* aiMove decide which move the AI will do */
func aiMove(board *Board) {
	bestValue := -1 << 31
	var bestMove []int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == Empty {
				board[i][j] = PlayerO
				moveValue := minMax(board, 0, false)
				board[i][j] = Empty

				if moveValue > bestValue {
					bestMove = []int{i, j}
					bestValue = moveValue
				}
			}
		}
	}

	board[bestMove[0]][bestMove[1]] = PlayerO
}
