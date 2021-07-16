package main

import (
	"farkle/game"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	ui := game.NewTermUI(os.Stdout)
	userInput := game.NewTermUserInput(os.Stdin, os.Stdout)
	g := game.NewGame(ui, userInput, 2000)
	g.CreatePlayers()
	g.Run()
}
