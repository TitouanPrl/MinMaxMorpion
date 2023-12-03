package Morpion

import (
	"Morpion/Game"
	"fmt"
	"math/rand"
)

func main() {

	/* Setting the board to a new game */
	board := Game.initNewBoard()

	turn := rand.Intn(2)
	fmt.Println("Tour initial :", turn)

	/* Running the game until all cells are full */
	for !Game.isVictory(&board, Game.PlayerX) && !Game.isVictory(&board, Game.PlayerO) && !Game.isDraw(&board) {
		Game.PrintPlayground(board)

		if turn == 0 {
			Game.playerMove(&board)
		} else {
			Game.aiMove(&board)
		}

		/* Changing player */
		turn = 1 - turn

		Game.CallClearTerminal()
	}

	Game.PrintPlayground(board)

	/* Printing the winner */
	if Game.isVictory(&board, Game.PlayerX) {
		fmt.Println("Vous avez gagné !")
	} else if Game.isVictory(&board, Game.PlayerO) {
		fmt.Println("L'IA a gagné...")
	} else {
		fmt.Println("Match nul, bien joué quand même")
	}
}
