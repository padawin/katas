package game

import (
	"bytes"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type mockUI struct {
}

func (m mockUI) clearScreen()                     {}
func (m mockUI) showBust()                        {}
func (m mockUI) showError(err error)              {}
func (m mockUI) showScore(p player)               {}
func (m mockUI) showWinner(p player)              {}
func (m mockUI) showThrow(throw []int)            {}
func (m mockUI) showTurnScore(turnScore int)      {}
func (m mockUI) showFinalTurnScore(turnScore int) {}

type mockUserInteractions struct {
	countCallReadPlayerName     int
	countCallReadDiceToKeep     int
	countCallReadEndThrowChoice int
	playerNames                 []string
	diceToKeep                  [][]int
	endThrowChoices             []turnChoice
}

func newMockUserInteractions(playerNames []string, diceToKeep [][]int, endThrowChoices []turnChoice) *mockUserInteractions {
	return &mockUserInteractions{0, 0, 0, playerNames, diceToKeep, endThrowChoices}
}

func (m *mockUserInteractions) readPlayerName(prompt string) string {
	if m.countCallReadPlayerName >= len(m.playerNames) {
		panic("Not enough fixtures in readPlayerName")
	}
	res := m.playerNames[m.countCallReadPlayerName]
	m.countCallReadPlayerName++
	return res
}
func (m *mockUserInteractions) readDiceToKeep(throwSize int) []int {
	if m.countCallReadDiceToKeep >= len(m.diceToKeep) {
		panic("Not enough fixtures in readDiceToKeep")
	}
	res := m.diceToKeep[m.countCallReadDiceToKeep]
	m.countCallReadDiceToKeep++
	return res
}
func (m *mockUserInteractions) readEndThrowChoice() turnChoice {
	if m.countCallReadEndThrowChoice >= len(m.endThrowChoices) {
		panic("Not enough fixtures in readEndThrowChoice")
	}
	res := m.endThrowChoices[m.countCallReadEndThrowChoice]
	m.countCallReadEndThrowChoice++
	return res
}

func Test_CreatePlayers(t *testing.T) {
	g := NewGomegaWithT(t)
	gameInstance := NewGame(
		mockUI{},
		newMockUserInteractions([]string{"John Doe", "Jack Smith"}, nil, nil),
		2000,
	)
	gameInstance.CreatePlayers()
	g.Expect(gameInstance.players[0].name).To(Equal("John Doe"))
	g.Expect(gameInstance.players[1].name).To(Equal("Jack Smith"))
}

func Test_GetWinnerNoPlayer(t *testing.T) {
	g := NewGomegaWithT(t)
	gameInstance := NewGame(mockUI{}, nil, 2000)
	winner := gameInstance.getWinner()
	g.Expect(winner).To(BeNil())
}

func Test_GetWinnerNoWinner(t *testing.T) {
	g := NewGomegaWithT(t)
	gameInstance := NewGame(
		mockUI{},
		newMockUserInteractions([]string{"John Doe", "Jack Smith"}, nil, nil),
		2000,
	)
	winner := gameInstance.getWinner()
	g.Expect(winner).To(BeNil())
}

func Test_GetWinnerWithPlayer1AsWinner(t *testing.T) {
	g := NewGomegaWithT(t)
	gameInstance := NewGame(
		mockUI{},
		newMockUserInteractions([]string{"John Doe", "Jack Smith"}, nil, nil),
		2000,
	)
	gameInstance.CreatePlayers()
	gameInstance.players[0].score = 2300
	gameInstance.players[1].score = 300
	winner := gameInstance.getWinner()
	g.Expect(winner).To(Equal(&gameInstance.players[0]))
}

func Test_GetWinnerWithPlayer2AsWinner(t *testing.T) {
	g := NewGomegaWithT(t)
	gameInstance := NewGame(
		mockUI{},
		newMockUserInteractions([]string{"John Doe", "Jack Smith"}, nil, nil),
		2000,
	)
	gameInstance.CreatePlayers()
	gameInstance.players[0].score = 300
	gameInstance.players[1].score = 2300
	winner := gameInstance.getWinner()
	g.Expect(winner).To(Equal(&gameInstance.players[1]))
}

var _ = Describe("Game.turn", func() {
	Context("When throw is bust", func() {
		var gameInstance Game
		var currentPlayer int
		var throw []int = []int{2, 3, 3, 6, 2}
		BeforeEach(func() {
			gameInstance = NewGame(
				mockUI{},
				newMockUserInteractions(nil, nil, nil),
				2000,
			)
			currentPlayer = gameInstance.currentPlayer
			gameInstance.turn(throw)
		})
		It("Changes the current player", func() {
			Expect(gameInstance.currentPlayer).To(Not(Equal(currentPlayer)))
		})
		It("The player who played got 0 points", func() {
			Expect(gameInstance.players[currentPlayer].score).To(Equal(0))
		})
		It("Resets the turn score to 0", func() {
			Expect(gameInstance.turnScore).To(Equal(0))
		})
		It("Resets the dice to the maxThrowSize, to throw all of them at the next turn", func() {
			Expect(gameInstance.throwSize).To(Equal(maxThrowSize))
		})
	})
	Context("When throw is not bust", func() {
		var gameInstance Game
		var currentPlayer int
		var playerDiceChoices [][]int
		var b bytes.Buffer
		var ui *TermUI
		Context("When the player chooses a all the dice", func() {
			Context("When the player chooses to throw again", func() {
				BeforeEach(func() {
					b = bytes.Buffer{}
					ui = NewTermUI(&b)
					playerEndTurnChoices := []turnChoice{throwAgain}
					playerDiceChoices = [][]int{{2, 1, 0, 3, 5, 4}}
					gameInstance = NewGame(
						ui,
						newMockUserInteractions(nil, playerDiceChoices, playerEndTurnChoices),
						2000,
					)
					currentPlayer = gameInstance.currentPlayer
					throw := []int{1, 1, 5, 2, 2, 2}
					gameInstance.turn(throw)
				})
				It("must not add the points to the player", func() {
					Expect(gameInstance.players[currentPlayer].score).To(Equal(0))
				})
				It("Changes not the current player", func() {
					Expect(gameInstance.currentPlayer).To(Equal(currentPlayer))
				})
				It("must add the points to game.turnScore", func() {
					Expect(gameInstance.turnScore).To(Equal(450))
				})
				It("Resets the dice to the maxThrowSize, to throw all of them at the next turn", func() {
					Expect(gameInstance.throwSize).To(Equal(maxThrowSize))
				})
			})
		})
		Context("When the player chooses a valid selection of dice", func() {
			BeforeEach(func() {
				b = bytes.Buffer{}
				ui = NewTermUI(&b)
				playerDiceChoices = [][]int{
					{0, 1, 2}, // Chooses the dice 1, 1 and 5; which is valid
				}
			})
			Context("When the player chooses to end their turn here", func() {
				BeforeEach(func() {
					playerEndTurnChoices := []turnChoice{endTurn}
					gameInstance = NewGame(
						ui,
						newMockUserInteractions(nil, playerDiceChoices, playerEndTurnChoices),
						2000,
					)
					currentPlayer = gameInstance.currentPlayer
					throw := []int{1, 1, 5, 2, 4, 2}
					gameInstance.turn(throw)
				})
				It("must add the points to the player", func() {
					Expect(gameInstance.players[currentPlayer].score).To(Equal(250))
				})
				It("Changes the current player", func() {
					Expect(gameInstance.currentPlayer).To(Not(Equal(currentPlayer)))
				})
				It("Resets the turn score to 0", func() {
					Expect(gameInstance.turnScore).To(Equal(0))
				})
				It("Resets the dice to the maxThrowSize, to throw all of them at the next turn", func() {
					Expect(gameInstance.throwSize).To(Equal(maxThrowSize))
				})
			})
			Context("When the player chooses to throw again", func() {
				BeforeEach(func() {
					playerEndTurnChoices := []turnChoice{throwAgain}
					gameInstance = NewGame(
						ui,
						newMockUserInteractions(nil, playerDiceChoices, playerEndTurnChoices),
						2000,
					)
					currentPlayer = gameInstance.currentPlayer
					throw := []int{1, 1, 5, 2, 4, 2}
					gameInstance.turn(throw)
				})
				It("must not add the points to the player", func() {
					Expect(gameInstance.players[currentPlayer].score).To(Equal(0))
				})
				It("Changes not the current player", func() {
					Expect(gameInstance.currentPlayer).To(Equal(currentPlayer))
				})
				It("must add the points to game.turnScore", func() {
					Expect(gameInstance.turnScore).To(Equal(250))
				})
				It("Resets the dice to the maxThrowSize, to throw all of them at the next turn", func() {
					Expect(gameInstance.throwSize).To(Equal(3))
				})
			})
			Context("When the player gets a bust after the second throw", func() {
				BeforeEach(func() {
					playerEndTurnChoices := []turnChoice{throwAgain}
					gameInstance = NewGame(
						mockUI{},
						newMockUserInteractions(nil, playerDiceChoices, playerEndTurnChoices),
						2000,
					)
					currentPlayer = gameInstance.currentPlayer
					throw := []int{1, 1, 5, 2, 4, 2}
					gameInstance.turn(throw)
					Expect(gameInstance.turnScore).To(Equal(250))
					throw = []int{2, 3, 4}
					gameInstance.turn(throw)
				})
				It("Changes the current player", func() {
					Expect(gameInstance.currentPlayer).To(Not(Equal(currentPlayer)))
				})
				It("The player who played got 0 points", func() {
					Expect(gameInstance.players[currentPlayer].score).To(Equal(0))
				})
				It("Resets the turn score to 0", func() {
					Expect(gameInstance.turnScore).To(Equal(0))
				})
				It("Resets the dice to the maxThrowSize, to throw all of them at the next turn", func() {
					Expect(gameInstance.throwSize).To(Equal(maxThrowSize))
				})
			})
		})
		Context("When the player chooses an invalid selection of dice, then a valid one", func() {
			BeforeEach(func() {
				b = bytes.Buffer{}
				ui = NewTermUI(&b)
				playerDiceChoices = [][]int{
					{0, 1, 3}, // Chooses the dice 1, 1 and 2; which is invalid
					{0, 1, 2}, // Chooses the dice 1, 1 and 5; which is valid
				}
				playerEndTurnChoices := []turnChoice{endTurn}
				gameInstance = NewGame(
					ui,
					newMockUserInteractions(nil, playerDiceChoices, playerEndTurnChoices),
					2000,
				)
				throw := []int{1, 1, 5, 2, 4, 2}
				gameInstance.turn(throw)
			})
			It("must display an error for the first user input", func() {
				expectedPartialOutput := `
This selection would bring you no point
This turn might bring you 250 points
`
				Expect(strings.Contains(b.String(), expectedPartialOutput)).To(BeTrue())
			})
		})
		/*
			If the player chooses valid selection of dice
				If the player chooses to throw again
					if the second throw is a bust
						[Bust case]
					else
						[..] As many times as the player throws
				If the player chooses to end turn
					check points and end turn
		*/
	})
})
