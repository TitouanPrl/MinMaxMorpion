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

/* initNewBoard returns an empty board */
func initNewBoard() Board {
	var board Board

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = Empty
		}
	}
	return board
}

/* minInt returns the min between 2 int */
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/* maxInt returns the max between 2 int */
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/* isFull checks if the board is full */
func isFull(board *Board) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == Empty {
				return false
			}
		}
	}
	return true
}
