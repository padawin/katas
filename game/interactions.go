package game

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type UserInput interface {
	readPlayerName(prompt string) string
	readDiceToKeep() []int
	readEndThrowChoice() turnChoice
}

type TermUserInput struct {
	input  io.Reader
	output io.Writer
}

func NewTermUserInput(input io.Reader, output io.Writer) *TermUserInput {
	return &TermUserInput{input, output}
}

func (i *TermUserInput) readPlayerName(prompt string) string {
	reader := bufio.NewReader(i.input)
	prompt = fmt.Sprintf("Player %s name: ", prompt)
	var name string
	for name == "" {
		fmt.Fprint(i.output, prompt)
		name, _ = reader.ReadString('\n')
		name = strings.Trim(name, " \n")
	}
	return name
}

func (i *TermUserInput) readDiceToKeep() []int {
	reader := bufio.NewReader(i.input)
	var selection string
	var dice []int
	var err error
	for {
		fmt.Fprint(i.output, "Select \033[5;32mdice number(s)\033[0m to keep: ")
		selection, _ = reader.ReadString('\n')
		selection = strings.Trim(selection, " \n")
		if selection == "" {
			fmt.Fprintln(i.output, "Please choose at least one dice")
			continue
		}
		dice, err = _stringToIndices(selection)
		if err != nil {
			fmt.Fprintln(i.output, err)
		} else {
			break
		}
	}

	return _removeDuplicateValues(dice)
}

func _stringToIndices(selection string) ([]int, error) {
	values := strings.Split(selection, " ")
	res := []int{}
	for _, valStr := range values {
		val, err := strconv.Atoi(valStr)
		if err != nil || val <= 0 {
			return res, fmt.Errorf("%s is not a valid die number", valStr)
		}
		// The user reads values from 1 to n, so we convert them in 0-based values
		res = append(res, val-1)
	}
	return res, nil
}

func _removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (i *TermUserInput) readEndThrowChoice() turnChoice {
	reader := bufio.NewReader(i.input)
	var selection string
	for {
		fmt.Fprint(i.output, "Do you want to (T)hrow again, or to (E)nd turn (t/e)? ")
		selection, _ = reader.ReadString('\n')
		selection = strings.Trim(selection, "\n")
		if selection != "e" && selection != "t" {
			fmt.Fprintf(i.output, "%s is not a valid choice\n", selection)
		} else {
			break
		}
	}
	return turnChoice(selection)
}
