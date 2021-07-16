package game

import (
	"farkle/farkle"
	"fmt"
	"math/rand"
	"time"
)

const maxThrowSize = 6

type player struct {
	name  string
	score int
}

type players [2]player

type turnChoice string

const (
	endTurn    turnChoice = "e"
	throwAgain            = "t"
)

type Game struct {
	ui            UI
	userInput     UserInput
	players       players
	maxScore      int
	currentPlayer int
	throwSize     int
	turnScore     int
}

func NewGame(ui UI, userInput UserInput, maxScore int) Game {
	ui.clearScreen()
	return Game{
		ui:            ui,
		userInput:     userInput,
		maxScore:      maxScore,
		currentPlayer: rand.Intn(2),
		throwSize:     maxThrowSize, turnScore: 0,
	}
}

func (g *Game) CreatePlayers() {
	name1 := g.userInput.readPlayerName("one")
	name2 := g.userInput.readPlayerName("two")
	g.players = players{{name1, 0}, {name2, 0}}
}

// Main loop, does not have unit tests
func (g *Game) Run() {
	var winner *player
	for winner == nil {
		g.ui.clearScreen()
		currentPlayer := g.currentPlayer
		throw, _ := farkle.RollNDice(g.throwSize)
		g.turn(throw)
		winner = g.getWinner()
		if g.currentPlayer != currentPlayer {
			time.Sleep(3 * time.Second)
		}
	}
	g.ui.showWinner(*winner)
}

func (g *Game) getWinner() *player {
	for _, player := range g.players {
		if player.score >= g.maxScore {
			return &player
		}
	}
	return nil
}

func (g *Game) turn(throw []int) {
	g.ui.showScore(g.players[g.currentPlayer])
	g.ui.showThrow(throw)
	if isBust(throw) {
		g.ui.showBust()
		g.turnScore = 0
		g.endTurn()
	} else {
		throwScore, countDiceKeptByPlayer, endThrowChoice := g.makePlayerSelectDice(throw)
		g.turnScore += throwScore
		if endThrowChoice == endTurn {
			g.ui.showFinalTurnScore(g.turnScore)
			g.endTurn()
		} else {
			g.updateThrowSize(countDiceKeptByPlayer)
		}
	}
}

func isBust(throw []int) bool {
	isBust, _ := farkle.IsBust(throw)
	return *isBust
}

func (g *Game) makePlayerSelectDice(throw []int) (int, int, turnChoice) {
	var throwScore int
	var dice []int
	var choice turnChoice
	for throwScore == 0 {
		diceIndices := g.userInput.readDiceToKeep(len(throw))
		dice = getDiceAtIndices(throw, diceIndices)
		throwScore = farkle.CalculateScore(dice)
		if throwScore == 0 {
			g.ui.showError(fmt.Errorf("This selection would bring you no point"))
			continue
		}
		g.ui.showTurnScore(g.turnScore + throwScore)
		choice = g.userInput.readEndThrowChoice()
	}
	return throwScore, len(dice), choice
}

func getDiceAtIndices(throw []int, diceIndices []int) []int {
	dice := []int{}
	for _, i := range diceIndices {
		dice = append(dice, throw[i])
	}
	return dice
}

func (g *Game) endTurn() {
	g.players[g.currentPlayer].score += g.turnScore
	g.turnScore = 0
	g.throwSize = maxThrowSize
	g.currentPlayer = (g.currentPlayer + 1) % 2
}

func (g *Game) updateThrowSize(countDiceKeptByPlayer int) {
	g.throwSize -= countDiceKeptByPlayer
	if g.throwSize == 0 {
		// The player re-throw the whole hand
		g.throwSize = maxThrowSize
	}
}
