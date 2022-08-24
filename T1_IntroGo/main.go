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

func calculateFrameScore(rolls ...string) (uint, error) {
	switch len(rolls) {
	case 1:
		return parseStringToInt(rolls[0])
	case 2:
		return calculateScoreForTwoRolls(rolls...)
	}

	return 0, errors.New("too many args")
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
