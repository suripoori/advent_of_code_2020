package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
	"errors"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
        return
	}
	strData := strings.Split(string(data), "\n")
	part1(strData)
	part2(strData)
}

func parseRow(str string) (int, int, string, string, error) {
	num1, err := strconv.Atoi(strings.Split(str, string('-'))[0])
	if err != nil {
		fmt.Println("Failed to convert num1")
		return -1, -1, string(' '), string(' '), errors.New("Failed to convert num1")
	}
	num2, err := strconv.Atoi(strings.Split(strings.Split(str, string(' '))[0], string('-'))[1])
	if err != nil {
		fmt.Println("Failed to convert num2")
		return -1, -1, string(' '), string(' '), errors.New("Failed to convert num2")
	}
	ch := string(strings.Split(str, string(' '))[1][0])
	password := strings.Split(str, string(' '))[2]

	return num1, num2, ch, password, nil
}

func part1(strData []string) {
	validPasswordCount := 0
	for _, str := range(strData) {
		minNum, maxNum, ch, password, err := parseRow(str)
		if err != nil {
			fmt.Print("Failing part 1")
			return
		}

		count := 0
		for _, char := range password {
			if string(char) == ch {
				count++
			}
		}

		if minNum <= count && count <= maxNum {
			validPasswordCount++
		}
	}

	fmt.Println(validPasswordCount)
}

func part2(strData []string) {
	validPasswordCount := 0
	for _, str := range(strData) {
		firstIndex, secondIndex, ch, password, err := parseRow(str)
		if err != nil {
			fmt.Print("Failing part 1")
			return
		}

		// Convert to zero based indexing
		firstIndex--
		secondIndex--

		validPass := false
		for _, index := range([]int{firstIndex, secondIndex}) {
			if string(password[index]) == ch {
				if validPass {
					validPass = false
					break
				} else {
					validPass = true
				}
			}
		}

		if validPass {
			validPasswordCount++
		}
	}

	fmt.Println(validPasswordCount)
}