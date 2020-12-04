package main

import (
	"fmt"
	"regexp"
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
		if !containsValidField(field, passport) {
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

func containsValidField(targetField string, passport[]string) bool {
	rules := make(map[string]string)
	rules["byr"] = `\b19[2-9]\d\b|\b200[0-2]\b`
	rules["iyr"] = `\b201\d\b|\b2020\b`
	rules["eyr"] = `\b202\d\b|\b2030\b`
	rules["hgt"] = `\b(1[5-8]\d|19[0-3])cm\b|\b(59|6\d|7[0-6])in\b`
	rules["hcl"] = `#[\da-f]{6}\b`
	rules["ecl"] = `\b(amb|blu|brn|gry|grn|hzl|oth)\b`
	rules["pid"] = `\b\d{9}\b`

	for _, entry := range passport {
		if targetField == entry[0:3] {
			// entry[4:] is the value of the field

			match, _ := regexp.MatchString(rules[targetField], entry[4:])
			return match
		}
	}

	return false
}


func main() {
	fmt.Println("day04-02 started")
	rawPassportList := loadInput("input")
	passportList := splitPassports(rawPassportList)
	validPassports := validatePassports(passportList)

	fmt.Printf("Total number of valid passports is: %d\n", validPassports)
}
