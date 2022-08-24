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

func calculateScoreForTwoRolls(rolls ...string) (sum uint, err error) {
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

	if v1 > 9 || v2 > 9 {
		return 0, errors.New("invalid score")
	}

	return v1 + v2, nil
}

func parseStringToInt(s string) (uint, error) {
	s = strings.ToLower(s)

	switch s {
	case "x":
		return 10, nil
	case "-":
		return 0, nil
	}

	i, err := strconv.Atoi(s)

	return uint(i), err
}

type Frame struct {
	score uint
	rolls []string
}

func (f *Frame) calculateScore() {
	var value uint
	var err error

	switch len(f.rolls) {
	case 1:
		value, err = parseStringToInt(f.rolls[0])
	case 2:
		value, err = calculateScoreForTwoRolls(f.rolls...)
	default:
		err = errors.New("too many args")
	}

	if err == nil {
		f.score = value
	}
}
