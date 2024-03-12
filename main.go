package main

import (
	"fmt"

	"github.com/Rayato159/tic-tac-toe-but-qlearning/src"
)

type Pos struct {
	X int
	Y int
}

func main() {
	var row, col int
	src.LoadAgent()

	fmt.Println("Start Tic Tac Toe Game!")
	fmt.Println("### Your Turn ###")

	for {
		if src.IsYourTurn() {
			fmt.Scan(&row, &col)

			if row < 0 || row > 2 || col < 0 || col > 2 {
				fmt.Println("Invalid Input!")
				continue
			}
		}
		row, col := src.Input(row, col)

		if src.IsDuplicate(row, col) {
			fmt.Println("Duplicate!")
			continue
		}

		src.MapUpdating(row, col)
		src.MapDrawing()

		if src.IsYourTurn() {
			if src.IsEnd('x') {
				fmt.Println("You Win!")
				break
			}
		} else {
			if src.IsEnd('o') {
				fmt.Println("You Lose!")
				break
			}
		}

		if src.IsDraw() {
			fmt.Println("Draw!")
			break
		}

		src.TurnChaning()
	}
}
