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

func hasVisibleNbrInDirection(floorPlan []string, x int, y int, dirX int, dirY int) bool {
	newX, newY := x, y
	newX += dirX
	newY += dirY

	for newY >= 0 && newX >= 0 && newY < len(floorPlan) && newX < len(floorPlan[y]) {
		if string(floorPlan[newY][newX]) == "#" {
			return true
		} else if string(floorPlan[newY][newX]) == "L" {
			return false
		}
		newX += dirX
		newY += dirY
	}

	return false
}

func getNumberOfVisibleNeighbors(floorPlan []string, x int, y int) int {
	// debug := false
	// if x == 3 && y == 1 {
	// 	debug = true
	// }

	visNbrs := 0

	// top left
	if hasVisibleNbrInDirection(floorPlan, x, y, -1, -1) {
		// fmt.Println("top left seen")
		visNbrs++
	}

	// top middle
	if hasVisibleNbrInDirection(floorPlan, x, y, 0, -1) {
		// fmt.Println("top middle seen")
		visNbrs++
	}

	// top right
	if hasVisibleNbrInDirection(floorPlan, x, y, 1, -1) {
		// fmt.Println("top right seen")
		visNbrs++
	}

	// middle left
	if hasVisibleNbrInDirection(floorPlan, x, y, -1, 0) {
		// fmt.Println("middle left seen")
		visNbrs++
	}

	// middle right
	if hasVisibleNbrInDirection(floorPlan, x, y, 1, 0) {
		// fmt.Println("middle right seen")
		visNbrs++
	}

	// bottom left
	if hasVisibleNbrInDirection(floorPlan, x, y, -1, 1) {
		// fmt.Println("bottom left seen")
		visNbrs++
	}

	// bottom middle
	if hasVisibleNbrInDirection(floorPlan, x, y, 0, 1) {
		// fmt.Println("bottom middle seen")
		visNbrs++
	}

	// bottom right
	if hasVisibleNbrInDirection(floorPlan, x, y, 1, 1) {
		// fmt.Println("bottom right seen")
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
	// fmt.Printf("at iteration %d, state is:\n", iteration)
	oldPlan := duplicateSlice(floorPlan)
	for !areTwoPlansEqual(newPlan, oldPlan) {
		// copy(oldPlan, newPlan)
		oldPlan = duplicateSlice(newPlan)
		newPlan = getNextStateAll(newPlan)
		iteration++
		// fmt.Printf("at iteration %d, state is:\n", iteration)
		// printFloorPlan(newPlan)
	}
	// fmt.Printf("at iteration %d, state has stabilized:\n", iteration)
	// printFloorPlan(newPlan)
	return countAllOccupiedStates(newPlan)
}

func main() {
	fmt.Println("day11-01 started")
	floorPlan := loadInputIntoListOfStrings("input")
	// test8Seen := loadInputIntoListOfStrings("test_8_occupied_seen")
	// testNoneSeen := loadInputIntoListOfStrings("test_no_occupied_seen")
	// testEmptyBlocking := loadInputIntoListOfStrings("test_empty_blocking")
	// fmt.Printf("floorPlan is: \n")
	// printFloorPlan(floorPlan)
	// fmt.Printf("After one iteration, the floor plan is now:\n")
	// printFloorPlan(getNextStateAll(floorPlan))
	// fmt.Printf("At initial state, floorplan is: \n")
	// printFloorPlan(floorPlan)
	fmt.Printf("number of occupied seats after floorPlan states have stabilized is: %d\n", partTwo(floorPlan))
	// fmt.Printf("seat at position <1,1> would see 0 occupied seats: getNumberOfVisibleNeighbors: %d\n", getNumberOfVisibleNeighbors(testEmptyBlocking, 1, 1))
	// fmt.Printf("seat at position <3,4> would see 8 occupied seats: getNumberOfVisibleNeighbors: %d\n", getNumberOfVisibleNeighbors(test8Seen, 3, 4))
	// fmt.Printf("seat at position <3,3> would see 0 occupied seats: getNumberOfVisibleNeighbors: %d\n", getNumberOfVisibleNeighbors(testNoneSeen, 3, 3))
}
