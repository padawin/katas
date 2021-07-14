package game

import (
	"fmt"
	"io"
	"strings"
)

type UI interface {
	showBust()
	showScore(player player)
	showTurnScore(turnScore int)
	showThrow(throw []int)
	showError(err error)
}

type TermUI struct {
	output io.Writer
}

func NewTermUI(output io.Writer) *TermUI {
	return &TermUI{output}
}

func (u *TermUI) showBust() {
	fmt.Fprintln(u.output, "Bust!")
}

func (u *TermUI) showScore(player player) {
	fmt.Fprintf(u.output, "%s, you have %d points\n", player.name, player.score)
}

func (u *TermUI) showTurnScore(turnScore int) {
	fmt.Fprintf(u.output, "This turn might bring you %d points\n", turnScore)
}

func (u *TermUI) showThrow(throw []int) {
	fmt.Fprintln(u.output, "Throw:")
	//Prints something along the lines of:
	//+-------+-------+-------+-------+-------+-------|
	//| DIE 1 | DIE 2 | DIE 3 | DIE 4 | DIE 5 | DIE 6 |
	//+-------+-------+-------+-------+-------+-------|
	//|   5   |   5   |   5   |   5   |   5   |   5   |
	//+-------+-------+-------+-------+-------+-------|
	fmt.Fprint(u.output, strings.Repeat("+-------", len(throw)), "+\n")
	for i := range throw {
		fmt.Fprintf(u.output, "| DIE \033[5;32m%d\033[0m ", i+1)
	}
	fmt.Fprintln(u.output, "|")
	fmt.Fprint(u.output, strings.Repeat("+-------", len(throw)), "+\n")
	for _, value := range throw {
		fmt.Fprintf(u.output, "|   %d   ", value)
	}
	fmt.Fprintln(u.output, "|")
	fmt.Fprint(u.output, strings.Repeat("+-------", len(throw)), "+\n")
}

func (u *TermUI) showError(err error) {
	fmt.Fprintln(u.output, err)
}
