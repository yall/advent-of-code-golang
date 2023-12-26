package main

import (
	"fmt"
	"os"

	"github.com/yall/advent-of-code-golang/pkg/utils"
)

func main() {
	lines := utils.LoadInput(os.Args[1])
	var sum int
	for _, line := range lines {
		digits := extractDigits(line)
		//fmt.Println(line, ":", digits)
		value := digitsValue(digits)
		sum += value
	}
	fmt.Println(sum)
}

var digitValuesMap = map[string]int{
	"zero":  0,
	"0":     0,
	"one":   1,
	"1":     1,
	"two":   2,
	"2":     2,
	"three": 3,
	"3":     3,
	"four":  4,
	"4":     4,
	"five":  5,
	"5":     5,
	"six":   6,
	"6":     6,
	"seven": 7,
	"7":     7,
	"eight": 8,
	"8":     8,
	"nine":  9,
	"9":     9,
}

func checkDigitWord(line string, i int, word string) bool {
	for j, c := range word {
		if i+j >= len(line) {
			return false
		}
		//fmt.Println(line[i+j], c)
		if rune(line[i+j]) != c {
			return false
		}
	}
	return true
}

func extractDigits(line string) []int {
	var digits []int
	for i, _ := range line {
		for word, v := range digitValuesMap {
			if checkDigitWord(line, i, word) {
				digits = append(digits, v)
			}
		}
	}
	return digits
}

func digitsValue(digits []int) int {
	f, l := digits[0], digits[len(digits)-1]
	return 10*f + l
}
