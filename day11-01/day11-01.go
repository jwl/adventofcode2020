package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func loadInputIntoListOfStrings(filename string) []string {
	input := []string{}

	buf, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(buf)
	for snl.Scan() {
		input = append(input, snl.Text())
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func getNumberOfAdjacentNeighbors(floorPlan []string, x int, y int) int {
	adjNbrs := 0

	// top left
	if y > 0 && x > 0 && string(floorPlan[y-1][x-1]) == "#" {
		adjNbrs++
	}
	// top middle
	if y > 0 && string(floorPlan[y-1][x]) == "#" {
		adjNbrs++
	}
	// top right
	if y > 0 && x < len(floorPlan[y])-1 && string(floorPlan[y-1][x+1]) == "#" {
		adjNbrs++
	}
	// left middle
	if x > 0 && string(floorPlan[y][x-1]) == "#" {
		adjNbrs++
	}
	// right middle
	if x < len(floorPlan[y])-1 && string(floorPlan[y][x+1]) == "#" {
		adjNbrs++
	}
	// bottom left
	if y < len(floorPlan)-1 && x > 0 && string(floorPlan[y+1][x-1]) == "#" {
		adjNbrs++
	}
	// bottom middle
	if y < len(floorPlan)-1 && string(floorPlan[y+1][x]) == "#" {
		adjNbrs++
	}
	// bottom right
	if y < len(floorPlan)-1 && x < len(floorPlan[y])-1 && string(floorPlan[y+1][x+1]) == "#" {
		adjNbrs++
	}

	return adjNbrs
}

func getNextState(floorPlan []string, x int, y int) string {
	// note that any given position is defined by floorPlan[y][x]
	if string(floorPlan[y][x]) == "L" && getNumberOfAdjacentNeighbors(floorPlan, x, y) == 0 {
		// If seat is empty, and there are no occupied seats adjacent to it,
		// it becomes occupied (#)
		return "#"

	} else if string(floorPlan[y][x]) == "#" && getNumberOfAdjacentNeighbors(floorPlan, x, y) >= 4 {
		// If seat is occupied and four or more seats adjacent to it are
		// also occupied, that seat becomes empty (L)
		return "L"
	}

	// else, the seat does not change
	return string(floorPlan[y][x])
}

func getNextStateAll(floorPlan []string) []string {
	newPlan := []string{}
	for y, row := range floorPlan {
		newRow := ""
		for x := range row {
			newC := getNextState(floorPlan, x, y)
			newRow += string(newC)
		}
		newPlan = append(newPlan, newRow)
	}
	return newPlan
}

func printFloorPlan(floorPlan []string) {
	for _, row := range floorPlan {
		fmt.Println(row)
	}
}

func countAllOccupiedStates(floorPlan []string) int {
	occupiedSeats := 0
	for _, row := range floorPlan {
		for _, c := range row {
			if string(c) == "#" {
				occupiedSeats++
			}
		}
	}
	return occupiedSeats
}

func duplicateSlice(src []string) []string {
	tmp := make([]string, len(src))
	copy(tmp, src)
	return tmp
}

func areTwoPlansEqual(p1 []string, p2 []string) bool {
	for y, row := range p1 {
		for x, c := range row {
			if string(c) != string(p2[y][x]) {
				return false
			}
		}
	}
	return true
}

func partOne(floorPlan []string) int {
	// iterates floorPlans until state stabilizes (ie. doesn't change)
	// then returns the number of occupied seats
	iteration := 1
	newPlan := getNextStateAll(floorPlan)
	oldPlan := duplicateSlice(floorPlan)
	for !areTwoPlansEqual(newPlan, oldPlan) {
		// copy(oldPlan, newPlan)
		oldPlan = duplicateSlice(newPlan)
		newPlan = getNextStateAll(newPlan)
		iteration++
	}
	fmt.Printf("at iteration %d, state has stabilized:\n", iteration)
	printFloorPlan(newPlan)
	return countAllOccupiedStates(newPlan)
}

func main() {
	fmt.Println("day11-01 started")
	floorPlan := loadInputIntoListOfStrings("input")
	// fmt.Printf("floorPlan is: \n")
	// printFloorPlan(floorPlan)
	// fmt.Printf("After one iteration, the floor plan is now:\n")
	// printFloorPlan(getNextStateAll(floorPlan))
	fmt.Printf("number of occupied seats after floorPlan states have stabilized is: %d\n", partOne(floorPlan))
}
