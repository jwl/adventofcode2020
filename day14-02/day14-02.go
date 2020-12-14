package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/jwl/adventofcode2020/aocutils"
)

func writeToMemory(memory map[int]int, address int, value int, bitmask string) {
	addrAsBinaryStr := applyBitmask(convertIntToBinaryString(address, 36), bitmask)
	re := regexp.MustCompile("X")
	xMatches := re.FindAllStringIndex(addrAsBinaryStr, -1)

	if len(xMatches) > 0 {
		// populate addresses with all possible addresses that result from addrAsBinaryStr
		addresses := []int{}
		// for every X, the number of addresses is 2^X
		numberOfAddresses := int(math.Pow(2, float64(len(xMatches))))

		for i := 0; i < numberOfAddresses; i++ {
			binaryAddrFragment := []rune(convertIntToBinaryString(i, len(xMatches)))
			tmpAddress := addrAsBinaryStr

			// replace all X's with 1's and 0's depending on what x is
			for a, xIndex := range xMatches {
				tmpAddress = aocutils.ReplaceAtIndex(tmpAddress, binaryAddrFragment[a], xIndex[0])
			}

			tmpAddress64, _ := strconv.ParseInt(tmpAddress, 2, 64)
			addresses = append(addresses, int(tmpAddress64))
		}

		for _, address := range addresses {
			memory[address] = value
		}

	} else {
		modifiedAddress, _ := strconv.ParseInt(addrAsBinaryStr, 2, 64)
		memory[int(modifiedAddress)] = value
	}
}

func convertIntToBinaryString(value int, length int) string {
	binaryStr := strconv.FormatInt(int64(value), 2)
	numberOfPaddedZeroes := length - len(binaryStr)

	for i := 0; i < numberOfPaddedZeroes; i++ {
		binaryStr = "0" + binaryStr
	}

	return binaryStr
}

func applyBitmask(value string, bitmask string) string {
	resultAsBinaryStr := ""
	for i, j := range bitmask {
		switch string(j) {
		case "X":
			resultAsBinaryStr = resultAsBinaryStr + "X"
		case "0":
			resultAsBinaryStr = resultAsBinaryStr + string(value[i])
		case "1":
			resultAsBinaryStr = resultAsBinaryStr + "1"
		}

	}

	return resultAsBinaryStr
}

func sumAllValuesInMemory(memory map[int]int) int {
	sum := 0
	for _, v := range memory {
		sum += v
	}
	return sum
}

func partTwo(input []string) {
	memory := make(map[int]int)
	bitmask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" // start with a bitmask that allows everything

	for _, line := range input {
		fmt.Printf("Processing line: %s\n", line)
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
	fmt.Println("day14-02 started")

	// main
	input := aocutils.LoadInputIntoListOfStrings("input")
	partTwo(input)
}
