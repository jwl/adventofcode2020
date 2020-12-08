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

// returns 0 if there's an error, accumulator value if successful
func execute(s state, lines []string) int {
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
			return -1
		}

		// check if the next instruction is one we've already done
		if contains(s.history, s.current) {
			fmt.Printf("Hit an infinite loop! state is: %#v\n", s)
			return -1
		} else if s.current >= len(lines) {
			fmt.Printf("Reached the end of the program! state is: %#v\n", s)
			return s.accumulator
		}
	}

	// how did you get here?
	return 0
}

func modifyInstructions(original []string, i int, newInstruction string) []string {
	modifiedLines := original
	modifiedLines[i] = newInstruction + " " + getArgumentStr(original[i])
	return modifiedLines
}

func iterateProgram(lines []string) {
	for i, line := range lines {
		fmt.Printf("original lines are: %#v\n", lines)
		fmt.Printf("Checking line %d\n", i)
		if getInstruction(line) == "nop" {
			modifiedLines := modifyInstructions(lines, i, "jmp")
			accumulatorValue := execute(state{[]int{}, 0, 0}, modifiedLines)
			if accumulatorValue > 0 {
				fmt.Printf("Found successful modification! Change is at line %d and accumulator value is %d\n", i, accumulatorValue)
				return
			} else {
				lines = modifyInstructions(lines, i, "nop")
			}
		} else if getInstruction(line) == "jmp" {
			modifiedLines := modifyInstructions(lines, i, "nop")
			accumulatorValue := execute(state{[]int{}, 0, 0}, modifiedLines)
			if accumulatorValue > 0 {
				fmt.Printf("Found successful modification! Change is at line %d and accumulator value is %d\n", i, accumulatorValue)
				return
			} else {
				lines = modifyInstructions(lines, i, "jmp")
			}
		}
	}
	fmt.Printf("Hmmm, we hit the end of the program without finding a solution...\n")
}

func main() {
	fmt.Println("day08-02 started")
	input := loadInputIntoListOfStrings("input")

	iterateProgram(input)
}
