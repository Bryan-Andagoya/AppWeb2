package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestLinkedListInsert(t *testing.T) {
	var linkedList LinkedList
	linkedList.Insert(Frame{rolls: []string{"4", "-"}})

	assert.Equal(t, uint(1), linkedList.length)
	assert.Equal(t, "4", linkedList.head.key.rolls[0])
	assert.Equal(t, "-", linkedList.head.key.rolls[1])
	assert.Equal(t, "4", linkedList.tail.key.rolls[0])
	assert.Equal(t, "-", linkedList.tail.key.rolls[1])
	assert.Nil(t, linkedList.head.prev)
	assert.Nil(t, linkedList.head.next)
	assert.Nil(t, linkedList.tail.prev)
	assert.Nil(t, linkedList.tail.next)

	linkedList.Insert(Frame{rolls: []string{"X"}})
	assert.Equal(t, uint(2), linkedList.length)
	assert.Equal(t, "4", linkedList.head.key.rolls[0])
	assert.Equal(t, "-", linkedList.head.key.rolls[1])
	assert.Equal(t, "X", linkedList.tail.key.rolls[0])
	assert.Nil(t, linkedList.head.prev)
	assert.NotNil(t, linkedList.head.next)
	assert.Equal(t, "X", linkedList.head.next.key.rolls[0])
	assert.Nil(t, linkedList.tail.next)
	assert.NotNil(t, linkedList.tail.prev)
	assert.Equal(t, "4", linkedList.tail.prev.key.rolls[0])
	assert.Equal(t, "-", linkedList.tail.prev.key.rolls[1])

	linkedList.Insert(Frame{rolls: []string{"5", "/"}})
	assert.Equal(t, uint(3), linkedList.length)
	assert.Equal(t, "4", linkedList.head.key.rolls[0])
	assert.Equal(t, "-", linkedList.head.key.rolls[1])
	assert.Equal(t, "5", linkedList.tail.key.rolls[0])
	assert.Equal(t, "/", linkedList.tail.key.rolls[1])
	assert.Nil(t, linkedList.head.prev)
	assert.NotNil(t, linkedList.head.next)
	assert.Equal(t, "X", linkedList.head.next.key.rolls[0])
	assert.Equal(t, "5", linkedList.head.next.next.key.rolls[0])
	assert.Equal(t, "/", linkedList.head.next.next.key.rolls[1])
	assert.Nil(t, linkedList.tail.next)
	assert.NotNil(t, linkedList.tail.prev)
	assert.Equal(t, "X", linkedList.tail.prev.key.rolls[0])

}
