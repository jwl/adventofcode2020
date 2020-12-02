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

func loadInput() []string {
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

func isValid(line string) bool {
	// come up with a regex that
	re := regexp.MustCompile(`(\d+)-(\d+) (\w): ([a-z]+)`)
	match := re.FindStringSubmatch(line)
	p1, _ := strconv.Atoi(match[1])
	p2, _ := strconv.Atoi(match[2])
	targetLetter := match[3]
	candidatePassword := match[4]

	// implement day02-02 logic for line validation
	combined := string(candidatePassword[p1-1]) + string(candidatePassword[p2-1])
	if strings.Count(combined, targetLetter) == 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println("day02-02 started")
	input := loadInput()
	valid := 0

	for _, e1 := range input {
		if isValid(e1) {
			valid++
		}
	}

	fmt.Println("End of input reached!")
	fmt.Println("Total number of valid passwords is:", valid)
}
