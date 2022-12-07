package main

import (
	"bufio"
	"fmt"
)

func ItemPriority(item byte) int {

	if item >= 'a' && item <= 'z' {
		return int(item) - 96

	} else if item >= 'A' && item <= 'Z' {
		return int(item) - 64 + 26
	}

	panic("Invalid input")
}

func FindNumberOccurrences(entry byte, target string) (int, string) {

	var result string
	var ct int

	for i := 0; i < len(target); i++ {
		if target[i] == entry {
			ct++
			result += string(entry)
		}
	}

	return ct, result
}

func commonItemPriority(commonItems string) int {

	if len(commonItems) == 0 {
		return 0
	}

	var ct int
	for i := 0; i < len(commonItems); i++ {
		ct += ItemPriority(commonItems[i])
	}

	return ct
}

func findCommonItems(comp1 string, comp2 string) string {
	var commonItems string

	for i := 0; i < len(comp1); i++ {
		ct, _ := FindNumberOccurrences(comp1[i], comp2)
		if ct > 0 {
			commonItems += string(comp1[i])
		}
	}

	return commonItems
}

var groupItems [3]string
var groupIdx int = 0
var groupPriority int = 0

func HandleElfGroups(inventory string) {

	groupItems[groupIdx] = inventory
	groupIdx++
	fmt.Println("added at idx: ", groupIdx, inventory)

	if groupIdx == 3 {
		c1c2Items := findCommonItems(groupItems[0], groupItems[1])
		allCommon := findCommonItems(groupItems[2], c1c2Items)
		fmt.Println("----------------------")
		fmt.Println(groupItems[0], groupItems[1], groupItems[2])
		fmt.Println("c1c2common: ", c1c2Items, "allcommon: ", allCommon)
		fmt.Println("----------------------")
		groupPriority += commonItemPriority(string(allCommon[0]))

		groupIdx = 0
		groupItems[0] = ""
		groupItems[1] = ""
		groupItems[2] = ""
	}

}

func Day_3() {
	fileIn, err := GetFileHandleByDay("3")

	if err != nil {
		panic("Input file expected")
	}
	defer fileIn.Close()

	// Read line by line
	var totalPriority int = 0
	groupIdx = 0

	s := bufio.NewScanner(fileIn)
	for s.Scan() {

		HandleElfGroups(s.Text())
		inLen := len(s.Text())
		comp1 := s.Text()[0 : inLen/2]
		comp2 := s.Text()[inLen/2:]
		commonItems := findCommonItems(comp1, comp2)

		sackPriority := commonItemPriority(string(commonItems[0]))
		//fmt.Println(comp1, comp2, commonItems, "Sack priority ", sackPriority)
		totalPriority += sackPriority
	}

	fmt.Println("Total Priority: ", totalPriority)
	fmt.Println("Group Priority: ", groupPriority)
}
