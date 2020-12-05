package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"encoding/hex"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
        return
	}
	passports := strings.Split(string(data), "\n\n")
	fmt.Println(part1(passports))
	fmt.Println(part2(passports))
}

func isValid(passport string, validateData bool) bool {
	
	requiredKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} // skip cid

	fields := strings.Split(passport, "\n")
	passportKeyValues := make([]string, 0)
	for _, field := range(fields) {
		if strings.Contains(field, " ") {
			passportKeyValues = append(passportKeyValues, strings.Split(field, " ")...)
		} else {
			passportKeyValues = append(passportKeyValues, field)
		}
	}

	passportKeys := make(map[string]string)
	for _, keyValue := range(passportKeyValues) {
		if strings.Contains(keyValue, ":") {
			key := strings.Split(keyValue, ":")[0]
			value := strings.Split(keyValue, ":")[1]
			passportKeys[key] = value
		}
	}

	for _, requiredKey := range(requiredKeys) {
		value, present := passportKeys[requiredKey]
		if !present {
			return false
		}
		if validateData && !isValidData(requiredKey, value) {
			return false
		}
	}
	return true
}

func isValidData(key string, data string) bool {
	switch key {
	case "byr":
		if len(data) != 4 {
			fmt.Print("Invalid byr len")
			return false
		}
		value, err := strconv.Atoi(data)
		if err != nil {
			fmt.Println("Invalid byr")
			return false
		}
		return value >= 1920 && value <= 2002
	case "iyr":
		if len(data) != 4 {
			fmt.Print("Invalid iyr len")
			return false
		}
		value, err := strconv.Atoi(data)
		if err != nil {
			fmt.Println("Invalid iyr")
			return false
		}
		return value >= 2010 && value <= 2020
	case "eyr":
		if len(data) != 4 {
			fmt.Print("Invalid eyr len")
			return false
		}
		value, err := strconv.Atoi(data)
		if err != nil {
			fmt.Println("Invalid eyr")
			return false
		}
		return value >= 2020 && value <= 2030
	case "hgt":
		value, err := strconv.Atoi(data[:len(data)-2])
		if err != nil {
			fmt.Println("Invalid hgt")
			return false
		}
		if string(data[len(data)-2:]) == string("cm") {
			return value >= 150 && value <= 193
		} else if string(data[len(data)-2:]) == string("in") {
			return value >= 59 && value <= 76
		}
		fmt.Println("Unknown height") 
		return false
	case "hcl":
		if string(data[0]) != string('#') {
			fmt.Println("First character is not #")
			return false
		}
		if len(data) != 7 {
			fmt.Println("Should contain exactly 6 digits after #")
			return false
		}
		_, err := hex.DecodeString(data[1:])
		if err != nil {
			fmt.Println("Invalid hex")
			return false
		}
		return true
	case "ecl":
		possibleColors := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
		_, present := possibleColors[data]
		if !present {
			fmt.Println("Invalid ecl")
		}
		return present
	case "pid":
		if len(data) != 9 {
			fmt.Println("Passport id should be a 9 digit number")
			return false
		}
		_, err := strconv.Atoi(data)
		if err != nil {
			fmt.Println("Unable to convert pid to a number")
			return false
		}
		return true
	default: 
		fmt.Println("Unknown key")
		return true
	}
}

func part1(passports []string) int {
	validCount := 0
	for _, passport := range(passports) {
		if isValid(passport, false) {
			validCount++
		}
	}
	return validCount
}

func part2(passports []string) int {
	validCount := 0
	for _, passport := range(passports) {
		if isValid(passport, true) {
			validCount++
		}
	}
	return validCount
}