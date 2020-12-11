package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	strData := strings.Split(string(data), "\n")
	right := 3
	down := 1
	fmt.Println(part1(strData, right, down))
	fmt.Println(part2(strData))
}

func part1(strData []string, right int, down int) int {
	rowLen := len(strData[0])
	rowNum := 0
	colNum := 0

	treeCount := 0

	for rowNum < len(strData) && len(strData[rowNum]) > 0 {
		if colNum >= rowLen-1 {
			colNum = colNum % rowLen
		}
		if string(strData[rowNum][colNum]) == string('#') {
			treeCount++
		}
		rowNum += down
		colNum += right
	}

	return treeCount
}

func part2(strData []string) int {
	product := 1
	product *= part1(strData, 1, 1)
	product *= part1(strData, 3, 1)
	product *= part1(strData, 5, 1)
	product *= part1(strData, 7, 1)
	product *= part1(strData, 1, 2)
	return product
}
