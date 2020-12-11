package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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
	numbersList := make([]int, 0)
	for _, number := range numbers {
		intNumber, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Could not convert value to int: ", number)
		}
		numbersList = append(numbersList, intNumber)
	}

	fmt.Println(part1(numbersList))
	fmt.Println(numWays(numbersList))
}

func getOneJoltTwoJoltThreeJolt(numbersList []int) (int, int, int) {
	sort.Ints(numbersList)
	oneJolt, twoJolt, threeJolt := 0, 0, 0

	if len(numbersList) == 0 {
		return oneJolt, twoJolt, threeJolt
	}

	deviceJolt := numbersList[len(numbersList)-1] + 3
	numbersList = append(numbersList, deviceJolt)
	lastNumber := 0
	for _, number := range numbersList {
		switch number - lastNumber {
		case 0:
			fmt.Println("Same joltage numbers!")
			return oneJolt, twoJolt, threeJolt
		case 1:
			oneJolt++
		case 2:
			twoJolt++
		case 3:
			threeJolt++
		default:
			fmt.Println("Difference more than 3 jolts! ", number, lastNumber)
			return oneJolt, twoJolt, threeJolt
		}
		lastNumber = number
	}
	return oneJolt, twoJolt, threeJolt
}

func part1(numbersList []int) int {
	oneJolt, _, threeJolt := getOneJoltTwoJoltThreeJolt(numbersList)
	return oneJolt * threeJolt
}

func numWays(numbersList []int) int {
	sort.Ints(numbersList)
	numbersList = append(make([]int, 1), numbersList...)
	numbersList = append(numbersList, numbersList[len(numbersList)-1]+3)
	ways := make([]int, len(numbersList))
	exists := make(map[int]int)
	for index, number := range numbersList {
		exists[number] = index
	}
	// There's only one way to get to the device
	// from the previous adapter.
	// But there can be up to 3 ways to get to previous adapter
	// if all 3 values exist.
	ways[len(numbersList)-1] = 1
	for i := len(numbersList) - 2; i >= 0; i-- {
		totalWays := 0
		for j := 1; j <= 3; j++ {
			if index, ok := exists[numbersList[i]+j]; ok {
				// ways[index] should already be filled by the
				// time we get here because index is always > i
				totalWays += ways[index]
			}
		}
		ways[i] = totalWays
	}
	// The source joltage is 0
	// So, total ways is ways[1] + ways[2] + ways[3]
	// because we can go up to 3 joltages higher from 0
	// if such adapters exist.
	totalWays := 0
	for joltage := 1; joltage <= 3; joltage++ {
		if index, ok := exists[joltage]; ok {
			totalWays += ways[index]
		}
	}
	return totalWays
}
