package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jwl/adventofcode2020/aocutils"
)

func writeToMemory(memory map[int]int, address int, value int, bitmask string) {
	valueAsBinaryStr := convertIntToBinaryString(value)
	valueAsInt := applyBitMask(valueAsBinaryStr, bitmask)

	memory[address] = valueAsInt
}

func convertIntToBinaryString(value int) string {
	binaryStr := strconv.FormatInt(int64(value), 2)
	numberOfPaddedZeroes := 36 - len(binaryStr)

	for i := 0; i < numberOfPaddedZeroes; i++ {
		binaryStr = "0" + binaryStr
	}

	return binaryStr
}

func applyBitMask(value string, bitmask string) int {
	resultAsBinaryStr := ""
	for i, j := range bitmask {
		switch string(j) {
		case "X":
			resultAsBinaryStr = resultAsBinaryStr + string(value[i])
		case "0":
			resultAsBinaryStr = resultAsBinaryStr + "0"
		case "1":
			resultAsBinaryStr = resultAsBinaryStr + "1"
		}

	}

	resultAsInt64, _ := strconv.ParseInt(resultAsBinaryStr, 2, 64)

	return int(resultAsInt64)
}

func sumAllValuesInMemory(memory map[int]int) int {
	sum := 0
	for _, v := range memory {
		sum += v
	}
	return sum
}

func partOne(input []string) {
	memory := make(map[int]int)
	bitmask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" // start with a bitmask that allows everything

	for _, line := range input {
		if line[0:4] == "mask" {
			bitmask = line[7:]
			continue
		}
		if line[0:3] == "mem" {
			re := regexp.MustCompile(`mem\[(\d+)\] \= (\d+)`)
			match := re.FindStringSubmatch(line)
			address, _ := strconv.Atoi(match[1])
			value, _ := strconv.Atoi(match[2])
			writeToMemory(memory, address, value, bitmask)
		}
	}

	sum := sumAllValuesInMemory(memory)

	fmt.Printf("memory is now: %#v\n", memory)
	fmt.Printf("Sum of all values in memory is: %d\n", sum)

}

func main() {
	fmt.Println("day14-01 started")

	// main
	input := aocutils.LoadInputIntoListOfStrings("input")
	partOne(input)
}
