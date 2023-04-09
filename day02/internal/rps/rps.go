// Package rps implements the game of rock, paper, scissors.
package rps

import (
	"os"
	"path/filepath"
	"strings"
)

// points maps the game state to a number of points.
var points = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
	"loss":     0,
	"draw":     3,
	"win":      6,
}

var winningMoves = map[string]string{
	"rock":     "paper",
	"paper":    "scissors",
	"scissors": "rock",
}

// roundResult determines if the player wins, draws, or looses the round.
func roundResult(playerMove string, oppMove string) (string, string) {
	var (
		playerResult string
		oppResult    string
	)
	if playerMove == "rock" && oppMove == "paper" {
		playerResult = "loss"
		oppResult = "win"
	}
	if playerMove == "paper" && oppMove == "scissors" {
		playerResult = "loss"
		oppResult = "win"
	}
	if playerMove == "scissors" && oppMove == "rock" {
		playerResult = "loss"
		oppResult = "win"
	}
	if playerMove == "rock" && oppMove == "scissors" {
		playerResult = "win"
		oppResult = "loss"
	}
	if playerMove == "scissors" && oppMove == "paper" {
		playerResult = "win"
		oppResult = "loss"
	}
	if playerMove == "paper" && oppMove == "rock" {
		playerResult = "win"
		oppResult = "loss"
	}
	if playerMove == "rock" && oppMove == "rock" {
		playerResult = "draw"
		oppResult = "draw"
	}
	if playerMove == "scissors" && oppMove == "scissors" {
		playerResult = "draw"
		oppResult = "draw"
	}
	if playerMove == "paper" && oppMove == "paper" {
		playerResult = "draw"
		oppResult = "draw"
	}
	return playerResult, oppResult
}

// encryptedStrategy maps the input to a set of allowable game states.
var encryptedStrategy = map[string]string{
	"A": "rock",
	"Y": "paper",
	"B": "paper",
	"X": "rock",
	"C": "scissors",
	"Z": "scissors",
}

type Player struct {
	// RoundScores keep strack of the total score for a given Player
	// on each round.
	RoundScores []int
}

func NewPlayer() Player {
	return Player{RoundScores: []int{}}
}

// Score calculates the total score for a player from every round.
func (p *Player) score() int {
	sum := 0
	for _, s := range p.RoundScores {
		sum += s
	}
	return sum
}

// StartGame begins the game and returns the players score.
func StartGame(fpath string) (int, error) {
	opponent := NewPlayer()
	p := NewPlayer()

	fp, err := filepath.Abs(fpath)
	if err != nil {
		return 0, err
	}

	data, err := os.ReadFile(fp)
	if err != nil {
		return 0, err
	}

	rounds := strings.Split(string(data), "\n")
	for _, round := range rounds {
		moves := strings.Split(round, " ")
		if len(moves) == 1 {
			continue
		}
		opponentMove := moves[0]
		recommendedMove := moves[1]

		oppPoints := 0
		playerPoints := 0

		// Decrypt the moves.
		decryptedOppMove := encryptedStrategy[opponentMove]
		decryptedRecMove := encryptedStrategy[recommendedMove]

		// Check the players round result.
		playerResult, opponentResult := roundResult(decryptedRecMove, decryptedOppMove)
		playerPoints += points[playerResult]
		oppPoints += points[opponentResult]

		// Calculate the points for these moves.
		oppPoints += points[decryptedOppMove]
		playerPoints += points[decryptedRecMove]

		// Add the points/score for each round to the player and opponent.
		p.RoundScores = append(p.RoundScores, playerPoints)
		opponent.RoundScores = append(opponent.RoundScores, oppPoints)
	}

	return p.score(), nil
}
