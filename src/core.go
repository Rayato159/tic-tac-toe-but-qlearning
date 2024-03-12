package src

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type Pos struct {
	X int
	Y int
}

var (
	Board = [3][3]rune{
		{'-', '-', '-'},
		{'-', '-', '-'},
		{'-', '-', '-'},
	}

	yourTurn = true

	agent = make(map[[3][3]rune]Pos)
)

func Input(row, col int) (int, int) {
	if yourTurn {
		return row, col
	}

	if pos, ok := agent[Board]; ok {
		return pos.X, pos.Y
	} else {
		return rand.Intn(3), rand.Intn(3)
	}
}

func LoadAgent() {
	rawAgent := make(map[string]Pos)

	rawJSON, _ := os.ReadFile("./agentOQTable.json")

	if err := json.Unmarshal(rawJSON, &rawAgent); err != nil {
		panic("Failed to load agent")
	}

	for k, v := range rawAgent {
		var runeMapping [3][3]rune

		for l := range 9 {
			runeMapping[l/3][l%3] = rune(k[l])
		}

		agent[runeMapping] = Pos{X: v.X, Y: v.Y}
	}
}

func MapUpdating(i, j int) {
	if yourTurn {
		Board[i][j] = 'x'
	} else {
		Board[i][j] = 'o'
	}
}

func MapDrawing() {
	for i := range 3 {
		for j := range 3 {
			fmt.Printf("%c", Board[i][j])
		}
		fmt.Println()
	}
}

func IsDuplicate(i, j int) bool {
	return Board[i][j] != '-'
}

func IsDraw() bool {
	return (Board[0][0] != Board[0][2] &&
		Board[2][0] != Board[2][2]) &&
		(Board[0][0] != '-' &&
			Board[0][2] != '-' &&
			Board[2][0] != '-' &&
			Board[2][2] != '-')
}

func TurnChaning() {
	yourTurn = !yourTurn

	if yourTurn {
		fmt.Println("### Your Turn ###")
	} else {
		fmt.Println("### Agent Turn ###")
	}
}

func IsYourTurn() bool {
	return yourTurn
}

func IsEnd(marker rune) bool {
	// row win
	if count := rowWining(marker); count == 3 {
		return true
	}

	// column win
	if count := columnWinning(marker); count == 3 {
		return true
	}

	// diagonal win
	if count := diagonalWinning(marker); count == 3 {
		return true
	}

	return false
}

func rowWining(marker rune) int {
	var count int

	for i := range 3 {
		for j := range 3 {
			if Board[i][j] == marker {
				count++
			}
		}
		if count == 3 {
			return count
		}
		count = 0
	}

	return count
}

func columnWinning(marker rune) int {
	var count int

	for i := range 3 {
		for j := range 3 {
			if Board[j][i] == marker {
				count++
			}
		}
		if count == 3 {
			return count
		}
		count = 0
	}

	return count
}

func diagonalWinning(marker rune) int {
	var count int

	for i := range 3 {
		if Board[i][i] == marker {
			count++
		}
	}
	if count == 3 {
		return count
	}

	count = 0

	for i := range 3 {
		if Board[i][2-i] == marker {
			count++
		}
	}

	return count
}
