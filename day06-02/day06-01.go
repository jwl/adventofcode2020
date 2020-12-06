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
		// fmt.Printf("individual is: %s\n", individual)
		for _, c := range individual {
			// fmt.Printf("c is: %s\n", string(c))
			if strings.Count(uniqueLetters, string(c)) == 0 {
				uniqueLetters = uniqueLetters + string(c)
			}
		}
	}
	// fmt.Printf("for group %#v, uniqueLetters are: %s\n", group, uniqueLetters)
	return uniqueLetters
}

func main() {
	fmt.Println("day06-01 started")
	// boardingPasses := loadInput("input")
	// fmt.Printf("Missing seatID is: %d\n", getMissingSeatID(boardingPasses))

	groupList := splitByBlanklines(loadInput("input"))

	sum := 0

	for _, group := range groupList {
		// fmt.Printf("group is:\n%s\n", group)
		// for _, individual := range splitByNewlines(group) {
		// 	fmt.Printf("\tindividual in group is: %s\n", individual)
		// }
		fmt.Printf("uniqueLetters of this group are: %s\n\n", getAllUniqueLetters(splitByNewlines(group)))
		sum += len(getAllUniqueLetters(splitByNewlines(group)))
	}

	fmt.Printf("sum is: %d\n", sum)
}
