package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)


func loadInput(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}


func splitPassports(rawString string) []string {
	// Takes a giant string and splits into a list of passports
	passportList := strings.Split(rawString, "\n\n")
	return passportList
}


func validatePassports(passportList []string) int {
	// Returns number of valid passports in passportList
	validPassports := 0
	for _, passport := range passportList {
		if validatePassport(passport) {
			validPassports++
		}
	}

	return validPassports
}


func validatePassport(rawPassport string) bool {
	// A valid passport must have:
	// byr, iyr, eyr, hgt, hcl, ecl, pid
	requiredFields := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	passport := strings.Fields(rawPassport)

	for _, field := range requiredFields {
		if !containsField(field, passport) {
			fmt.Printf("Field <%s> is missing from passport <%v>\n", field, passport)
			return false
		}
	}

	return true
}


func containsField(targetField string, passport []string) bool {
	for _, entry := range passport {
		if targetField == entry[0:3] {
			return true
		}
	}
	return false
}


func main() {
	fmt.Println("day04-01 started")
	rawPassportList := loadInput("input")
	passportList := splitPassports(rawPassportList)
	validPassports := validatePassports(passportList)

	fmt.Printf("Total number of valid passports is: %d\n", validPassports)
}
