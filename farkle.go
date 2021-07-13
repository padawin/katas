package farkle

import (
	"fmt"
	"math"
	"math/rand"
)

const maxDice = 6

// Defines the score of some values when they appear in 1 or 2 dice only
var scoreUnit = map[int]int{
	1: 100,
	5: 50,
}

// straight == all dice have a different value
const scoreStraight = 1200

// Base for the score for numbers appearing 3 or more times
const baseScoreThreeOfAKind = 100

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
		valueScore := calculateValueScore(value, count)
		if valueScore == 0 {
			score = 0
			break
		}
		score += valueScore
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
	} else if base, found = scoreUnit[value]; found {
		score += base * count
	} else {
		score = 0
	}
	return score
}
