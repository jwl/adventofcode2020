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

// func getAllContained(bagRules []string, bagType string) []string {
// 	// get all bagtypes immediately contained, then recursively search those to get those that they contain
// 	allPossibleContained := []string{}
// 	// get all immediate containers
// 	allPossibleContained = append(allPossibleContained, getImmediateContained(bagRules, bagType)...)

// 	indirectContained := []string{}
// 	for _, possibleContainerBagType := range allPossibleContained {
// 		indirectContained = append(indirectContained, getAllContained(bagRules, possibleContainerBagType)...)
// 	}

// 	// append indirectContainers to allPossibleContainers and remove all duplicates
// 	// allPossibleContained = removeDuplicates(append(allPossibleContained, indirectContained...))

// 	return allPossibleContained
// }

// func removeDuplicates(elements []string) []string {
// 	result := []string{}
// 	encountered := make(map[string]bool)
// 	for _, element := range elements {
// 		if !encountered[element] {
// 			result = append(result, element)
// 			encountered[element] = true
// 		}
// 	}
// 	return result
// }

// func getImmediateContained(bagRules []string, bagType string) []string {
// 	// search for the rule about what bags that this bagtype contains and return them
// 	immediateContained := []string{}
// 	for _, rule := range bagRules {
// 		// fmt.Printf("in getImmediateContainers, looking at rule <%s> for bagtype <%s>\n", rule, bagType)
// 		// if containsTarget(rule, bagType) {
// 		// 	immediateContained = append(immediateContained, getHolderType(rule))
// 		// }
// 		if getHolderType(rule) == bagType {
// 			return getHeldTypes(rule)
// 		}
// 	}

// 	return immediateContained
// }

func getTotalBags(rules []string, bagType string) int {
	// Note that you start with 1 bag; the top most bag
	totalBags := 1
	contentRules := findDefinition(rules, bagType)
	if len(contentRules) == 0 {
		return 1
	}
	for _, contentRule := range contentRules {
		// fmt.Printf("contentRule is: %s\n", contentRule)
		// quantity, _ := strconv.Atoi(strings.Split(contentRule, "")[0])
		quantity := getQuantity(contentRule)
		bagType := getBagType(contentRule)
		// fmt.Printf("quantity is: %d, bagType is: %s\n", quantity, bagType)
		bagsForBagType := getTotalBags(rules, bagType)
		// fmt.Printf("quantity is %d, bagType is: %s, bagsForBagType is: %d, total bags for this type is: %d\n", quantity, bagType, bagsForBagType, quantity*bagsForBagType)
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
	// fmt.Printf("match is: %#v\n", match)
	quantity, _ := strconv.Atoi(match[1])
	return quantity
}

func getBagType(ruleFragment string) string {
	// given input similar to "1 dark olive bag"
	// return the bag type, ie. "dark olive"
	re := regexp.MustCompile(`(\d) ([\w]+ [\w]+) bag[s]*`)
	match := re.FindStringSubmatch(ruleFragment)
	// fmt.Printf("match is: %#v\n", match)
	return match[2]
}

func getHolderType(rule string) string {
	// in a given rule, this returns the type of the bag that does the holding
	// ie. the type that appears to the left of 'contain'
	re := regexp.MustCompile(`^([\w]+ [\w]+)`)
	holderType := re.FindStringSubmatch(rule)

	return holderType[0]
}

func getHeldTypes(rule string) []string {
	splitRule := strings.Split(rule, " contain ")
	fmt.Printf("splitRule[1] is: %#v\n", splitRule[1])
	if strings.Contains(splitRule[1], "no other bags") {
		return []string{}
	}
	re := regexp.MustCompile(`(\d) ([\w]+ [\w]+) bag[s]*`)
	rawHeldTypes := re.FindAllStringSubmatch(rule, -1)

	fmt.Printf("rawHeldTypes is: <%#v>\n", rawHeldTypes)

	heldTypes := []string{}

	for _, rawHeldType := range rawHeldTypes {
		heldTypes = append(heldTypes, rawHeldType[2])
	}

	fmt.Printf("heldTypes is: <%#v>\n", heldTypes)
	return heldTypes
}

func main() {
	fmt.Println("day07-02 started")

	input := loadInputIntoListOfStrings("input")

	bags := getTotalBags(input, "shiny gold")
	fmt.Printf("Total number of bags inside is: %d\n", (bags - 1))

	// fmt.Printf("findDefinition for 'shiny gold' is: %#v\n", findDefinition(input, "shiny gold"))
	// fmt.Printf("getQuantity for '1 dark olive bag' is: %d\n", getQuantity("1 dark olive bag"))
	// fmt.Printf("getBagType for '1 dark olive bag' is: %s\n", getBagType("1 dark olive bag"))
}
