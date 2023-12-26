package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/yall/advent-of-code-golang/pkg/utils"
)

type (
	Game struct {
		Id    int
		Hands []map[string]int
	}
)

const maxRed = 12
const maxGreen = 13
const maxBlue = 14

func main() {
	lines := utils.LoadInput(os.Args[1])

	result := solve(lines)
	fmt.Println(result)
}

func solve(lines []string) int {
	var result int

	for _, line := range lines {
		game := parseGame(line)
		//fmt.Println(game)
		if isGameValid(game) {
			result += game.Id
		}
	}

	return result
}

func isGameValid(game *Game) bool {
	for _, hand := range game.Hands {
		if !isHandValid(hand) {
			return false
		}
	}
	return true
}

func isHandValid(hand map[string]int) bool {

	return isThrowPossible("red", hand["red"]) &&
		isThrowPossible("green", hand["green"]) &&
		isThrowPossible("blue", hand["blue"])
}

func isThrowPossible(color string, count int) bool {
	switch color {
	case "red":
		return count <= maxRed
	case "green":
		return count <= maxGreen
	case "blue":
		return count <= maxBlue
	default:
		return false
	}
}

func parseGame(line string) *Game {
	parts := strings.Split(line, ": ")
	id, _ := strconv.Atoi(parts[0][5:])

	var hands = make([]map[string]int, 0)
	for _, hand := range strings.Split(parts[1], "; ") {
		handMap := make(map[string]int)
		cards := strings.Split(hand, ", ")
		for _, card := range cards {
			cardParts := strings.Split(card, " ")
			cardValue, _ := strconv.Atoi(cardParts[0])
			handMap[cardParts[1]] = cardValue
		}
		hands = append(hands, handMap)
	}

	return &Game{Id: id, Hands: hands}
}
