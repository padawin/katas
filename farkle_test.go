package farkle

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_RollNDiceReturnsNInt(t *testing.T) {
	g := NewGomegaWithT(t)
	var res []int
	var err error
	res, err = RollNDice(7)
	g.Expect(len(res)).To(Equal(7))
	g.Expect(err).To(BeNil())
	res, err = RollNDice(0)
	g.Expect(err).To(BeNil())
}

func Test_RollNDiceReturnsErrorIfNIsNegative(t *testing.T) {
	g := NewGomegaWithT(t)
	res, err := RollNDice(-7)
	g.Expect(len(res)).To(Equal(0))
	g.Expect(err).To(Equal(fmt.Errorf("the number of dice to roll must be positive")))
}

func Test_CalculateScoreReturns0For0values(t *testing.T) {
	g := NewGomegaWithT(t)
	res := CalculateScore([]int{})
	g.Expect(res).To(Equal(0))
}

func Test_CalculateScoreForDifferentValues(t *testing.T) {
	type testData struct {
		play  []int
		score int
	}
	fixtures := []testData{
		{[]int{1}, 100},
		{[]int{1, 1}, 200},
		{[]int{1, 1, 1}, 1000},
		{[]int{1, 1, 1, 1}, 2000},
		{[]int{1, 1, 1, 1, 1}, 4000},
		{[]int{1, 1, 1, 1, 1, 1}, 8000},
		{[]int{2}, 0},
		{[]int{2, 2}, 0},
		{[]int{2, 2, 2}, 200},
		{[]int{2, 2, 2, 2}, 400},
		{[]int{2, 2, 2, 2, 2}, 800},
		{[]int{2, 2, 2, 2, 2, 2}, 1600},
		{[]int{3}, 0},
		{[]int{3, 3}, 0},
		{[]int{3, 3, 3}, 300},
		{[]int{3, 3, 3, 3}, 600},
		{[]int{3, 3, 3, 3, 3}, 1200},
		{[]int{3, 3, 3, 3, 3, 3}, 2400},
		{[]int{4}, 0},
		{[]int{4, 4}, 0},
		{[]int{4, 4, 4}, 400},
		{[]int{4, 4, 4, 4}, 800},
		{[]int{4, 4, 4, 4, 4}, 1600},
		{[]int{4, 4, 4, 4, 4, 4}, 3200},
		{[]int{5}, 50},
		{[]int{5, 5}, 100},
		{[]int{5, 5, 5}, 500},
		{[]int{5, 5, 5, 5}, 1000},
		{[]int{5, 5, 5, 5, 5}, 2000},
		{[]int{5, 5, 5, 5, 5, 5}, 4000},
		{[]int{6}, 0},
		{[]int{6, 6}, 0},
		{[]int{6, 6, 6}, 600},
		{[]int{6, 6, 6, 6}, 1200},
		{[]int{6, 6, 6, 6, 6}, 2400},
		{[]int{6, 6, 6, 6, 6, 6}, 4800},
		// combinations
		{[]int{1, 5, 5}, 200},
		{[]int{1, 1, 5}, 250},
		{[]int{4, 4, 4, 1, 5}, 550},
		// straight
		{[]int{1, 2, 3, 4, 5, 6}, 1200},
		{[]int{3, 1, 6, 4, 2, 5}, 1200},
	}
	for _, fixture := range fixtures {
		score := CalculateScore(fixture.play)
		if score != fixture.score {
			t.Errorf("Score: With play %v, expected %d but got %d ", fixture.play, fixture.score, score)
		}
	}
}
