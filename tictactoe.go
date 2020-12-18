package main

import (
	"fmt"

	"zartosht.dev/programs/go/games/tictoctoe/boolean"
)

func main() {
	var columns, rows, xWon, oWon int
	fmt.Print("How many columns> ")
	fmt.Scan(&columns)
	fmt.Print("How many rows> ")
	fmt.Scan(&rows)
	boardTiles := make([]interface{}, rows*columns)
	// Create a tic-tac-toe board.
	const top, middle, base, void = "|¯¯¯¯¯", "|  %s  ", "|_____", "|     "
	var board string
	for row := 0; row < rows; row++ {
		for column := 0; column < columns*3; column++ {
			if column < columns {
				if row == 0 {
					board += top
				} else {
					board += void
				}
			} else if column < columns*2 {
				board += middle
			} else {
				board += base
			}
			if column == (columns)-1 || column == (columns*2)-1 || column == (columns*3)-1 {
				board += "|\n"
			}
		}
	}

	r := boolean.New()
	for i := range boardTiles {
		if r.Bool() {
			boardTiles[i] = "X"
			xWon++
		} else {
			boardTiles[i] = "O"
			oWon++
		}
	}

	if xWon > oWon {
		fmt.Println("X won the game!")
	} else if oWon > xWon {
		fmt.Println("O won the game!")
	} else {
		fmt.Println("TIE!")
	}

	fmt.Printf(board, boardTiles...)
}
