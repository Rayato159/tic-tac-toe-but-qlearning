package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/Rayato159/tic-tac-toe-but-qlearning/src"
)

var (
	agentX = &Agent{
		QTable: make(map[[3][3]rune]map[src.Pos][]float64),
	}

	agentO = &Agent{
		QTable: make(map[[3][3]rune]map[src.Pos][]float64),
	}
)

type Agent struct {
	QTable map[[3][3]rune]map[src.Pos][]float64
	Action src.Pos
}

// Q(s, a) = Q(s, a) + α * [r + γ * maxQ(s', a') - Q(s, a)]
func main() {
	agentXActionRecording := make(map[string]src.Pos)
	agentOActionRecording := make(map[string]src.Pos)

	qXScoreRecord := make(map[src.Pos][]float64)
	qOScoreRecord := make(map[src.Pos][]float64)

	const learningRate = 0.1
	const discountFactor = 0.9

	for i := range 1000 {
		var pos src.Pos
		if i == 0 {
			pos = src.Pos{X: 0, Y: 0}
		} else {
			pos = src.Pos{X: rand.Intn(3), Y: rand.Intn(3)}
		}

		var isX bool
		if i%2 == 0 {
			isX = true
		} else {
			isX = false
		}

		var updatedQ float64
		if i == 0 {
			updatedQ = q(isX, pos)
		} else {
			updatedQ = q(isX, pos) + learningRate*(discountFactor*maxQ(isX, pos)-q(isX, pos))
		}

		if isX {
			if agentX.QTable[src.Board] == nil {
				agentX.QTable[src.Board] = make(map[src.Pos][]float64)
				agentX.QTable[src.Board][pos] = make([]float64, 0)
			}

			qXScoreRecord[pos] = append(qXScoreRecord[pos], updatedQ)
			agentX.QTable[src.Board][pos] = qXScoreRecord[pos]

			agentXActionRecording[runeToStringConverting(src.Board)] = pos
		} else {
			if agentO.QTable[src.Board] == nil {
				agentO.QTable[src.Board] = make(map[src.Pos][]float64)
				agentO.QTable[src.Board][pos] = make([]float64, 0)
			}

			qOScoreRecord[pos] = append(qOScoreRecord[pos], updatedQ)
			agentO.QTable[src.Board][pos] = qOScoreRecord[pos]

			agentOActionRecording[runeToStringConverting(src.Board)] = pos
		}
	}

	agentOResult, _ := json.MarshalIndent(agentOActionRecording, "", "\t")

	os.WriteFile("agentOQTable.json", agentOResult, 0777)
}

func maxQ(isX bool, s src.Pos) float64 {
	var max float64

	if isX {
		for _, v := range agentX.QTable[src.Board][s] {
			if v > max {
				max = v
			}
		}
	} else {
		for _, v := range agentO.QTable[src.Board][s] {
			if v > max {
				max = v
			}
		}
	}

	return max
}

func q(isX bool, action src.Pos) float64 {
	row, col := action.X, action.Y

	src.MapDrawing()
	fmt.Println("*****")

	for {
		if src.IsDuplicate(row, col) {
			return 0
		}

		src.MapUpdating(row, col)
		src.MapDrawing()

		xTurn := src.IsYourTurn()

		if isX {
			if xTurn {
				if src.IsEnd('x') {
					return 1.0
				}
			} else {
				if src.IsEnd('o') {
					return -2.0
				}
			}
		} else {
			if xTurn {
				if src.IsEnd('x') {
					return -2.0
				}
			} else {
				if src.IsEnd('o') {
					return 1.0
				}
			}
		}

		if src.IsDraw() {
			return 1.0
		}

		src.TurnChaning()
		fmt.Println("*****")
	}
}

func runeToStringConverting(s [3][3]rune) string {
	var result string

	for i := range 3 {
		for j := range 3 {
			result += string(s[i][j])
		}
	}

	return result
}
