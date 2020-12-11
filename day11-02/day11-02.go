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

func hasVisibleNbrInDirection(floorPlan []string, x int, y int, dx int, dy int) bool {
	newX, newY := x, y
	newX += dx
	newY += dy

	for newY >= 0 && newX >= 0 && newY < len(floorPlan) && newX < len(floorPlan[y]) {
		if string(floorPlan[newY][newX]) == "#" {
			return true
		} else if string(floorPlan[newY][newX]) == "L" {
			return false
		}
		newX += dx
		newY += dy
	}

	return false
}

func getNumberOfVisibleNeighbors(floorPlan []string, x int, y int) int {
	visNbrs := 0

	// top left
	if hasVisibleNbrInDirection(floorPlan, x, y, -1, -1) {
		visNbrs++
	}
	// top middle
	if hasVisibleNbrInDirection(floorPlan, x, y, 0, -1) {
		visNbrs++
	}
	// top right
	if hasVisibleNbrInDirection(floorPlan, x, y, 1, -1) {
		visNbrs++
	}
	// middle left
	if hasVisibleNbrInDirection(floorPlan, x, y, -1, 0) {
		visNbrs++
	}
	// middle right
	if hasVisibleNbrInDirection(floorPlan, x, y, 1, 0) {
		visNbrs++
	}
	// bottom left
	if hasVisibleNbrInDirection(floorPlan, x, y, -1, 1) {
		visNbrs++
	}
	// bottom middle
	if hasVisibleNbrInDirection(floorPlan, x, y, 0, 1) {
		visNbrs++
	}
	// bottom right
	if hasVisibleNbrInDirection(floorPlan, x, y, 1, 1) {
		visNbrs++
	}
	return visNbrs
}

func getNextState(floorPlan []string, x int, y int) string {
	// note that any given position is defined by floorPlan[y][x]
	if string(floorPlan[y][x]) == "L" && getNumberOfVisibleNeighbors(floorPlan, x, y) == 0 {
		// If seat is empty, and there are no occupied seats adjacent to it,
		// it becomes occupied (#)
		return "#"

	} else if string(floorPlan[y][x]) == "#" && getNumberOfVisibleNeighbors(floorPlan, x, y) >= 5 {
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

func partTwo(floorPlan []string) int {
	// iterates floorPlans until state stabilizes (ie. doesn't change)
	// then returns the number of occupied seats
	iteration := 1
	newPlan := getNextStateAll(floorPlan)
	oldPlan := duplicateSlice(floorPlan)
	for !areTwoPlansEqual(newPlan, oldPlan) {
		oldPlan = duplicateSlice(newPlan)
		newPlan = getNextStateAll(newPlan)
		iteration++
	}
	return countAllOccupiedStates(newPlan)
}

func main() {
	fmt.Println("day11-01 started")
	floorPlan := loadInputIntoListOfStrings("input")
	fmt.Printf("number of occupied seats after floorPlan states have stabilized is: %d\n", partTwo(floorPlan))
}
