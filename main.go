package Morpion

import (
	"fmt"
	"math/rand"
)

func main() {

	/* Setting the board to a new game */
	board := initNewBoard()

	turn := rand.Intn(2)
	fmt.Println("Tour initial :", turn)

	/* Running the game until all cells are full */
	for !isVictory(&board, PlayerX) && !isVictory(&board, PlayerO) && !isDraw(&board) {
		PrintPlayground(board)

		if turn == 0 {
			playerMove(&board)
		} else {
			aiMove(&board)
		}

		/* Changing player */
		turn = 1 - turn

		CallClearTerminal()
	}

	PrintPlayground(board)

	/* Printing the winner */
	if isVictory(&board, PlayerX) {
		fmt.Println("Vous avez gagné !")
	} else if isVictory(&board, PlayerO) {
		fmt.Println("L'IA a gagné...")
	} else {
		fmt.Println("Match nul, bien joué quand même")
	}
}
