package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World")
}

func calculateFrameScore(rolls ...string) (int, error) {
	switch len(rolls) {
	case 1:
		return parseStringToInt(rolls[0])
	case 2:
		return calculateScoreForTwoRolls(rolls...)
	}

	return 0, errors.New("too many args")
}

func calculateScoreForTwoRolls(rolls ...string) (sum int, err error) {
	v1, err := parseStringToInt(rolls[0])

	if err != nil {
		return 0, err
	}
	if rolls[1] == "/" {
		return 10, nil
	}

	v2, err := parseStringToInt(rolls[1])
	if err != nil {
		return 0, err
	}
	return v1 + v2, nil
}

func parseStringToInt(s string) (int, error) {
	if strings.ToLower(s) == "x" {
		return 10, nil
	}

	return strconv.Atoi(s)
}
