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

func duplicateSlice(src []int) []int {
	tmp := make([]int, len(src))
	copy(tmp, src)
	return tmp
}

func findNextHighestByN(list []int, x int, n int) int {
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

func getJoltDifferences(list []int) {
	// number of 1, 2 and 3 jolt differences
	// the final jump is always 3 jolts so d3 starts at 1
	d1, d2, d3 := 0, 0, 1
	largest := getLargestElement(list)

	x := 0
	for true {
		if x >= largest {
			break
		}
		if findNextHighestByN(list, x, 1) != -1 {
			// finalChain = append(finalChain, list[i])
			d1++
			x = x + 1
		} else if findNextHighestByN(list, x, 2) != -1 {
			d2++
			x = x + 2
		} else if findNextHighestByN(list, x, 3) != -1 {
			d3++
			x = x + 3
		} else {
			fmt.Printf("Couldn't find next highest from x: %d\n", x)
		}

	}
	fmt.Printf("Reached end, d1 is: %d, d3 is %d\n", d1, d3)
}

func main() {
	fmt.Println("day10-01 started")
	input := loadInputIntoListOfStrings("input")
	getJoltDifferences(input)
}
