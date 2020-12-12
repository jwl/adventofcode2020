package main

import (
	"fmt"
	"strconv"

	"github.com/jwl/adventofcode2020/aocutils"
)

type state struct {
	sx int // positive values mean east
	sy int // positive values mean north
	// dir int
	// dir = 0 -> east
	// dir = 90 -> north
	// dir = 180 -> west
	// dir = 270 -> south
	wx int // waypoint coordinates x
	wy int // waypoint coordinates y
}

func getWPRelativePos(s state) (int, int) {
	// returns x and y relative distances of waypoint from ship
	// positive x means waypoint is east
	// positive y means waypoint is north
	return s.wx - s.sx, s.wy - s.sy
}

func getNextState(s state, instruction string) state {
	action := string(instruction[0])
	value, _ := strconv.Atoi(instruction[1:])
	switch action {
	case "N":
		s.wy += value
	case "S":
		s.wy -= value
	case "E":
		s.wx += value
	case "W":
		s.wx -= value
	case "L":
		switch value {
		case 90:
			// identical to R270
			oldRelativeX, oldRelativeY := getWPRelativePos(s)
			newRelativeX := -(oldRelativeY)
			newRelativeY := oldRelativeX
			s.wx = s.sx + newRelativeX
			s.wy = s.sy + newRelativeY
		case 180:
			oldRelativeX, oldRelativeY := getWPRelativePos(s)
			newRelativeX := -(oldRelativeX)
			newRelativeY := -(oldRelativeY)
			s.wx = s.sx + newRelativeX
			s.wy = s.sy + newRelativeY
		case 270:
			// identical to R90
			oldRelativeX, oldRelativeY := getWPRelativePos(s)
			newRelativeX := oldRelativeY
			newRelativeY := -(oldRelativeX)
			s.wx = s.sx + newRelativeX
			s.wy = s.sy + newRelativeY
		}
	case "R":
		switch value {
		case 90:
			// identical to L270
			oldRelativeX, oldRelativeY := getWPRelativePos(s)
			newRelativeX := oldRelativeY
			newRelativeY := -(oldRelativeX)
			s.wx = s.sx + newRelativeX
			s.wy = s.sy + newRelativeY
		case 180:
			oldRelativeX, oldRelativeY := getWPRelativePos(s)
			newRelativeX := -(oldRelativeX)
			newRelativeY := -(oldRelativeY)
			s.wx = s.sx + newRelativeX
			s.wy = s.sy + newRelativeY
		case 270:
			// identical to L90
			oldRelativeX, oldRelativeY := getWPRelativePos(s)
			newRelativeX := -(oldRelativeY)
			newRelativeY := oldRelativeX
			s.wx = s.sx + newRelativeX
			s.wy = s.sy + newRelativeY
		}
	case "F":
		relX, relY := getWPRelativePos(s)
		s.sx += relX * value
		s.sy += relY * value
		s.wx += relX * value
		s.wy += relY * value
	default:
		fmt.Printf("Invalid action: %#v, quitting!\n", action)
		return s
	}

	return s
}

func partTwo(input []string) int {
	// returns manhattan distance from origin
	s := state{0, 0, 10, 1}

	for _, instruction := range input {
		s = getNextState(s, instruction)
	}

	fmt.Printf("Final state is: %#v\n", s)
	return (aocutils.Abs(s.sx) + aocutils.Abs(s.sy))
}

func main() {
	fmt.Println("day12-01 started")

	input := aocutils.LoadInputIntoListOfStrings("input")

	fmt.Printf("Manhattan distance is: %d\n", partTwo(input))
}
