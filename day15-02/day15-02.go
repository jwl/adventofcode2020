package main

import (
	"fmt"
)

func partTwo(startingNumbers []int, finalTurn int) int {
	fmt.Printf("Starting numbers are: %#v\n", startingNumbers)
	// note that the 2020th number spoken will be at turn 2019
	turn := 0

	// gameHistory[number][array of ints representing turns number was spoken]
	lookup := make(map[int][]int) // append turns
	history := []int{}            // append spokenNumbers

	// for turn < 2020 {
	for turn < finalTurn {
		if turn < len(startingNumbers) {
			spokenNumber := startingNumbers[turn]
			lookup[spokenNumber] = append(lookup[spokenNumber], turn)
			history = append(history, spokenNumber)
		} else {
			// consider the last number spoken
			previousNum := history[turn-1]

			numHistory := lookup[previousNum]
			if len(numHistory) >= 2 {
				// If it had been spoken before, get the difference between the turn from when
				// it was last spoken, and the turn when it was spoken before that
				// if it had been spoken before, numHistory must be more than 2
				lastTurn := numHistory[len(lookup[previousNum])-1]
				beforeLastTurn := numHistory[len(lookup[previousNum])-2]
				spokenNumber := lastTurn - beforeLastTurn
				lookup[spokenNumber] = append(lookup[spokenNumber], turn)
				history = append(history, spokenNumber)

			} else {
				// If this is a new number, spokenNumber is 0
				spokenNumber := 0
				lookup[spokenNumber] = append(lookup[spokenNumber], turn)
				history = append(history, spokenNumber)

			}
		}
		turn++
	}

	return history[finalTurn-1]
}

func main() {
	fmt.Println("day15-01 started")

	// // sample1, answer is 436
	// sample1 := []int{0, 3, 6}
	// fmt.Printf("the 2020th number spoken for %#v should be 175594, is actually: %d\n", sample1, partTwo(sample1, 30000000))

	// // sample2, answer is 1
	// sample2 := []int{1, 3, 2}
	// fmt.Printf("the 2020th number spoken for %#v should be 2578, is actually: %d\n", sample2, partTwo(sample2, 30000000))

	// // sample3, answer is 10
	// sample3 := []int{2, 1, 3}
	// fmt.Printf("the 2020th number spoken for %#v should be 3544142, is actually: %d\n", sample3, partTwo(sample3, 30000000))

	// // sample4, answer is 27
	// sample4 := []int{1, 2, 3}
	// fmt.Printf("the 2020th number spoken for %#v should be 261214, is actually: %d\n", sample4, partTwo(sample4, 30000000))

	// // sample5, answer is 78
	// sample5 := []int{2, 3, 1}
	// fmt.Printf("the 2020th number spoken for %#v should be 6895259, is actually: %d\n", sample5, partTwo(sample5, 30000000))

	// // sample6, answer is 438
	// sample6 := []int{3, 2, 1}
	// fmt.Printf("the 2020th number spoken for %#v should be 18, is actually: %d\n", sample6, partTwo(sample6, 30000000))

	// // sample7, answer is 1838
	// sample7 := []int{3, 1, 2}
	// fmt.Printf("the 2020th number spoken for %#v should be 362, is actually: %d\n", sample7, partTwo(sample7, 30000000))

	// actual input
	input := []int{6, 3, 15, 13, 1, 0}
	fmt.Printf("the 2020th number spoken for %#v will be: %d\n", input, partTwo(input, 30000000))
}
