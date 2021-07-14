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
		// Selecting a combination worth nothing along with some marking points
		// invalidates the score
		{[]int{1, 1, 3}, 0},
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

func Test_IsBustWithNoValueShouldReturnAnError(t *testing.T) {
	g := NewGomegaWithT(t)
	res, err := IsBust([]int{})
	g.Expect(res).To(BeNil())
	g.Expect(err).To(Not(BeNil()))
}

func Test_IsBust(t *testing.T) {
	g := NewGomegaWithT(t)
	type testData struct {
		play   []int
		isBust bool
	}
	fixtures := []testData{
		{[]int{1}, false},
		{[]int{1, 1}, false},
		{[]int{1, 1, 1}, false},
		{[]int{1, 1, 1, 1}, false},
		{[]int{1, 1, 1, 1, 1}, false},
		{[]int{1, 1, 1, 1, 1, 1}, false},
		{[]int{2, 2, 2}, false},
		{[]int{2, 2, 2, 2}, false},
		{[]int{2, 2, 2, 2, 2}, false},
		{[]int{2, 2, 2, 2, 2, 2}, false},
		{[]int{3, 3, 3}, false},
		{[]int{3, 3, 3, 3}, false},
		{[]int{3, 3, 3, 3, 3}, false},
		{[]int{3, 3, 3, 3, 3, 3}, false},
		{[]int{4, 4, 4}, false},
		{[]int{4, 4, 4, 4}, false},
		{[]int{4, 4, 4, 4, 4}, false},
		{[]int{4, 4, 4, 4, 4, 4}, false},
		{[]int{5}, false},
		{[]int{5, 5}, false},
		{[]int{5, 5, 5}, false},
		{[]int{5, 5, 5, 5}, false},
		{[]int{5, 5, 5, 5, 5}, false},
		{[]int{5, 5, 5, 5, 5, 5}, false},
		{[]int{6, 6, 6}, false},
		{[]int{6, 6, 6, 6}, false},
		{[]int{6, 6, 6, 6, 6}, false},
		{[]int{6, 6, 6, 6, 6, 6}, false},
		// combinations
		{[]int{1, 5, 5}, false},
		{[]int{1, 1, 5}, false},
		{[]int{4, 4, 4, 1, 5}, false},
		// even though 3 is worth nothing, [1, 1] makes points, so it is not a
		// bust
		{[]int{1, 1, 3}, false},
		// straight
		{[]int{1, 2, 3, 4, 5, 6}, false},
		{[]int{3, 1, 6, 4, 2, 5}, false},
		// Bust cases
		{[]int{2}, true},
		{[]int{3}, true},
		{[]int{4}, true},
		{[]int{6}, true},
		{[]int{2, 2}, true},
		{[]int{3, 3}, true},
		{[]int{4, 4}, true},
		{[]int{6, 6}, true},
		// Combination of busts
		{[]int{6, 6, 3, 2, 2, 4}, true},
	}
	for _, fixture := range fixtures {
		isBust, err := IsBust(fixture.play)
		g.Expect(err).To(BeNil())
		if *isBust != fixture.isBust {
			t.Errorf("Bust: With play %v, expected %v but got %v ", fixture.play, fixture.isBust, *isBust)
		}
	}
}
