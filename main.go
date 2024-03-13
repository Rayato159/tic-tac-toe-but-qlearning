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
	isX := true
	var row, col int
	src.LoadAgent()

	fmt.Println("Start Tic Tac Toe Game!")

	for {
		if isX {
			fmt.Scan(&row, &col)

			if row < 0 || row > 2 || col < 0 || col > 2 {
				fmt.Println("Invalid Input!")
				continue
			}
		}
		row, col := src.Input(row, col, isX)

		for src.IsDuplicate(row, col) {
			fmt.Println("Duplicate!")
			continue
		}

		src.MapUpdating(row, col, isX)
		src.MapDrawing()

		src.IsGameOver = src.IsEnd('x') || src.IsEnd('o') || src.IsDraw()
		if src.IsGameOver {
			break
		}

		isX = !isX

		fmt.Println("***************")
	}
}
