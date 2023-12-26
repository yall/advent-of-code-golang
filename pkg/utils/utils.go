package utils

import (
	"os"
	"strings"
)

func LoadInput(path string) []string {
	dat, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	return strings.Split(string(dat), "\n")
}
