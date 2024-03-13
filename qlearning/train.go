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
		QTable: make(map[[3][3]rune]map[src.Pos]float64),
	}

	agentO = &Agent{
		QTable: make(map[[3][3]rune]map[src.Pos]float64),
	}
)

type Agent struct {
	QTable map[[3][3]rune]map[src.Pos]float64
	Action src.Pos
}

func main() {
	const learningRate = 0.1
	const discountFactor = 0.9
	const numEpisodes = 1000

	for episode := range numEpisodes {
		playEpisode(episode, learningRate, discountFactor)
	}

	agentOResult := make(map[string]src.Pos)
	for k := range agentO.QTable {
		agentOResult[convertBoardToKey(k)] = selectAction(k, agentO)
	}

	agentOResultJSON, _ := json.MarshalIndent(agentOResult, "", "\t")
	os.WriteFile("agentOQTable.json", agentOResultJSON, 0777)
}

func convertBoardToKey(board [3][3]rune) string {
	key := ""
	for _, row := range board {
		for _, cell := range row {
			key += string(cell)
		}
	}
	return key
}

func playEpisode(episode int, learningRate, discountFactor float64) {
	isX := episode%2 == 0

	for !src.IsGameOver {
		var agent *Agent
		if isX {
			agent = agentX
		} else {
			agent = agentO
		}

		agent.Action = src.Pos{X: rand.Intn(3), Y: rand.Intn(3)}

		updatedQ := q(isX, agent.Action) + learningRate*(discountFactor*maxQ(src.Board, agent)-q(isX, agent.Action))
		updateAgentQTable(agent, src.Board, agent.Action, updatedQ)

		fmt.Println("Episode:", episode)

		src.IsGameOver = src.IsEnd('x') || src.IsEnd('o') || src.IsDraw()

		isX = !isX // Switch player turn
	}
}

func maxQ(board [3][3]rune, agent *Agent) float64 {
	max := -1000.0

	for _, v := range agent.QTable[board] {
		if v > max {
			max = v
		}
	}

	return max
}

func updateAgentQTable(agent *Agent, board [3][3]rune, pos src.Pos, updatedQ float64) {
	if agent.QTable[board] == nil {
		agent.QTable[board] = make(map[src.Pos]float64)
	}
	agent.QTable[board][pos] = updatedQ
}

func selectAction(board [3][3]rune, agent *Agent) src.Pos {
	// Select the action with the highest Q-value for the current state
	var bestAction src.Pos
	for pos := range agent.QTable[board] {
		bestAction = pos
	}

	return bestAction
}

func q(isX bool, action src.Pos) float64 {
	row, col := action.X, action.Y

	if src.IsDuplicate(row, col) {
		return -10.0
	}

	src.MapUpdating(row, col, isX)
	src.MapDrawing()

	if src.IsEnd('x') {
		if isX {
			return 1.0 // Player X wins
		} else {
			return -1.0 // Player O loses
		}
	}
	if src.IsEnd('o') {
		if isX {
			return -1.0 // Player X loses
		} else {
			return 1.0 // Player O wins
		}
	}

	return 0 // Game continues
}
