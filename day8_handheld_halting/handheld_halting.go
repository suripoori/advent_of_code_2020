package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
        return
	}
	instructions := strings.Split(string(data), "\n")
	instructionList := getInstructionList(instructions)
	fmt.Println(part1(instructionList))
	fmt.Println(part2(instructionList))
}

type Instruction struct {
	index int
	operation string
	argument int
}

func part1(instructionList []Instruction) (int) {
	acc, _ := getAccValue(0, instructionList)
	return acc
}

func part2(instructionList []Instruction) (int) {
	for index, instruction := range(instructionList) {
		if instruction.operation == "jmp" {
			newInstructionList := make([]Instruction, len(instructionList))
			copy(newInstructionList, instructionList)
			newInstructionList[index].operation = "nop"
			acc, ip := getAccValue(0, newInstructionList)
			if ip >= len(newInstructionList) {
				return acc
			}
		} else if instruction.operation == "nop" {
			newInstructionList := make([]Instruction, len(instructionList))
			copy(newInstructionList, instructionList)
			newInstructionList[index].operation = "jmp"
			acc, ip := getAccValue(0, newInstructionList)
			if ip >= len(newInstructionList) {
				return acc
			}
		}
	}
	return 0
}

func getAccValue(acc int, instructionList []Instruction) (int, int) {
	visited := make(map[int]bool)
	ip := 0
	for {
		if ip >= len(instructionList) {
			break
		}
		instruction := instructionList[ip]
		_, present := visited[instruction.index]
		if present {
			return acc, ip
		}
		visited[instruction.index] = true

		switch instruction.operation {
		case "nop":
			ip++
		case "acc":
			acc += instruction.argument
			ip++
		case "jmp":
			ip += instruction.argument
		} 
	}
	return acc, ip
}

func getInstructionList(instructions []string) ([]Instruction) {
	var output = []Instruction{}

	for index, instruction := range(instructions) {
		if len(instruction) == 0 {
			continue
		}
		operation := strings.Split(instruction, " ")[0]
		argument := strings.Split(instruction, " ")[1]
		argVal, err := strconv.Atoi(argument[1:])
		if err != nil {
			fmt.Println("Unable to convert instruction argument", argument)
			break
		}
		if argument[0] == '-' {
			argVal *= -1
		}
		instructionObj := Instruction{index: index, operation: operation, argument: argVal}
		output = append(output, instructionObj)
	}
	return output
}