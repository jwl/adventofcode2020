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

func load_input() []string {
	path := "input"
	input := []string{}

	buf, err := os.Open(path)
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

func is_valid(line string) bool {
	// come up with a regex that
	re := regexp.MustCompile(`(\d+)-(\d+) (\w): ([a-z]+)`)
	match := re.FindStringSubmatch(line)
	min, _ := strconv.Atoi(match[1])
	max, _ := strconv.Atoi(match[2])
	target_letter := match[3]
	candidate_password := match[4]

	target_count := strings.Count(candidate_password, target_letter)
	if target_count >= min && target_count <= max {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("day02-01 started")
	input := load_input()
	valid := 0
	for _, e1 := range input {
		if is_valid(e1) {
			valid++
		}

		// fmt.Println("Line", i+1, ": ", e1)
		// fmt.Println("Is this password valid? ", is_valid(e1))

		// TODO: once you can parse one line, let the loop process every line in the file
	}
	fmt.Println("End of input reached!")
	fmt.Println("Total number of valid passwords is:", valid)
}
