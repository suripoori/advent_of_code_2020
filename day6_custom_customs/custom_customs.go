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
	groupDatas := strings.Split(string(data), "\n\n")
	fmt.Println(part1(groupDatas))
	fmt.Println(part2(groupDatas))
}

func groupCountAny(groupData string) (int) {
	answers := strings.Split(groupData, "\n")
	yesses := make(map[string]bool)

	for _, answer := range(answers) {
		for _, c := range(answer) {
			yesses[string(c)] = true
		}
	}
	return len(yesses)
}

func part1(groupDatas []string) (int) {
	count := 0
	for _, groupData := range(groupDatas) {
		count += groupCountAny(groupData)
	}
	return count
}

func intersect(m1 map[string]bool, m2 map[string]bool) (map[string]bool) {
	output := make(map[string]bool)
	for key, _ := range(m1) {
		_, present := m2[key]
		if present {
			output[key] = true
		}
	}
	for key, _ := range(m2) {
		_, present := m1[key]
		if present {
			output[key] = true
		}
	}
	return output
}

func groupCountAll(groupData string) (int) {
	answers := strings.Split(groupData, "\n")
	var allYesses map[string]bool

	for _, answer := range(answers) {
		if answer == string("") {
			continue
		}

		personYesses := make(map[string]bool)
		for _, c := range(answer) {
			personYesses[string(c)] = true
		}
		if allYesses == nil {
			allYesses = personYesses
		} else {
			allYesses = intersect(allYesses, personYesses)
		}
	}
	return len(allYesses)
}

func part2(groupDatas []string) (int) {
	count := 0
	for _, groupData := range(groupDatas) {
		count += groupCountAll(groupData)
	}
	return count
}