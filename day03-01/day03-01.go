package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//var mapInput string = ""

func loadInput() []string {
	path := "input"
	input := []string{}

	buf, err := os.Open(path)
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

// func isTree(treeMap []string, x int, y int) bool {
// 	// returns true if the coordinate (x, y) is a tree on the map
// 	// first, get max x of map
// 	//fmt.Println("treeMap[0] is:", treeMap[y])
// 	length := len(string(treeMap[y]))
// 	//fmt.Println("length is:", length)

// 	// modulo the x coordinate to get the effective x on the raw map
// 	effectiveX := x % (length)
// 	//fmt.Println("effectiveX is:", effectiveX)

// 	//fmt.Printf("At x: %d, y: %d, treeMap[y][x] is is %v \n", x, y, string(treeMap[y][effectiveX]))

// 	if string(treeMap[y][effectiveX]) == "#" {
// 		return true
// 	}
// 	return false

// 


func isTree(treeMapLine string, x int, y int) bool {
	// returns true if the coordinate (x, y) is a tree on the map
	length := len(string(treeMapLine))

	// modulo the x coordinate to get the effective x on the raw map
	effectiveX := x % (length)

	if string(treeMapLine[effectiveX]) == "#" {
		return true
	}
	return false

}


func main() {
	fmt.Println("day03-01 started")
	treeMap := loadInput()

	height := len(treeMap)

	fmt.Printf("Height of this mountain is: %d\n", height)

	treesHit := 0

	for y, _ := range treeMap {
		// if isTree(treeMap, (3 * y), y) {
		if isTree(treeMap[y], (3 * y), y) {
			treesHit++
		}
	}

	fmt.Printf("treesHit: %d\n", treesHit)
}
