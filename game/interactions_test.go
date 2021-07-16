package game

import (
	"bytes"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadPlayerName(t *testing.T) {
	g := NewGomegaWithT(t)
	input := bytes.Buffer{}
	output := bytes.Buffer{}
	ui := NewTermUserInput(&input, &output)
	input.WriteString("John Doe")
	name := ui.readPlayerName("one")
	g.Expect(name).To(Equal("John Doe"))
	g.Expect(output.String()).To(Equal("Player one name: "))
}

func Test_ReadPlayerNameWithSpaces(t *testing.T) {
	g := NewGomegaWithT(t)
	input := bytes.Buffer{}
	output := bytes.Buffer{}
	ui := NewTermUserInput(&input, &output)
	input.WriteString("   John Doe   ")
	name := ui.readPlayerName("one")
	g.Expect(name).To(Equal("John Doe"))
	g.Expect(output.String()).To(Equal("Player one name: "))
}

func Test_ReadPlayerNameWithEmptyName(t *testing.T) {
	g := NewGomegaWithT(t)
	input := bytes.Buffer{}
	output := bytes.Buffer{}
	ui := NewTermUserInput(&input, &output)
	input.WriteString("\n")
	input.WriteString("John Doe")
	name := ui.readPlayerName("one")
	expectedOutput := []string{
		"Player one name: ",
		"Player one name: ",
	}
	g.Expect(name).To(Equal("John Doe"))
	g.Expect(output.String()).To(Equal(strings.Join(expectedOutput, "")))
}

func Test_ReadDiceToKeep(t *testing.T) {
	g := NewGomegaWithT(t)
	input := bytes.Buffer{}
	output := bytes.Buffer{}
	ui := NewTermUserInput(&input, &output)
	input.WriteString("1 4 3 2")
	diceIndices := ui.readDiceToKeep(5)
	g.Expect(output.String()).To(Equal("Select \033[5;32mdice number(s)\033[0m to keep: "))
	g.Expect(diceIndices).To(Equal([]int{0, 3, 2, 1}))
}

func Test_ReadDiceToKeepInvalidInput(t *testing.T) {
	g := NewGomegaWithT(t)
	input := bytes.Buffer{}
	output := bytes.Buffer{}
	input.WriteString("\n")
	input.WriteString("   \n")
	input.WriteString("foo bar\n")
	input.WriteString("1 2 3 foo bar\n")
	input.WriteString("-1 2 3\n")
	input.WriteString("1 4 0 2\n")
	input.WriteString("1 4 3 2 8\n")
	input.WriteString("1 4 3 2")
	ui := NewTermUserInput(&input, &output)
	diceIndices := ui.readDiceToKeep(5)
	expectedOutput := []string{
		"Select \033[5;32mdice number(s)\033[0m to keep: ",
		"Please choose at least one dice\n",
		"Select \033[5;32mdice number(s)\033[0m to keep: ",
		"Please choose at least one dice\n",
		"Select \033[5;32mdice number(s)\033[0m to keep: ",
		"foo is not a valid die number\n",
		"Select \033[5;32mdice number(s)\033[0m to keep: ",
		"foo is not a valid die number\n",
		"Select \033[5;32mdice number(s)\033[0m to keep: ",
		"-1 is not a valid die number\n",
		"Select \033[5;32mdice number(s)\033[0m to keep: ",
		"0 is not a valid die number\n",
		"Select \033[5;32mdice number(s)\033[0m to keep: ",
		"8 is not a valid die number\n",
		"Select \033[5;32mdice number(s)\033[0m to keep: ",
	}
	g.Expect(output.String()).To(Equal(strings.Join(expectedOutput, "")))
	g.Expect(diceIndices).To(Equal([]int{0, 3, 2, 1}))
}

func Test_ReadDiceToKeepDuplicatedNumbers(t *testing.T) {
	g := NewGomegaWithT(t)
	input := bytes.Buffer{}
	output := bytes.Buffer{}
	ui := NewTermUserInput(&input, &output)
	input.WriteString("1 1 1 1")
	diceIndices := ui.readDiceToKeep(2)
	g.Expect(output.String()).To(Equal("Select \033[5;32mdice number(s)\033[0m to keep: "))
	g.Expect(diceIndices).To(Equal([]int{0}))
}

func Test_ReadEndThrowChoiceEndTurn(t *testing.T) {
	g := NewGomegaWithT(t)
	input := bytes.Buffer{}
	output := bytes.Buffer{}
	ui := NewTermUserInput(&input, &output)
	input.WriteString("e")
	choice := ui.readEndThrowChoice()
	g.Expect(output.String()).To(Equal("Do you want to (T)hrow again, or to (E)nd turn (t/e)? "))
	g.Expect(choice).To(Equal(turnChoice(endTurn)))
}

func Test_ReadEndThrowChoiceThrowAgain(t *testing.T) {
	g := NewGomegaWithT(t)
	input := bytes.Buffer{}
	output := bytes.Buffer{}
	ui := NewTermUserInput(&input, &output)
	input.WriteString("t")
	choice := ui.readEndThrowChoice()
	g.Expect(output.String()).To(Equal("Do you want to (T)hrow again, or to (E)nd turn (t/e)? "))
	g.Expect(choice).To(Equal(turnChoice(throwAgain)))
}

func Test_ReadEndThrowChoiceInvalid(t *testing.T) {
	g := NewGomegaWithT(t)
	input := bytes.Buffer{}
	output := bytes.Buffer{}
	ui := NewTermUserInput(&input, &output)
	input.WriteString("hello\n")
	input.WriteString("e")
	choice := ui.readEndThrowChoice()
	expectedOutput := []string{
		"Do you want to (T)hrow again, or to (E)nd turn (t/e)? ",
		"hello is not a valid choice\n",
		"Do you want to (T)hrow again, or to (E)nd turn (t/e)? ",
	}
	g.Expect(output.String()).To(Equal(strings.Join(expectedOutput, "")))
	g.Expect(choice).To(Equal(turnChoice(endTurn)))
}
