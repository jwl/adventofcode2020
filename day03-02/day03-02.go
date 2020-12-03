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

func getTreesForSlope(treeMap []string, slopeX int, slopeY int) int {
	 // for a given slope, return the number of trees hit
	 treesHit := 0
	 height := len(treeMap)

	 for i := 0; i * slopeY < height; i++ {
	 	if isTree(treeMap[i * slopeY], i * slopeX, i * slopeY) {
	 		treesHit++
	 	}
	 }

	 fmt.Printf("For slope (%d, %d), treesHit is: %d\n", slopeX, slopeY, treesHit)
	 return treesHit

}

func main() {
	fmt.Println("day03-01 started")
	treeMap := loadInput()

	slope1TreesHit := getTreesForSlope(treeMap, 1, 1)
	slope2TreesHit := getTreesForSlope(treeMap, 3, 1)
	slope3TreesHit := getTreesForSlope(treeMap, 5, 1)
	slope4TreesHit := getTreesForSlope(treeMap, 7, 1)
	slope5TreesHit := getTreesForSlope(treeMap, 1, 2)

	fmt.Printf("Product of all trees hit on all slopes is %d\n", slope1TreesHit * slope2TreesHit * slope3TreesHit * slope4TreesHit * slope5TreesHit)
}
