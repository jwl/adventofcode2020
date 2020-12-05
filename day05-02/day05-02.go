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

func convertRowToDecimal(rowString string) int {
	binaryString := strings.ReplaceAll(rowString, "F", "0")
	binaryString = strings.ReplaceAll(binaryString, "B", "1")
	result, _ := strconv.ParseInt(binaryString, 2, 64)

	return int(result)
}

func convertColumnToDecimal(colString string) int {
	binaryString := strings.ReplaceAll(colString, "L", "0")
	binaryString = strings.ReplaceAll(binaryString, "R", "1")
	result, _ := strconv.ParseInt(binaryString, 2, 64)

	return int(result)
}

func calcSeatID(row int, col int) int {
	return row*8 + col
}

func getHighestSeatID(boardingPassList []string) int {
	// iterate through each boardingpass, calculate its seat ID
	highestSeatID := 0
	seatID := -1
	for _, rawBoardingPass := range boardingPassList {
		seatID = calcSeatID(convertRowToDecimal(rawBoardingPass[:7]), convertColumnToDecimal(rawBoardingPass[7:]))
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	return highestSeatID
}

func main() {
	// f = 0, b = 1
	fmt.Println("day05-02 started")
	boardingPasses := loadInput("input")

	// fmt.Printf("boardingPasses: %#v\n", boardingPasses)

	highestSeatID := getHighestSeatID(boardingPasses)

	fmt.Printf("highest seat ID is: %d\n", highestSeatID)

	// fmt.Printf("test FBFBBFF should be 44: %d\n", convertToDecimal("FBFBBFF"))

}
