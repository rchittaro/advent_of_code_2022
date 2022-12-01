package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
)

type ElfNode struct {
	next     *ElfNode
	elf      int
	calories int
}

type ElfList struct {
	head *ElfNode
}

func (elfList *ElfList) Insert(newElfEntry *ElfNode) {

	if newElfEntry == nil {
		panic("should not be nil")
	}

	if elfList.head == nil {
		elfList.head = newElfEntry
		return
	} else if newElfEntry.calories > elfList.head.calories {
		tmp := elfList.head
		elfList.head = newElfEntry
		newElfEntry.next = tmp
		return
	}

	currentNode := elfList.head
	var previousNode *ElfNode = nil

	for currentNode != nil {
		if newElfEntry.calories > currentNode.calories {
			newElfEntry.next = currentNode
			previousNode.next = newElfEntry
			return
		}

		previousNode = currentNode
		currentNode = currentNode.next
	}

	// Insert at the end if we get here.
	previousNode.next = newElfEntry
}

func day_1_part1() (int, int, int) {

	fileIn, err := GetFileHandleByDay("1")
	if err != nil {
		log.Fatal(err)
		panic("Input file expected")
	}
	defer fileIn.Close()

	var elfList ElfList
	var currentElf *ElfNode = nil
	elveCt := 0

	// Read line by line
	s := bufio.NewScanner(fileIn)
	for s.Scan() {

		// Do we need to start a new Elf node
		if currentElf == nil {

			elveCt++
			currentElf = new(ElfNode)
			currentElf.elf = elveCt
		}

		if len(s.Text()) == 0 {
			elfList.Insert(currentElf)
			elveCt++
			currentElf = nil
		} else {
			calorieEntry, err := strconv.Atoi(s.Text())

			// There should not be any bad input. Implies data file is corrupt
			if err != nil {
				panic("Bad input value: " + s.Text())
			}

			currentElf.calories += calorieEntry
		}
	}

	top3Total := elfList.head.calories + elfList.head.next.calories + elfList.head.next.next.calories
	return elfList.head.elf, elfList.head.calories, top3Total
}

func Day_1() {
	elfKey, elfCalCarried, top3 := day_1_part1()

	fmt.Println("key   Calories   top3")
	fmt.Println(elfKey, elfCalCarried, top3)
}
