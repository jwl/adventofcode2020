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

func duplicateSlice(src []int) []int {
	tmp := make([]int, len(src))
	copy(tmp, src)
	return tmp
}

func findNextHighestByN(x int, n int, list []int) int {
	// returns index of element that is exactly n higher than x if it exists
	// if such an element doesn't exist, return -1
	for i, element := range list {
		if element == x+n {
			return i
		}
	}

	return -1
}

func remove(s []int, i int) []int {
	// return a slice with the element at index i removed
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
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
			fmt.Printf("lookupTable[%d] is: %d\n", src+1, lookupTable[src+1])
		}
		paths += lookupTable[src+1]
	}
	if isPathBetween(src, src+2, list) {
		if lookupTable[src+2] == 0 {
			lookupTable[src+2] = getNumberOfPaths(src+2, target, list)
			fmt.Printf("lookupTable[%d] is: %d\n", src+2, lookupTable[src+2])
		}
		paths += lookupTable[src+2]
	}
	if isPathBetween(src, src+3, list) {
		if lookupTable[src+3] == 0 {
			lookupTable[src+3] = getNumberOfPaths(src+3, target, list)
			fmt.Printf("lookupTable[%d] is: %d\n", src+3, lookupTable[src+3])
		}
		paths += lookupTable[src+3]
	}
	fmt.Printf("From src %d to target %d, number of paths is %d\n", src, target, paths)
	return paths
}

func main() {
	fmt.Println("day10-02 started")
	input := loadInputIntoListOfStrings("input")
	max := getLargestElement(input)
	// input = append(input, max)
	input = append(input, 0)
	sort.Ints(input)
	lookupTable = make(map[int]int)
	fmt.Printf("input is now: %#v\n", input)
	arrangements := getNumberOfPaths(0, max+3, input)
	fmt.Printf("Number of arrangements is: %d\n", arrangements)

}