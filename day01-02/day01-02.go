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
	for i, e1 := range input {
		for j, e2 := range input[i:] {
			for _, e3 := range input[j:] {
				if e1+e2+e3 == 2020 {
					fmt.Println("The three elements are: ", e1, ", ", e2, " and ", e3)
					fmt.Println("Their product is: ", e1*e2*e3)
					return
				}
			}
		}
	}
	fmt.Println("End of input reached!")
}
