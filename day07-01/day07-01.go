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
	// fmt.Printf("### getAllContainers called, looking for all containers that hold bagtype: <%s>\n", bagType)
	allPossibleContainers := []string{}
	// get all immediate containers
	allPossibleContainers = append(allPossibleContainers, getImmediateContainers(bagRules, bagType)...)

	// call getAllContainers on all immediate containers
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
		// fmt.Printf("in getImmediateContainers, looking at rule <%s> for bagtype <%s>\n", rule, bagType)
		if containsTarget(rule, bagType) {
			immediateContainers = append(immediateContainers, getHolderType(rule))
		}
	}

	// fmt.Printf("for bagtype %s, immediateContainers is %#v\n", bagType, immediateContainers)

	return immediateContainers
}

func containsTarget(rule string, bagType string) bool {
	// fmt.Printf("\tchecking if rule <%s> contains bagtype <%s>\n", rule, bagType)
	// in a given rule, this returns whether the rule has the bagtype on the right hand side of the string
	splitRule := strings.Split(rule, "contain")

	// fmt.Printf("rule for containsTarget: %s\n", rule)
	// fmt.Printf("bagType for containsTarget: %#v\n", bagType)
	// fmt.Printf("splitRule for containsTarget: %#v\n", splitRule)

	return strings.Contains(splitRule[1], bagType)
}

func getHolderType(rule string) string {
	// in a given rule, this returns the type of the bag that does the holding
	// ie. the type that appears to the left of 'contain'
	// fmt.Printf("\tin getHolderType, rule to get Holder type of is: <%s>\n", rule)
	// splitRule := strings.Split(rule, "contain")
	// for _, s := range splitRule {
	// 	fmt.Printf("string in splitRule: %s\n", s)
	// }
	re := regexp.MustCompile(`^([\w]+ [\w]+)`)
	holderType := re.FindStringSubmatch(rule)

	// fmt.Printf("holderType: %#v\n", holderType)

	return holderType[0]
}

func main() {
	input := loadInputIntoListOfStrings("input")

	fmt.Println("day06-01 started")
	// fmt.Println("for sample_input, the correct answer is 4")

	// fmt.Printf("The number of bagtypes that can hold a shiny gold bag is: %d\n", len(getAllContainers(input, "shiny gold")))

	// fmt.Printf("getHolderType: %#v\n", containsTarget("light red bags contain 1 bright white bag, 2 muted yellow bags.", "muted yellow"))
	allContainers := getAllContainers(input, "shiny gold")
	// fmt.Printf("allContainers is: %#v\n", allContainers)
	fmt.Printf("getAllContainers for input is: %d\n", len(allContainers))
}
