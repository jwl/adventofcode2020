package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var lookupTable map[int]int

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

func getLargestElement(list []int) int {
	largest := list[0]
	for i := 0; i < len(list); i++ {
		if list[i] > largest {
			largest = list[i]
		}
	}
	return largest
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isPathBetween(x int, y int, list []int) bool {
	if contains(list, x) && contains(list, y) {
		return true
	}
	return false
}

func getNumberOfPaths(src int, target int, list []int) int {
	if (target - src) <= 3 {
		return 1
	}
	paths := 0
	if isPathBetween(src, src+1, list) {
		if lookupTable[src+1] == 0 {
			lookupTable[src+1] = getNumberOfPaths(src+1, target, list)
		}
		paths += lookupTable[src+1]
	}
	if isPathBetween(src, src+2, list) {
		if lookupTable[src+2] == 0 {
			lookupTable[src+2] = getNumberOfPaths(src+2, target, list)
		}
		paths += lookupTable[src+2]
	}
	if isPathBetween(src, src+3, list) {
		if lookupTable[src+3] == 0 {
			lookupTable[src+3] = getNumberOfPaths(src+3, target, list)
		}
		paths += lookupTable[src+3]
	}
	return paths
}

func main() {
	fmt.Println("day10-02 started")
	input := loadInputIntoListOfStrings("input")
	max := getLargestElement(input)
	input = append(input, 0)
	sort.Ints(input)
	lookupTable = make(map[int]int)
	arrangements := getNumberOfPaths(0, max+3, input)
	fmt.Printf("Number of arrangements is: %d\n", arrangements)

}
