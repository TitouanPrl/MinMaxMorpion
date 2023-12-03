package main

import (
	"Morpion/Game"
	"fmt"
	"math/rand"
)

func main() {

	/* Ask which solving method to use */
	var input int

	fmt.Println("Choisissez la méthode de résolution :")
	fmt.Println("1 - MinMax")
	fmt.Println("2 - Optimisation Alpha Beta")

	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	} else if input != 1 && input != 2 {
		panic(fmt.Errorf("Erreur de saisie"))
	}
	fmt.Println("Vous avez choisi la méthode de résolution :", input)

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
			Game.AiMove(&board, input)
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
