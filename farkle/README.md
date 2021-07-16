# Falkle

Write a class Greed with a score() method that accepts an array of die values (up to 6). Scoring rules are as follows:
Farkle is a dice game. The rules can be found here: https://en.wikipedia.org/wiki/Greed_%28dice_game%29

## Scoring

    A single one (100)
    A single five (50)
    Triple ones [1,1,1] (1000)
    Triple twos [2,2,2] (200)
    Triple threes [3,3,3] (300)
    Triple fours [4,4,4] (400)
    Triple fives [5,5,5] (500)
    Triple sixes [6,6,6] (600)

    Four-of-a-kind (Multiply Triple Score by 2)
    Five-of-a-kind (Multiply Triple Score by 4)

    Six-of-a-kind (Multiply Triple Score by 8)

    Three Pairs [2,2,3,3,4,4] (800)

    Straight [1,2,3,4,5,6] (1200)

## Usage

### Run tests

	go test ./...

### Run tests with code coverage

	go test ./... -coverprofile=/tmp/coverage.out
	go tool cover -html=/tmp/coverage.out

### Play game

	go run main.go
