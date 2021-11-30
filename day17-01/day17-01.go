package main

import (
	"fmt"

	"github.com/jwl/adventofcode2020/aocutils"
)

maxArr := 1000

func getSumOfActive(state [maxArr][maxArr][maxArr]rune) int {
	max := len(state[0])
	sum := 0
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			for k := 0; k < max; k++ {
				if state[i][j][k] == "#" {
					sum++
				}
			}
		}
	}

	return sum
}

func getNextState(state [maxArr][maxArr][maxArr]rune) [maxArr][maxArr][maxArr]rune {
	nextState := state

	return nextState
}

func convertRawInputToState([]string) [maxArr][maxArr][maxArr]rune {

}

func partOne(input []string) int {
	numberOfCycles := 6

	// state := [maxArr][maxArr][maxArr]rune{}

	state := convertRawInputToState(input)

	for i := 0; i <= numberOfCycles; i++ {
		state = getNextState(state)
	}

	sum := getSumOfActive(state)
	return sum
}

func main() {
	fmt.Println("dayXX-YY started")

	input := aocutils.LoadInputIntoListOfStrings("sample1")
	// input := aocutils.LoadInputIntoListOfStrings("input")

	fmt.Printf("Number of cubes active after 6 cycles: %d\n", partOne(input))

}
