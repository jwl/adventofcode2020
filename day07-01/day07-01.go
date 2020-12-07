package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func loadInputIntoListOfStrings(filename string) []string {
	input := []string{}

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
		input = append(input, snl.Text())
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func loadInputIntoGiantString(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

func splitByBlanklines(rawString string) []string {
	groupList := strings.Split(rawString, "\n\n")
	return groupList
}

func splitByNewlines(rawString string) []string {
	splitStrings := strings.Split(rawString, "\n")
	return splitStrings
}

func getAllContainers(bagRules []string, bagType string) []string {
	allPossibleContainers := []string{}
	// get all immediate containers
	allPossibleContainers = append(allPossibleContainers, getImmediateContainers(bagRules, bagType)...)

	indirectContainers := []string{}
	for _, possibleContainerBagType := range allPossibleContainers {
		indirectContainers = append(indirectContainers, getAllContainers(bagRules, possibleContainerBagType)...)
	}

	// append indirectContainers to allPossibleContainers and remove all duplicates
	allPossibleContainers = removeDuplicates(append(allPossibleContainers, indirectContainers...))

	return allPossibleContainers
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

func getImmediateContainers(bagRules []string, bagType string) []string {
	immediateContainers := []string{}
	for _, rule := range bagRules {
		if containsTarget(rule, bagType) {
			immediateContainers = append(immediateContainers, getHolderType(rule))
		}
	}

	return immediateContainers
}

func containsTarget(rule string, bagType string) bool {
	splitRule := strings.Split(rule, "contain")
	return strings.Contains(splitRule[1], bagType)
}

func getHolderType(rule string) string {
	// in a given rule, this returns the type of the bag that does the holding
	// ie. the type that appears to the left of 'contain'
	re := regexp.MustCompile(`^([\w]+ [\w]+)`)
	holderType := re.FindStringSubmatch(rule)

	return holderType[0]
}

func main() {
	fmt.Println("day07-01 started")
	input := loadInputIntoListOfStrings("input")
	allContainers := getAllContainers(input, "shiny gold")
	fmt.Printf("getAllContainers for input is: %d\n", len(allContainers))
}
