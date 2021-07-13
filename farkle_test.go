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
