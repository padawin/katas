package farkle

import (
	"fmt"
	"math/rand"
)

const maxDice = 6

func RollNDice(n int) ([]int, error) {
	res := []int{}
	if n < 0 {
		return res, fmt.Errorf("the number of dice to roll must be positive")
	}
	for n > 0 {
		res = append(res, rand.Intn(maxDice)+1)
		n--
	}
	return res, nil
}
