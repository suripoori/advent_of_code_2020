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
	fmt.Println(getNumOccupiedSeats(getInitialState(data), false, 4))
	fmt.Println(getNumOccupiedSeats(getInitialState(data), true, 5))
}

func getInitialState(data []byte) [][]rune {
	rows := strings.Split(string(data), "\n")
	initialState := make([][]rune, len(rows)-1)

	for i, row := range rows[:len(rows)-1] {
		initialState[i] = []rune(row)
	}
	return initialState
}

func printState(state [][]rune) {
	fmt.Println("\n-------------------------")
	for _, row := range state {
		fmt.Println(string(row))
	}
	fmt.Println("-------------------------")
}

func getAdjacentOccupiedSeats(row int, col int, state [][]rune) int {
	adjCount := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i >= len(state) || i < 0 {
				continue
			}
			if j >= len(state[i]) || j < 0 {
				continue
			}
			if j == col && i == row {
				continue
			}
			if state[i][j] == '#' {
				adjCount++
			}
		}
	}
	return adjCount
}

func visiblyOccupiedCount(row int, col int, rowIncrement int, colIncrement int, state [][]rune) int {
	i := row
	j := col
	for {
		i = i + rowIncrement
		j = j + colIncrement
		if i < 0 || i == len(state) {
			return 0
		}
		if j < 0 || j == len(state[row]) {
			return 0
		}
		if state[i][j] == '#' {
			return 1
		}
		if state[i][j] == 'L' {
			return 0
		}
	}
}

func getVisibleOccupiedSeats(row int, col int, state [][]rune) int {
	adjCount := 0
	adjCount += visiblyOccupiedCount(row, col, -1, 0, state)
	adjCount += visiblyOccupiedCount(row, col, 1, 0, state)
	adjCount += visiblyOccupiedCount(row, col, 0, -1, state)
	adjCount += visiblyOccupiedCount(row, col, 0, 1, state)
	adjCount += visiblyOccupiedCount(row, col, -1, -1, state)
	adjCount += visiblyOccupiedCount(row, col, 1, -1, state)
	adjCount += visiblyOccupiedCount(row, col, -1, 1, state)
	adjCount += visiblyOccupiedCount(row, col, 1, 1, state)
	return adjCount
}

func getNumOccupiedSeats(initialState [][]rune, visibilityMethod bool, tolerance int) int {
	for {
		changed := false
		newState := make([][]rune, len(initialState))

		//printState(initialState)
		for row := 0; row < len(initialState); row++ {
			newState[row] = make([]rune, len(initialState[row]))
			copy(newState[row], initialState[row])
			for col := 0; col < len(initialState[row]); col++ {
				adjCount := 0
				if visibilityMethod {
					adjCount = getVisibleOccupiedSeats(row, col, initialState)
				} else {
					adjCount = getAdjacentOccupiedSeats(row, col, initialState)
				}
				if adjCount == 0 && initialState[row][col] == 'L' {
					newState[row][col] = '#'
					changed = true
				} else if adjCount >= tolerance && initialState[row][col] == '#' {
					newState[row][col] = 'L'
					changed = true
				}
			}
		}
		if !changed {
			break
		}

		for row := 0; row < len(initialState); row++ {
			copy(initialState[row], newState[row])
		}
	}

	occupiedCount := 0
	for row := 0; row < len(initialState); row++ {
		for col := 0; col < len(initialState[row]); col++ {
			if initialState[row][col] == '#' {
				occupiedCount++
			}
		}
	}

	return occupiedCount
}
