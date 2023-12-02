package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"os"
)

type Round struct {
	Red int
	Green int
	Blue int
}

type BagState = Round

type Game struct {
	Id int
	Rounds []Round
}

func (g Game) String() string {
	roundsStr := []string{}
	for _, round := range g.Rounds {
		roundStr := []string{}
		if round.Red > 0 {
			roundStr = append(roundStr, fmt.Sprintf("%d red", round.Red))
		}
		if round.Blue > 0 {
			roundStr = append(roundStr, fmt.Sprintf("%d blue", round.Blue))
		}
		if round.Green > 0 {
			roundStr = append(roundStr, fmt.Sprintf("%d green", round.Green))
		}
		roundsStr = append(roundsStr, strings.Join(roundStr, ", "))
	}
	return fmt.Sprintf("Game %d: %s", g.Id, strings.Join(roundsStr, "; "))
}


func parseGame(line string) (Game, error) {
	mainParts := strings.Split(line, ": ")
	if len(mainParts) != 2 {
		return Game{}, fmt.Errorf("Expected two parts w/ colon")
	}
	gameIdParts := strings.Split(mainParts[0], " ")
	if len(gameIdParts) != 2 {
		return Game{}, fmt.Errorf("Error parsing game ID")
	}
	gameId, err := strconv.Atoi(gameIdParts[1])
	if err != nil {
		return Game{}, fmt.Errorf("Error parsing numeric ID")
	}

	var rounds []Round
	roundStrings := strings.Split(mainParts[1], "; ")
	for _, roundString := range(roundStrings) {
		var round Round
		elements := strings.Split(roundString, ", ")
		for _, element := range(elements) {
			numColor := strings.Split(element, " ")
			if len(numColor) != 2 {
				return Game{}, fmt.Errorf("Expected number and color")
			}
			num, err := strconv.Atoi(numColor[0])
			if err != nil {
				return Game{}, fmt.Errorf("Couldn't parse num")
			}
			switch numColor[1] {
			case "red":
				round.Red = num
			case "blue":
				round.Blue = num
			case "green":
				round.Green = num
			default:
				return Game{}, fmt.Errorf("Unrecognized color %s", numColor[1])
			}
		}
		rounds = append(rounds, round)
	}

	return Game{gameId, rounds}, nil
}

func isPossibleBagState(bag BagState, game Game) bool {
	for _, round := range(game.Rounds) {
		if round.Red > bag.Red || round.Blue > bag.Blue || round.Green > bag.Green {
			return false
		}
	}
	return true
}

func powerOfMinCubes(game Game) int {
	var r, g, b int
	for _, round := range(game.Rounds) {
		r = max(r, round.Red)
		g = max(g, round.Green)
		b = max(b, round.Blue)
	}
	return r * g * b
}

func parseTest(input string) {
	game, err := parseGame(input)
	if err != nil {
		fmt.Printf("Error!  %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Parsing '%s'\n", input)
	fmt.Printf("Parsed  '%s'\n\n", game.String())
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lineNum := 0
	sum := 0
	var line string

	for scanner.Scan() {
		line = scanner.Text()
		game, err := parseGame(line)
		if err != nil {
			fmt.Printf("Got error %s\n", err)
			os.Exit(1)
		}
		// Part 1:
		// if isPossibleBagState(BagState{12, 13, 14}, game) {
		// 	sum += game.Id
		// }

		// Part 2:
		sum += powerOfMinCubes(game)
		lineNum += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading lines: %v", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", sum)
}
