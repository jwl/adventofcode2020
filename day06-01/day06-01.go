package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func loadInput(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

func splitByBlanklines(rawString string) []string {
	// Takes a giant string and splits into a list of passports
	groupList := strings.Split(rawString, "\n\n")
	return groupList
}

func splitByNewlines(rawString string) []string {
	splitStrings := strings.Split(rawString, "\n")
	return splitStrings
}

func getAllUniqueLetters(group []string) string {
	uniqueLetters := ""
	for _, individual := range group {
		for _, c := range individual {
			if strings.Count(uniqueLetters, string(c)) == 0 {
				uniqueLetters = uniqueLetters + string(c)
			}
		}
	}
	return uniqueLetters
}

func main() {
	fmt.Println("day06-01 started")
	groupList := splitByBlanklines(loadInput("input"))
	sum := 0

	for _, group := range groupList {
		sum += len(getAllUniqueLetters(splitByNewlines(group)))
	}

	fmt.Printf("sum is: %d\n", sum)
}
