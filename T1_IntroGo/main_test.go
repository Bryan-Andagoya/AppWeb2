package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert.NotPanics(t, main)
}

func TestParseStringToInt(t *testing.T) {
	v, err := parseStringToInt("4")
	assert.Equal(t, uint(4), v)
	assert.Nil(t, err)
	v, err = parseStringToInt("abc")
	assert.Equal(t, uint(0), v)
	assert.NotNil(t, err)
	v, err = parseStringToInt("x")
	assert.Equal(t, uint(10), v)
	assert.Nil(t, err)
	v, err = parseStringToInt("X")
	assert.Equal(t, uint(10), v)
	assert.Nil(t, err)
	v, err = parseStringToInt("-")
	assert.Equal(t, uint(0), v)
	assert.Nil(t, err)
}

func TestCalculateScoreForTwoRolls(t *testing.T) {
	v, err := calculateScoreForTwoRolls("4", "4")
	assert.Equal(t, uint(8), v)
	assert.Nil(t, err)
	v, err = calculateScoreForTwoRolls("abc", "4")
	assert.Equal(t, uint(0), v)
	assert.NotNil(t, err)
	v, err = calculateScoreForTwoRolls("4", "abc")
	assert.Equal(t, uint(0), v)
	assert.NotNil(t, err)
	v, err = calculateScoreForTwoRolls("4", "/")
	assert.Equal(t, uint(10), v)
	assert.Nil(t, err)
	v, err = calculateScoreForTwoRolls("/", "4")
	assert.Equal(t, uint(0), v)
	assert.NotNil(t, err)
	v, err = calculateScoreForTwoRolls("/", "/")
	assert.Equal(t, uint(0), v)
	assert.NotNil(t, err)
	v, err = calculateScoreForTwoRolls("abc", "/")
	assert.Equal(t, uint(0), v)
	assert.NotNil(t, err)
	v, err = calculateScoreForTwoRolls("9", "-")
	assert.Equal(t, uint(9), v)
	assert.Nil(t, err)
	v, err = calculateScoreForTwoRolls("-", "2")
	assert.Equal(t, uint(2), v)
	assert.Nil(t, err)
	v, err = calculateScoreForTwoRolls("-", "-")
	assert.Equal(t, uint(0), v)
	assert.Nil(t, err)
	v, err = calculateScoreForTwoRolls("abc", "-")
	assert.Equal(t, uint(0), v)
	assert.NotNil(t, err)
	v, err = calculateScoreForTwoRolls("-", "abc")
	assert.Equal(t, uint(0), v)
	assert.NotNil(t, err)
	v, err = calculateScoreForTwoRolls("11", "23")
	assert.Equal(t, uint(0), v)
	assert.NotNil(t, err)
	v, err = calculateScoreForTwoRolls("-1", "-23")
	assert.Equal(t, uint(0), v)
	assert.NotNil(t, err)
	v, err = calculateScoreForTwoRolls("-", "/")
	assert.Equal(t, uint(10), v)
	assert.Nil(t, err)
	v, err = calculateScoreForTwoRolls("/", "-")
	assert.Equal(t, uint(0), v)
	assert.NotNil(t, err)
}

func TestFrameCalculateScore(t *testing.T) {
	f := Frame{}
	f.calculateScore()
	assert.Equal(t, uint(0), f.score)
	f = Frame{rolls: []string{"4", "4"}}
	f.calculateScore()
	assert.Equal(t, uint(8), f.score)
	f = Frame{rolls: []string{"4", "/"}}
	f.calculateScore()
	assert.Equal(t, uint(10), f.score)
	f = Frame{rolls: []string{"x"}}
	f.calculateScore()
	assert.Equal(t, uint(10), f.score)
	f = Frame{rolls: []string{"x", "x"}}
	f.calculateScore()
	assert.Equal(t, uint(0), f.score)
	f = Frame{rolls: []string{"23", "3"}}
	f.calculateScore()
	assert.Equal(t, uint(0), f.score)
	f = Frame{rolls: []string{"3", "3", "4"}}
	f.calculateScore()
	assert.Equal(t, uint(0), f.score)
	f = Frame{rolls: []string{"-1", "3"}}
	f.calculateScore()
	assert.Equal(t, uint(0), f.score)
	f = Frame{rolls: []string{"/", "3"}}
	f.calculateScore()
	assert.Equal(t, uint(0), f.score)
	f = Frame{rolls: []string{"/", "/"}}
	f.calculateScore()
	assert.Equal(t, uint(0), f.score)
	f = Frame{rolls: []string{"/", "-"}}
	f.calculateScore()
	assert.Equal(t, uint(0), f.score)
	f = Frame{rolls: []string{"-", "-"}}
	f.calculateScore()
	assert.Equal(t, uint(0), f.score)
	f = Frame{rolls: []string{"-", "8"}}
	f.calculateScore()
	assert.Equal(t, uint(8), f.score)
	f = Frame{rolls: []string{"8", "-"}}
	f.calculateScore()
	assert.Equal(t, uint(8), f.score)
	f = Frame{rolls: []string{"abc", "-"}}
	f.calculateScore()
	assert.Equal(t, uint(0), f.score)
	f = Frame{rolls: []string{"abc", "/"}}
	f.calculateScore()
	assert.Equal(t, uint(0), f.score)
	f = Frame{rolls: []string{"4", "abc"}}
	f.calculateScore()
	assert.Equal(t, uint(0), f.score)
	f = Frame{rolls: []string{"abc", "4"}}
	f.calculateScore()
	assert.Equal(t, uint(0), f.score)
}

func TestFrameStackIsEmpty(t *testing.T) {
	var frameStack FrameStack
	assert.Equal(t, true, frameStack.IsEmpty())

	frameStack.Push(Frame{})
	assert.Equal(t, false, frameStack.IsEmpty())
}

func TestFrameStackPush(t *testing.T) {
	var frameStack FrameStack
	assert.Equal(t, 0, len(frameStack))

	frameStack.Push(Frame{})
	assert.Equal(t, 1, len(frameStack))
}

func TestFrameStackPop(t *testing.T) {
	var frameStack FrameStack
	frameStack.Push(Frame{})
	assert.Equal(t, 1, len(frameStack))

	frameStack.Push(Frame{})
	assert.Equal(t, 2, len(frameStack))

	_, ok := frameStack.Pop()
	assert.Equal(t, 1, len((frameStack)))
	assert.Equal(t, true, ok)

	_, ok = frameStack.Pop()
	assert.Equal(t, 0, len((frameStack)))
	assert.Equal(t, true, ok)

	_, ok = frameStack.Pop()
	assert.Equal(t, false, ok)
}
