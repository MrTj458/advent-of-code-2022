package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	rockPoints       = 1
	paperPoints      = 2
	scissorPoints    = 3
	winPoints        = 6
	tiePoints        = 3
	losePoints       = 0
	playerLose       = "X"
	playerTie        = "Y"
	playerWin        = "Z"
	opponentRock     = "A"
	opponentPaper    = "B"
	opponentScissors = "C"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0
	for scanner.Scan() {
		row := scanner.Text()
		opponentChoice := strings.Split(row, " ")[0]
		playerChoice := strings.Split(row, " ")[1]

		switch playerChoice {
		case playerLose:
			switch opponentChoice {
			case opponentRock:
				total += scissorPoints + losePoints
			case opponentPaper:
				total += rockPoints + losePoints
			case opponentScissors:
				total += paperPoints + losePoints
			}
		case playerTie:
			switch opponentChoice {
			case opponentRock:
				total += rockPoints + tiePoints
			case opponentPaper:
				total += paperPoints + tiePoints
			case opponentScissors:
				total += scissorPoints + tiePoints
			}
		case playerWin:
			switch opponentChoice {
			case opponentRock:
				total += paperPoints + winPoints
			case opponentPaper:
				total += scissorPoints + winPoints
			case opponentScissors:
				total += rockPoints + winPoints
			}
		}
	}

	fmt.Println("Points:", total)
}
