package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jwl/adventofcode2020/aocutils"
)

func getIDs(rawIDs string) []int {
	stringIDs := strings.Split(rawIDs, ",")
	fmt.Printf("stringIDs are: %#v\n", stringIDs)
	intIDs := []int{}

	for _, e := range stringIDs {
		if e != "x" {
			id, _ := strconv.Atoi(e)
			intIDs = append(intIDs, id)
		}
	}

	return intIDs
}

func getNextBus(earliestTimestamp int, busID int) int {
	return ((earliestTimestamp / busID) + 1) * busID
}

func partOne(input []string) int {
	earliestTimestamp, _ := strconv.Atoi(input[0])
	ids := getIDs(input[1])

	fmt.Println(earliestTimestamp)
	fmt.Println(ids)

	earliestBus := getNextBus(earliestTimestamp, ids[0])
	earliestID := ids[0]
	for _, id := range ids {
		// fmt.Printf("for id %d and earliestTimestamp %d, next availble bus is at %d\n", id, earliestTimestamp, getNextBus(earliestTimestamp, id))
		newEarliest := getNextBus(earliestTimestamp, id)
		if newEarliest < earliestBus {
			earliestBus = newEarliest
			earliestID = id
		}
	}
	fmt.Printf("Earliest bus is ID %d, it departs at %d and you need to wait %d minutes.\n", earliestID, earliestBus, earliestBus-earliestTimestamp)
	fmt.Printf("Multiplying the ID and the number of minutes you get: %d\n", earliestID*(earliestBus-earliestTimestamp))
	return earliestBus - earliestTimestamp
}

func main() {
	fmt.Println("day13-01 started")

	input := aocutils.LoadInputIntoListOfStrings("input")

	fmt.Printf("ID of earliest bus multiplied by number of minutes waited is: %d\n", partOne(input))
}
