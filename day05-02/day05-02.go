package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func loadInput(filename string) []string {
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

func getHighestSeatID(boardingPassList []string) int {
	highestSeatID := 0
	seatID := -1
	for _, rawBoardingPass := range boardingPassList {
		seatID = convertBoardingPassToDecimal(rawBoardingPass)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	return highestSeatID
}

func getLowestSeatID(boardingPassList []string) int {
	lowestSeatID := 9999999
	seatID := -1
	for _, rawBoardingPass := range boardingPassList {
		seatID = convertBoardingPassToDecimal(rawBoardingPass)
		if seatID < lowestSeatID {
			lowestSeatID = seatID
		}
	}
	return lowestSeatID
}

func getMissingSeatID(boardingPassList []string) int {
	highestSeatID := getHighestSeatID(boardingPassList)
	lowestSeatID := getLowestSeatID(boardingPassList)

	theoreticalSum := getSum(highestSeatID) - getSum(lowestSeatID-1)
	actualSum := 0

	for _, boardingPass := range boardingPassList {
		fmt.Printf("%d\n", convertBoardingPassToDecimal(boardingPass))
		actualSum += convertBoardingPassToDecimal(boardingPass)
	}

	return theoreticalSum - actualSum
}

func getSum(x int) int {
	var sum int
	for i := 0; i <= x; i++ {
		sum += i
	}
	return sum
}

func convertBoardingPassToDecimal(boardingPass string) int {
	// Every boarding pass is just a binary number with F's and L's as 0's and B's and R's as 1's.
	binaryString := strings.ReplaceAll(boardingPass, "F", "0")
	binaryString = strings.ReplaceAll(binaryString, "B", "1")
	binaryString = strings.ReplaceAll(binaryString, "L", "0")
	binaryString = strings.ReplaceAll(binaryString, "R", "1")

	result, _ := strconv.ParseInt(binaryString, 2, 64)

	return int(result)
}

func main() {
	fmt.Println("day05-02 started")
	boardingPasses := loadInput("input")
	fmt.Printf("Missing seatID is: %d\n", getMissingSeatID(boardingPasses))

}
