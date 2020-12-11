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
	numbersList := make([]int, 0)
	for _, number := range numbers {
		intNumber, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Could not convert value to int: ", number)
		}
		numbersList = append(numbersList, intNumber)
	}

	fmt.Println(getInvalidNumber(25, numbersList))
	fmt.Println(getEncryptionWeakness(25, numbersList))
}

func getInvalidNumber(preambleLength int, numbersList []int) int {
	preamble := make([]int, 0)

	for _, number := range numbersList {
		if len(preamble) == preambleLength {
			if !isNumberValid(number, preamble) {
				return number
			}
		}
		preamble = append(preamble, number)
		if len(preamble) > preambleLength {
			preamble = preamble[1:]
		}
	}
	return -1
}

func isNumberValid(number int, preamble []int) bool {
	preambleMap := make(map[int]int)
	for i, n := range preamble {
		preambleMap[n] = i
	}

	for i, n := range preamble {
		index, present := preambleMap[number-n]
		if present && index != i {
			return true
		}
	}
	return false
}

func getContiguousList(invalidNumber int, numbersList []int) []int {
	contiguousList := make([]int, 0)
	sum := 0
	var popped int
	for _, number := range numbersList {
		if sum+number < invalidNumber {
			contiguousList = append(contiguousList, number)
			sum += number
		} else if sum+number == invalidNumber {
			contiguousList = append(contiguousList, number)
			return contiguousList
		} else {
			for len(contiguousList) > 0 {
				popped, contiguousList = contiguousList[0], contiguousList[1:]
				sum -= popped
				if sum+number < invalidNumber {
					contiguousList = append(contiguousList, number)
					sum += number
					break
				} else if sum+number == invalidNumber {
					contiguousList = append(contiguousList, number)
					sum += number
					return contiguousList
				}
			}
		}
	}
	return make([]int, 0)
}
func getEncryptionWeakness(preambleLength int, numbersList []int) int {
	invalidNumber := getInvalidNumber(preambleLength, numbersList)
	contiguousList := getContiguousList(invalidNumber, numbersList)
	var min int
	var max int
	for i, number := range contiguousList {
		if i == 0 || min > number {
			min = number
		}
		if i == 0 || max < number {
			max = number
		}
	}
	return min + max
}
