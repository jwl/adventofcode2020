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

func parseRawInput(input []string) (map[string][]int, string, [][]int) {
	// returns rawRules, rawMyTicket and rawNearbyTickets in that order
	rules := make(map[string][]int)
	lineNumber := 0

	// begin parsing rules
	for i, line := range input {
		lineNumber = i
		if line == "" {
			break
		} else {
			re := regexp.MustCompile(`([\w]+): ([\d]+)-([\d]+) or ([\d]+)-([\d]+)`)
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
	myTicket := input[lineNumber+2]

	// remaining lines should be nearby tickets
	rawNearbyTickets := input[lineNumber+5:]
	nearbyTickets := [][]int{}
	for _, rawTicket := range rawNearbyTickets {
		nearbyTickets = append(nearbyTickets, splitStringToIntArray(rawTicket))
	}

	return rules, myTicket, nearbyTickets
}

func partOne(input []string) int {
	fmt.Printf("beginning partOne...\n")
	rules, myTicket, nearbyTickets := parseRawInput(input)

	fmt.Printf("rules are: %#v\n", rules)
	fmt.Printf("myTicket is: %s\n", myTicket)
	fmt.Printf("nearbyTickets are: %#v\n", nearbyTickets)

	ticketScanningErrorRate := 0
	validTickets := [][]int{}
	for _, ticket := range nearbyTickets {
		validTicketFlag := true

		for _, num := range ticket {
			if !isNumberValid(rules, num) {
				fmt.Printf("found invalid num %d in ticket: %#v\n", num, ticket)
				ticketScanningErrorRate += num
				validTicketFlag = false
				break
			}
		}

		if validTicketFlag {
			validTickets = append(validTickets, ticket)
		}
	}

	fmt.Printf("List of valid tickets is:\n")
	for _, vTicket := range validTickets {
		for i, num := range vTicket {
			fmt.Printf("%d", num)
			if i < len(vTicket)-1 {
				fmt.Printf(",")
			}
		}
		fmt.Println()
		// fmt.Println(vTicket)
	}

	fmt.Printf("ticketScanningErrorRate is: %d\n", ticketScanningErrorRate)
	return ticketScanningErrorRate
}

func main() {
	fmt.Println("day16-01 started")

	input := aocutils.LoadInputIntoListOfStrings("sample1")

	fmt.Printf("Ticket scanning error rate is: %d\n", partOne(input))
}
