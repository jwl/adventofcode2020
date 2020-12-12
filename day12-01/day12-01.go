package main

import (
	"fmt"
	"strconv"

	"github.com/jwl/adventofcode2020/aocutils"
)

type state struct {
	x   int // positive values mean east
	y   int // positive values mean north
	dir int
	// dir = 0 -> east
	// dir = 90 -> north
	// dir = 180 -> west
	// dir = 270 -> south
}

func getNextState(s state, instruction string) state {
	action := string(instruction[0])
	value, _ := strconv.Atoi(instruction[1:])
	switch action {
	case "N":
		s.y += value
	case "S":
		s.y -= value
	case "E":
		s.x += value
	case "W":
		s.x -= value
	case "L":
		s.dir = (s.dir + value) % 360
		if s.dir < 0 {
			s.dir += 360
		}
	case "R":
		s.dir = (s.dir - value) % 360
		if s.dir < 0 {
			s.dir += 360
		}
	case "F":
		switch s.dir {
		case 0:
			s.x += value
		case 90:
			s.y += value
		case 180:
			s.x -= value
		case 270:
			s.y -= value
		}
	default:
		fmt.Printf("Invalid action: %#v, quitting!\n", action)
		return s
	}

	return s
}

func partOne(input []string) int {
	// returns manhattan distance from origin
	s := state{0, 0, 0}

	for _, instruction := range input {
		s = getNextState(s, instruction)
	}

	fmt.Printf("Final state is: %#v\n", s)
	return (aocutils.Abs(s.x) + aocutils.Abs(s.y))
}

func main() {
	fmt.Println("day12-01 started")

	// input := []int{1, 2, 3}
	// fmt.Printf("input contains 3: %#v", aocutils.Contains(input, 3))

	input := aocutils.LoadInputIntoListOfStrings("input")

	fmt.Printf("Manhattan distance is: %d\n", partOne(input))

	// testState := state{0, 0, 0}
	// fmt.Printf("getNextState is: %#v\n", getNextState(testState, "N10"))
}
