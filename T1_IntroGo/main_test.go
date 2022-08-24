package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBowling(t *testing.T) {
	v, err := calculateFrameScore("4", "4")
	assert.Equal(t, 8, v)
	assert.Nil(t, err)
	v, err = calculateFrameScore("4", "/")
	assert.Equal(t, 10, v)
	assert.Nil(t, err)
	v, err = calculateFrameScore("x")
	assert.Equal(t, 10, v)
	assert.Nil(t, err)
	v, err = calculateFrameScore("4", "3", "3")
	assert.Equal(t, 0, v)
	assert.NotNil(t, err)
	v, err = calculateFrameScore("4", "abc")
	assert.Equal(t, 0, v)
	assert.NotNil(t, err)
	v, err = calculateFrameScore("abc", "5")
	assert.Equal(t, 0, v)
	assert.NotNil(t, err)
	v, err = calculateFrameScore()
	assert.Equal(t, 0, v)
	assert.NotNil(t, err)
}

func TestParseStringToInt(t *testing.T) {
	v, err := parseStringToInt("4")
	assert.Equal(t, 4, v)
	assert.Nil(t, err)
	v, err = parseStringToInt("abc")
	assert.Equal(t, 0, v)
	assert.NotNil(t, err)
	v, err = parseStringToInt("x")
	assert.Equal(t, 10, v)
	assert.Nil(t, err)
	v, err = parseStringToInt("X")
	assert.Equal(t, 10, v)
	assert.Nil(t, err)
}

func TestCalculateScoreForTwoRolls(t *testing.T) {
	v, err := calculateScoreForTwoRolls("4", "4")
	assert.Equal(t, 8, v)
	assert.Nil(t, err)
	v, err = calculateScoreForTwoRolls("abc", "4")
	assert.Equal(t, 0, v)
	assert.NotNil(t, err)
	v, err = calculateScoreForTwoRolls("4", "abc")
	assert.Equal(t, 0, v)
	assert.NotNil(t, err)
	v, err = calculateScoreForTwoRolls("4", "/")
	assert.Equal(t, 10, v)
	assert.Nil(t, err)
	v, err = calculateScoreForTwoRolls("/", "4")
	assert.Equal(t, 0, v)
	assert.NotNil(t, err)
	v, err = calculateScoreForTwoRolls("/", "/")
	assert.Equal(t, 0, v)
	assert.NotNil(t, err)
	v, err = calculateScoreForTwoRolls("abc", "/")
	assert.Equal(t, 0, v)
	assert.NotNil(t, err)
}
