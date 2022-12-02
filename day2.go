package main

import (
	"bufio"
	"fmt"
	"strings"
)

var scoreMap = map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}

func calculateMyScore(opponentPlay string, myPlay string) int {

	baseScore := scoreMap[myPlay]

}

func Day_2() {
	fileIn, err := GetFileHandleByDay("2")

	if err != nil {
		panic("Input file expected")
	}
	defer fileIn.Close()

	// Read line by line
	s := bufio.NewScanner(fileIn)
	for s.Scan() {

		for i := 1; i <= 3; i++ {
			roundPlay := strings.Fields(s.Text())

			opponentPlay := roundPlay[0]
			myPlay := roundPlay[1]

			fmt.Println(opponentPlay, scoreMap[opponentPlay],
				myPlay, scoreMap[myPlay])

		}
	}

}
