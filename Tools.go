package Morpion

import (
	"fmt"
)

/* Setting the symbols representing the players */
const (
	Empty   = ' '
	PlayerX = 'X'
	PlayerO = 'O'
)

/* Board is the type for the board game */
type Board [3][3]rune

/* PrintPlayground displays the playground */
func PrintPlayground(board Board) {
	lenght := 4

	/* Printing top line */
	fmt.Printf("┌")
	for i := 0; i < lenght-2; i++ {
		fmt.Printf("───┬")
	}
	fmt.Printf("───┐\n")

	for i := 0; i < lenght-1; i++ {

		/* Printing line */
		for j := 0; j < lenght-1; j++ {
			if board[i][j] == 0 {
				fmt.Printf("│   ")
			} else {
				fmt.Printf("│ %v ", board[i][j])
			}
		}
		fmt.Printf("│\n")

		/* Printing separator */
		if i < lenght-2 {
			fmt.Printf("├")
			for k := 0; k < lenght-2; k++ {
				fmt.Printf("───┼")
			}
			fmt.Printf("───┤\n")
		}
	}

	/* Printing last line */
	fmt.Printf("└")
	for i := 0; i < lenght-2; i++ {
		fmt.Printf("───┴")
	}
	fmt.Printf("───┘\n")
}
