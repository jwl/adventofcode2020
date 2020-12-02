package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func load_input() []int {
	path := "input.txt"
	input := []int{}

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
		// fmt.Println(snl.Text())
		i, _ := strconv.Atoi(snl.Text())
		input = append(input, i)
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func main() {
	fmt.Println("day01-01 started")
	input := load_input()
	for i, element := range input {
		for _, inner_element := range input[i:] {
			if element+inner_element == 2020 {
				fmt.Println("The two elements are: ", element, " and ", inner_element)
				fmt.Println("Their product is: ", element*inner_element)

			}

		}

	}
	fmt.Println("End of input reached!")
}
