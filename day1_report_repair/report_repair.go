package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	numbers := strings.Split(string(data), "\n")
	numbers = numbers[:len(numbers)-1]
	part1(numbers)
	part2(numbers)
}

func part1(numbers []string) {
	numSet := make(map[int]bool)

	for _, num := range numbers {
		i, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Failed to convert to int")
		}

		if _, ok := numSet[2020-i]; ok {
			fmt.Println(i * (2020 - i))
			break
		} else {
			numSet[i] = true
		}
	}
}

func part2(numbers []string) {
	numSet := make(map[int]bool)
	numList := []int{}

	for _, num := range numbers {
		i, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Failed to convert to int:", num, err)
			continue
		}
		numSet[i] = true
		numList = append(numList, i)
	}

	for index1, num1 := range numList {
		for _, num2 := range numList[index1:] {
			if _, ok := numSet[2020-num1-num2]; ok {
				fmt.Println(num1 * num2 * (2020 - num1 - num2))
				return
			}
		}
	}

}
