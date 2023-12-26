package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yall/advent-of-code-golang/pkg/utils"
)

func main() {
	lines := utils.LoadInput(os.Args[1])
	var sum int
	for _, line := range lines {
		digits := extractDigits(line)
		value := digitsValue(digits)
		sum += value
	}
	fmt.Println(sum)
}

func extractDigits(line string) []int {
	var digits []int
	for _, c := range line {
		if c >= '0' && c <= '9' {
			i, _ := strconv.Atoi(string(c))
			digits = append(digits, i)
		}
	}
	return digits
}

func digitsValue(digits []int) int {
	f, l := digits[0], digits[len(digits)-1]
	return 10*f + l
}
