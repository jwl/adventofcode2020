package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

func main() {
	fmt.Println("day11-02 started")
}
