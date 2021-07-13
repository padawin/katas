package farkle

import (
	"fmt"
	"math"
	"math/rand"
)

const maxDice = 6

// Score for numbers appearing 1 or 2 times
const basePointsUnit = 0

// straight == all dice have a different value
const scoreStraight = 1200

// Base for the score for numbers appearing 3 or more times
const baseScoreThreeOfAKind = 100

// Override score for numbers appearing 1 or 2 times
var specialScoreUnit = map[int]int{
	1: 100,
	5: 50,
}

// Override score for numbers appearing 3 or more times
var specialScoreThreeOfAKind = map[int]int{
	1: 1000,
}

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

func CalculateScore(values []int) int {
	// Group per values
	countsPerValue := map[int]int{}
	for _, value := range values {
		countsPerValue[value]++
	}
	if len(countsPerValue) == maxDice {
		return scoreStraight
	}
	score := 0
	for value, count := range countsPerValue {
		score += calculateValueScore(value, count)
	}
	return score
}

func calculateValueScore(value, count int) int {
	score := 0
	var base int
	var found bool
	if count >= 3 {
		if base, found = specialScoreThreeOfAKind[value]; !found {
			base = value * baseScoreThreeOfAKind
		}
		score += base * int(math.Pow(2, float64(count-3)))
	} else {
		if base, found = specialScoreUnit[value]; !found {
			base = basePointsUnit
		}
		score += base * count
	}
	return score
}
