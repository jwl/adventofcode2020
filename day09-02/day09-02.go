package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func loadInputIntoListOfStrings(filename string) []int {
	input := []int{}

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
		i, _ := strconv.Atoi(snl.Text())
		input = append(input, i)
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func remove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func duplicateSlice(src []int) []int {
	tmp := make([]int, len(src))
	copy(tmp, src)
	return tmp
}

func isValidSet(list []int, frameSize int) bool {
	target := list[len(list)-1]
	set := list[:len(list)-1]

	for i := range set {
		tmp := duplicateSlice(set)
		tmp = remove(tmp, i)
		for _, number := range tmp {
			if set[i]+number == target {
				// this is a valid set
				return true
			}
		}
	}

	return false
}

func getSumOfSlice(slice []int) int {
	i := 0
	for _, number := range slice {
		i += number
	}
	return i
}

func getSumOfSmallestLargest(slice []int) int {
	smallestNumber := slice[0]
	largestNumber := slice[0]
	for _, element := range slice {
		if element < smallestNumber {
			smallestNumber = element
		}
	}
	for _, element := range slice {
		if element > largestNumber {
			largestNumber = element
		}
	}

	return smallestNumber + largestNumber
}

func findEncryptionWeakness(invalidNumber int, input []int) int {
	fmt.Printf("Finding encryption weakness with invalidNumber: %d\n", invalidNumber)
	for i := range input {
		j := i + 1
		for j < len(input) {
			if getSumOfSlice(input[i:j]) == invalidNumber {
				return getSumOfSmallestLargest(input[i:j])
			}
			j++
		}
	}
	// couldn't find a solution...
	return -1
}

func main() {
	fmt.Println("day09-01 started")
	input := loadInputIntoListOfStrings("input")

	frameSize := 25
	maxIndex := len(input) - 1
	foundInvalidSet := false
	i := frameSize

	for true {
		if i > maxIndex {
			break
		}

		if !isValidSet(input[i-frameSize:i+1], frameSize) {
			foundInvalidSet = true
			break
		}

		i++
	}

	if foundInvalidSet {
		fmt.Printf("Found invalid set at index %d, which has a value of %d\n", i, input[i])
		fmt.Printf("Encryption weakness is: %d\n", findEncryptionWeakness(input[i], input))
	} else {
		fmt.Printf("Couldn't find an invalid set...\n")
	}

}
