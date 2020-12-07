package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

func getTotalBags(rules []string, bagType string) int {
	// Note that you start with 1 bag; the top most bag
	totalBags := 1
	contentRules := findDefinition(rules, bagType)
	if len(contentRules) == 0 {
		return 1
	}
	for _, contentRule := range contentRules {
		quantity := getQuantity(contentRule)
		bagType := getBagType(contentRule)
		bagsForBagType := getTotalBags(rules, bagType)
		totalBags += quantity * bagsForBagType

	}
	fmt.Printf("for bagtype: %s, total number of bags inside is: %d\n", bagType, totalBags)
	return totalBags
}

func findDefinition(rules []string, bagType string) []string {
	// finds definition for a given bagtype and returns what it contains
	// as an array of strings denoting quantity and type
	// ie. ["1 light blue", "3 dark green"]
	// if this definition contains no other bags, return empty array
	contents := []string{}
	for _, rule := range rules {
		splitRule := strings.Split(rule, "contain")

		// if left-hand side contains bagType, split the right-hand side and return as an array
		if strings.Contains(splitRule[0], bagType) {
			re := regexp.MustCompile(`(\d) ([\w]+ [\w]+) bag[s]*`)
			rawContents := re.FindAllStringSubmatch(strings.TrimSpace(splitRule[1]), -1)

			for _, quantityAndType := range rawContents {
				contents = append(contents, quantityAndType[0])
			}
		}
	}
	return contents
}

func getQuantity(ruleFragment string) int {
	// given input similar to "1 dark olive bag"
	// return the quantity, ie. 1
	if strings.Contains(ruleFragment, "no other bags") {
		return 1
	}
	re := regexp.MustCompile(`(\d) ([\w]+ [\w]+) bag[s]*`)
	match := re.FindStringSubmatch(ruleFragment)
	quantity, _ := strconv.Atoi(match[1])
	return quantity
}

func getBagType(ruleFragment string) string {
	// given input similar to "1 dark olive bag"
	// return the bag type, ie. "dark olive"
	re := regexp.MustCompile(`(\d) ([\w]+ [\w]+) bag[s]*`)
	match := re.FindStringSubmatch(ruleFragment)
	return match[2]
}

func main() {
	fmt.Println("day07-02 started")

	input := loadInputIntoListOfStrings("input")

	bags := getTotalBags(input, "shiny gold")
	fmt.Printf("Total number of bags inside is: %d\n", (bags - 1))
}
