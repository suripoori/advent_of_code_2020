package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
        return
	}
	seats := strings.Split(string(data), "\n")
	maxSeatId, mySeatId := part1and2(seats)
	fmt.Println(maxSeatId)
	fmt.Println(mySeatId)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func getSeatId(seat string) int {
	row := seat[:8]
	minRow := 0
	maxRow := 128
	for _, c := range(row) {
		if string(c) == string('B') {
			minRow = int((minRow + maxRow) / 2)
		} else if string(c) == string('F') {
			maxRow = int((minRow + maxRow) / 2)
		}
	}
	rowNum := minRow
	minCol := 0
	maxCol := 8
	for _, c := range(seat[7:]) {
		if string(c) == string('R') {
			minCol = int((minCol + maxCol) / 2)
		} else if string(c) == string('L') {
			maxCol = int((minCol + maxCol) / 2)
		}
	}
	colNum := minCol
	seatId := rowNum * 8 + colNum
	return seatId
}

func part1and2(seats []string) (int, int) {
	maxSeatId := 0
	allSeatIds := make(map[int]bool)
	for _, seat := range(seats[:len(seats)-1]) {
		seatId := getSeatId(seat)
		maxSeatId = max(maxSeatId, seatId)
		allSeatIds[seatId] = true
	}

	for i := 1; i <= 1023; i++ {
		_, present := allSeatIds[i]
		if present {
			continue
		}
		_, present_minus_1 := allSeatIds[i-1]
		if !present_minus_1 {
			continue
		}
		_, present_plus_1 := allSeatIds[i+1]
		if !present_plus_1 {
			continue
		}
		// If both myseatId + 1 and mySeatId - 1 are present and mySeatId is not present
		// then that is mySeatId
		return maxSeatId, i
	}
	return -1, -1 
}