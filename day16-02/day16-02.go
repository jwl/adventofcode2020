package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jwl/adventofcode2020/aocutils"
)

func splitStringToIntArray(str string) []int {
	s := strings.Split(str, ",")
	t := []int{}
	for _, v := range s {
		num, _ := strconv.Atoi(v)
		t = append(t, num)
	}
	return t
}

func isNumberValid(rules map[string][]int, target int) bool {
	// if a number doesn't fit in ANY rules, the entire ticket is invalid
	for _, rule := range rules {
		if aocutils.Contains(rule, target) {
			return true
		}
	}
	return false
}

func parseRawInput(input []string) (map[string][]int, []int, [][]int) {
	// returns rawRules, rawMyTicket and rawNearbyTickets in that order
	rules := make(map[string][]int)
	lineNumber := 0

	// begin parsing rules
	for i, line := range input {
		lineNumber = i
		if line == "" {
			break
		} else {
			re := regexp.MustCompile(`([\w]+[ \w]*): ([\d]+)-([\d]+) or ([\d]+)-([\d]+)`)
			match := re.FindStringSubmatch(line)
			fieldName := match[1]
			r1Lower, _ := strconv.Atoi(match[2])
			r1Upper, _ := strconv.Atoi(match[3])
			r2Lower, _ := strconv.Atoi(match[4])
			r2Upper, _ := strconv.Atoi(match[5])
			r1, r2 := []int{}, []int{}
			for x := r1Lower; x <= r1Upper; x++ {
				r1 = append(r1, x)
			}
			for y := r2Lower; y <= r2Upper; y++ {
				r2 = append(r2, y)
			}
			rules[fieldName] = append(r1, r2...)
		}

	}

	// get myTicket
	// myTicket := input[lineNumber+2]
	myTicket := splitStringToIntArray(input[lineNumber+2])

	// remaining lines should be nearby tickets
	rawNearbyTickets := input[lineNumber+5:]
	nearbyTickets := [][]int{}
	for _, rawTicket := range rawNearbyTickets {
		nearbyTickets = append(nearbyTickets, splitStringToIntArray(rawTicket))
	}

	return rules, myTicket, nearbyTickets
}

// func getPositionOfField(fieldName string, fieldRule []int, tickets [][]int, alreadyUsed []int) int {
// 	maxFieldNum := len(tickets[0])

// 	for i := 0; i < maxFieldNum; i++ {
// 		if aocutils.Contains(alreadyUsed, i) {
// 			continue
// 		}
// 		validPos := true

// 		// go through all tickets to see if ticket[i] matches fieldRule
// 		for _, ticket := range tickets {
// 			if !aocutils.Contains(fieldRule, ticket[i]) {
// 				validPos = false
// 				break
// 			}
// 		}

// 		if validPos {
// 			return i
// 		}
// 	}

// 	return -1
// }

func isValidBoard(fields map[string]int, rules map[string][]int, tickets [][]int) bool {
	// check if:
	// a. every field has a valid position
	// b. every field has a unique position
	fmt.Printf("\t Checking board of size %d: %#v ... ", len(fields), fields)
	alreadySeen := []int{}

	for fieldName, position := range fields {

		// check if unique
		if aocutils.Contains(alreadySeen, position) {
			fmt.Printf("INVALID\n")
			return false
		}

		alreadySeen = append(alreadySeen, position)

		// check if valid
		for _, ticket := range tickets {
			if !aocutils.Contains(rules[fieldName], ticket[position]) {
				fmt.Printf("INVALID\n")
				return false
			}
		}

	}

	fmt.Printf("GOOD!\n")
	return true
}

func solve(fields map[string]int, rules map[string][]int, tickets [][]int, usedNames []string) map[string]int {
	// recursive!
	fmt.Printf("\tlen(fields): %d, usedNames: %#v\n", len(fields), usedNames)
	if len(fields) >= len(rules) {
		fmt.Printf("Reached end condition! Returning...\n")
		return fields
	}

	maxField := len(tickets[0])
	fmt.Printf("maxField is: %d\n", maxField)

	tmpFields := make(map[string]int)
	// copy fields map
	for key, value := range fields {
		tmpFields[key] = value
	}

	for fieldName := range rules {
		if aocutils.ContainsString(usedNames, fieldName) {
			continue
		}
		usedNames = append(usedNames, fieldName)
		for i := 0; i < maxField; i++ {
			tmpFields[fieldName] = i
			if isValidBoard(tmpFields, rules, tickets) {
				tmpFields = solve(tmpFields, rules, tickets, usedNames)
				if len(tmpFields) == len(rules) {
					return tmpFields
				}
			}

		}
	}

	return fields
}

func partTwo(input []string) int {
	fmt.Printf("beginning partTwo...\n")
	rules, myTicket, tickets := parseRawInput(input)

	// fmt.Printf("rules are: %#v\n", rules)
	// fmt.Printf("myTicket is: %#v\n", myTicket)
	// fmt.Printf("tickets are: %#v\n", tickets)

	fields := make(map[string]int)

	fields = solve(fields, rules, tickets, []string{})
	// alreadyUsed := []int{}

	// for fieldName, fieldRule := range rules {
	// 	fields[fieldName] = getPositionOfField(fieldName, fieldRule, tickets, alreadyUsed)
	// 	alreadyUsed = append(alreadyUsed, fields[fieldName])
	// 	fmt.Printf("position of %s is %d\n", fieldName, fields[fieldName])
	// }

	result := 1

	for fieldName, fieldNum := range fields {
		if strings.Contains(fieldName, "departure") {
			// departureValue, _ := strconv.Atoi(myTicket)
			result *= myTicket[fieldNum]
		}
	}

	return result
}

func partTwoAlt(input []string) int {
	rules, myTicket, tickets := parseRawInput(input)
	var validTickets [][]int

search:
	for _, ticket := range tickets {
	check:
		for _, posValue := range ticket {
			for _, rule := range rules {
				if aocutils.Contains(rule, posValue) {
					continue check
				}
			}

			// invalid value
			continue search
		}

		validTickets = append(validTickets, ticket)
	}
	if len(validTickets) == 0 {
		panic("no valid tickets")
	}

	// For each rule, which columns could it be?
	fieldLocationCandidates := map[string]map[int]struct{}{}
	for fieldName, fieldRule := range rules {
		if _, ok := fieldLocationCandidates[fieldName]; !ok {
			fieldLocationCandidates[fieldName] = map[int]struct{}{}
		}

	fieldNumberSearch:
		for fieldNumber := 0; fieldNumber < len(myTicket); fieldNumber++ {
			for _, t := range validTickets {
				if !aocutils.Contains(fieldRule, t[fieldNumber]) {
					continue fieldNumberSearch
				}
			}

			fieldLocationCandidates[fieldName][fieldNumber] = struct{}{}
		}
	}

	// while there are more candidates
	fieldLocations := map[string]int{}
	for len(fieldLocationCandidates) > 0 {
		for idx, candidates := range fieldLocationCandidates {
			//   find a candidate with only one valid location
			if len(candidates) == 1 {
				//   remember its location
				location := -1
				for location = range candidates {
				}

				fieldLocations[idx] = location

				//   remove its location from consideration of all other candidates
				for k := range fieldLocationCandidates {
					delete(fieldLocationCandidates[k], location)
				}

				delete(fieldLocationCandidates, idx)
				break
			}
		}
	}

	checksum := 1
	for name, location := range fieldLocations {
		if strings.HasPrefix(name, "departure") {
			checksum *= myTicket[location]
		}
	}

	return checksum
}

func main() {
	fmt.Println("day16-02 started")

	// input := aocutils.LoadInputIntoListOfStrings("sample2")
	input := aocutils.LoadInputIntoListOfStrings("input")

	// fmt.Printf("Ticket scanning error rate is: %d\n", partOne(input))
	// fmt.Printf("Product of the 6 departure field values: %d\n", partTwo(input))
	fmt.Printf("Product of the 6 departure field values: %d\n", partTwoAlt(input))

}
