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

func getAllCommonLetters(group []string) string {
	commonLetters := strings.Split(group[0], "")

	for _, individual := range group {
		commonLetters = intersection(commonLetters, strings.Split(individual, ""))
		if len(commonLetters) < 1 {
			return ""
		}
	}

	return strings.Join(commonLetters, "")
}

func intersection(s1 []string, s2 []string) []string {
	result := []string{}
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			result = append(result, e)
		}
	}
	return removeDuplicates(result)
}

func removeDuplicates(elements []string) []string {
	result := []string{}
	encountered := make(map[string]bool)
	for _, element := range elements {
		if !encountered[element] {
			result = append(result, element)
			encountered[element] = true
		}
	}
	return result
}

func main() {
	fmt.Println("day06-01 started")
	groupList := splitByBlanklines(loadInput("input"))
	sum := 0

	for _, group := range groupList {
		sum += len(getAllCommonLetters(splitByNewlines(group)))
	}

	fmt.Printf("sum is: %d\n", sum)
}
