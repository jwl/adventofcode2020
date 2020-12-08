package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type state struct {
	history     []int
	accumulator int
	current     int
}

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

func getInstruction(line string) string {
	re := regexp.MustCompile(`([\w]+) ([+-][\d]+)$`)
	match := re.FindStringSubmatch(line)
	return match[1]
}

func getArgument(line string) int {
	re := regexp.MustCompile(`([\w]+) ([+-][\d]+)$`)
	match := re.FindStringSubmatch(line)
	i, _ := strconv.Atoi(match[2])
	return i
}

func getArgumentStr(line string) string {
	re := regexp.MustCompile(`([\w]+) ([+-][\d]+)$`)
	match := re.FindStringSubmatch(line)
	return match[2]
}

func contains(list []int, i int) bool {
	for _, e := range list {
		if e == i {
			return true
		}
	}
	return false
}

// returns 0 if there's an error, 1 if no issues
func execute(s state, lines []string) int {
	// while loop
	stillRunning := true
	for stillRunning {
		s.history = append(s.history, s.current)

		switch getInstruction(lines[s.current]) {
		case "nop":
			s.current++
		case "acc":
			s.accumulator += getArgument(lines[s.current])
			s.current++
		case "jmp":
			s.current += getArgument(lines[s.current])
		default:
			fmt.Printf("Unknown instruction: <%s>\nEnding program!\n", lines[s.current])
			return 0
		}

		// check if the next instruction is one we've already done
		if contains(s.history, s.current) {
			fmt.Printf("Hit an infinite loop! state is: %#v\n", s)
			return 0
		} else if s.current >= len(lines) {
			fmt.Printf("Reached the end of the program! state is: %#v\n", s)
			return 1
		}
	}

	// how did you get here?
	return 0
}

func main() {
	fmt.Println("day08-01 started")
	input := loadInputIntoListOfStrings("input")

	var initialState state = state{[]int{}, 0, 0}

	execute(initialState, input)
}
