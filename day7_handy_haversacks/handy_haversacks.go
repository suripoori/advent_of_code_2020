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
	bagDatas := strings.Split(string(data), "\n")
	bagMaps := getBagMaps(bagDatas)
	fmt.Println(part1(bagMaps, "shiny gold"))
	fmt.Println(part2(bagMaps, "shiny gold"))
}

func part1(bagMaps map[string]map[string]int, bagColorOfInterest string) int {
	count := 0
	for bagColor := range bagMaps {
		if canBagHoldBag(bagMaps, bagColor, bagColorOfInterest) {
			count++
		}
	}
	return count
}

func canBagHoldBag(bagMaps map[string]map[string]int, bagInQuestion string, bagColorOfInterest string) bool {
	for bagColor := range bagMaps[bagInQuestion] {
		if bagColor == bagColorOfInterest {
			return true
		}
		if canBagHoldBag(bagMaps, bagColor, bagColorOfInterest) {
			return true
		}
	}
	return false
}

func part2(bagMaps map[string]map[string]int, bagColorOfInterest string) int {
	totalBagCount := getTotalBagCount(bagMaps, bagColorOfInterest)
	// Since we only want to know how many total bags must be present inside the 1 shiny gold bag,
	// we subtract 1 from our total here.
	return totalBagCount - 1
}

func getTotalBagCount(bagMaps map[string]map[string]int, bagColorOfInterest string) int {
	totalCount := 1
	for bagColor, count := range bagMaps[bagColorOfInterest] {
		totalCount += count * getTotalBagCount(bagMaps, bagColor)
	}
	return totalCount
}

func getBagMaps(bagDatas []string) map[string]map[string]int {
	var bagMaps = map[string]map[string]int{}

	for _, bagData := range bagDatas {
		if len(bagData) == 0 {
			break
		}

		bagColor := strings.Trim(strings.Split(bagData, "bags")[0], " ")
		bagContents := strings.Trim(strings.Split(bagData, "contain")[1], " ")
		bagCounts := make(map[string]int)
		for _, bag := range strings.Split(bagContents, "bag") {
			bagContent := strings.Trim(bag, "s")
			bagContent = strings.Trim(bagContent, ",")
			bagContent = strings.Trim(bagContent, " ")
			bagContent = strings.Trim(bagContent, ".")
			if strings.Contains(bagContent, "no other") {
				break
			}
			if len(bagContent) == 0 {
				break
			}
			count, err := strconv.Atoi(strings.Split(bagContent, " ")[0])
			if err != nil {
				fmt.Println("Failed to convert integer count", bagContent)
				break
			}
			color := strings.Split(bagContent, " ")[1] + " " + strings.Split(bagContent, " ")[2]
			bagCounts[color] = count
		}
		bagMaps[bagColor] = bagCounts
	}
	return bagMaps
}
