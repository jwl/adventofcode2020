package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

func isTree(treeMapLine string, x int) bool {
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

	for y := range treeMap {
		if isTree(treeMap[y], (3 * y)) {
			treesHit++
		}
	}

	fmt.Printf("treesHit: %d\n", treesHit)
}
