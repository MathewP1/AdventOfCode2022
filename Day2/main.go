package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Choice uint8
type Result int

const (
	Rock          Choice = 1
	Paper         Choice = 2
	Scissors      Choice = 3
	UnknownChoice Choice = 4
)

const (
	Win           Result = 6
	Draw          Result = 3
	Loss          Result = 0
	UnknownResult Result = -1
)

func charToChoice(c uint8) Choice {
	if c == 'X' || c == 'A' {
		return Rock
	}
	if c == 'Y' || c == 'B' {
		return Paper
	}
	if c == 'Z' || c == 'C' {
		return Scissors
	}
	return UnknownChoice
}

func charToResult(c uint8) Result {
	if c == 'X' {
		return Loss
	}
	if c == 'Y' {
		return Draw
	}
	if c == 'Z' {
		return Win
	}
	return UnknownResult
}

// returns map that tells the winning choice against everything
func getWinMap() map[Choice]Choice {
	return map[Choice]Choice{
		Rock:     Paper,
		Paper:    Scissors,
		Scissors: Rock,
	}
}

func getLoseMap() map[Choice]Choice {
	return map[Choice]Choice{
		Rock:     Scissors,
		Paper:    Rock,
		Scissors: Paper,
	}
}

func Task1(filePath string) (score int) {
	score = 0
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	winMap := getWinMap()
	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if line == "" {
			continue
		}
		splitString := strings.Split(line, " ")
		if len(splitString) != 2 && len(splitString[1]) != 1 {
			panic("Unexpected parsing error")
		}

		myOpponent := charToChoice(splitString[0][0])
		myChoice := charToChoice(splitString[1][0])

		winAgainstOpponent := winMap[myOpponent]

		if myChoice == myOpponent {
			fmt.Println("Draw!")
			score += 3
		} else if myChoice == winAgainstOpponent {
			fmt.Println("Win!")
			score += 6
		} else {
			fmt.Println("Loss!")
		}

		score += int(myChoice)
	}
	return score
}

func Task2(filePath string) (score int) {
	score = 0
	winMap := getWinMap()
	loseMap := getLoseMap()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if line == "" {
			continue
		}
		splitString := strings.Split(line, " ")
		if len(splitString) != 2 && len(splitString[1]) != 1 {
			panic("Unexpected parsing error")
		}

		myOpponent := charToChoice(splitString[0][0])
		neededResult := charToResult(splitString[1][0])

		myChoice := UnknownChoice
		if neededResult == Win {
			myChoice = winMap[myOpponent]
		} else if neededResult == Draw {
			myChoice = myOpponent
		} else {
			myChoice = loseMap[myOpponent]
		}

		score += int(myChoice)
		score += int(neededResult)
	}

	return score
}

func main() {
	score1 := Task1("input.txt")
	fmt.Println("Task1 score is ", score1)
	score2 := Task2("input.txt")
	fmt.Println("Task2 score is ", score2)

}
