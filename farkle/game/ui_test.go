package game

import (
	"bytes"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ShowBust(t *testing.T) {
	g := NewGomegaWithT(t)
	b := bytes.Buffer{}
	ui := NewTermUI(&b)
	ui.showBust()
	g.Expect(b.String()).To(Equal("Bust!\n"))
}

func Test_ShowScore(t *testing.T) {
	g := NewGomegaWithT(t)
	b := bytes.Buffer{}
	ui := NewTermUI(&b)
	p := player{"John Doe", 42}
	ui.showScore(p)
	g.Expect(b.String()).To(Equal("John Doe, you have 42 points\n"))
}

func Test_ShowWinner(t *testing.T) {
	g := NewGomegaWithT(t)
	b := bytes.Buffer{}
	ui := NewTermUI(&b)
	p := player{"John Doe", 42}
	ui.showWinner(p)
	g.Expect(b.String()).To(Equal("Congratulation! John Doe, you won with 42 points!\n"))
}

func Test_ShowTurnScore(t *testing.T) {
	g := NewGomegaWithT(t)
	b := bytes.Buffer{}
	ui := NewTermUI(&b)
	ui.showTurnScore(42)
	g.Expect(b.String()).To(Equal("This turn might bring you 42 points\n"))
}

func Test_ShowFinalTurnScore(t *testing.T) {
	g := NewGomegaWithT(t)
	b := bytes.Buffer{}
	ui := NewTermUI(&b)
	ui.showFinalTurnScore(42)
	g.Expect(b.String()).To(Equal("You got 42 points!\n"))
}

func Test_ShowThrow(t *testing.T) {
	g := NewGomegaWithT(t)
	b := bytes.Buffer{}
	ui := NewTermUI(&b)
	throw := []int{1, 1, 3, 5, 2}
	ui.showThrow(throw)
	expected := "Throw:\n+-------+-------+-------+-------+-------+\n| DIE \033[5;32m1\033[0m | DIE \033[5;32m2\033[0m | DIE \033[5;32m3\033[0m | DIE \033[5;32m4\033[0m | DIE \033[5;32m5\033[0m |\n+-------+-------+-------+-------+-------+\n|   1   |   1   |   3   |   5   |   2   |\n+-------+-------+-------+-------+-------+\n"
	g.Expect(b.String()).To(Equal(expected))
}

func Test_ShowError(t *testing.T) {
	g := NewGomegaWithT(t)
	b := bytes.Buffer{}
	ui := NewTermUI(&b)
	ui.showError(fmt.Errorf("An error"))
	g.Expect(b.String()).To(Equal("An error\n"))
}
