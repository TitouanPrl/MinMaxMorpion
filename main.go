package main

import (
	"Morpion/Game"
	"fmt"
	"math/rand"
)

func main() {

	/* Setting the board to a new game */
	board := Game.InitNewBoard()

	depth := 9
	turn := rand.Intn(2)

	/* Running the game until all cells are full */
	for !Game.IsVictory(&board, Game.PlayerX) && !Game.IsVictory(&board, Game.PlayerO) && !Game.IsDraw(&board) {
		Game.PrintPlayground(board)

		if turn == 0 {
			Game.PlayerMove(&board)
			depth--
		} else {
			Game.AiMove(&board)
			depth--
		}

		/* Changing player */
		turn = 1 - turn

		Game.CallClearTerminal()
	}

	Game.PrintPlayground(board)

	/* Printing the winner */
	if Game.IsVictory(&board, Game.PlayerX) {
		fmt.Println("Vous avez gagné !")
	} else if Game.IsVictory(&board, Game.PlayerO) {
		fmt.Println("L'IA a gagné...")
	} else {
		fmt.Println("Match nul, bien joué quand même")
	}
}
