package main

import (
	"bufio"
	"fmt"
	"strings"
)

// (1 for Rock, 2 for Paper, and 3 for Scissors)

var scoreMap = map[string]int{
	"A": 1, // Rock
	"B": 2, // Paper
	"C": 3, // Scissors
	"X": 1, // Rock
	"Y": 2, // Paper
	"Z": 3, // Scissors
}

type GameResult int

const (
	WIN  GameResult = 6
	LOSE            = 0
	DRAW            = 3
)

type DecisionNode struct {
	beats *DecisionNode
	play  int
}

var decisionTree *DecisionNode

func InitDecisionEngine() {

	decisionTree = new(DecisionNode)
	scissorNode := new(DecisionNode)
	paperNode := new(DecisionNode)

	decisionTree.play = 1
	scissorNode.play = 3
	paperNode.play = 2

	decisionTree.beats = scissorNode // rock beats scissors
	scissorNode.beats = paperNode    // scissors beats paper
	paperNode.beats = decisionTree   // paper beats rock

}

func GetDecisionNode(forValue int) *DecisionNode {

	if decisionTree.play == forValue {
		return decisionTree
	}

	currentNode := decisionTree.beats
	for {
		if currentNode == decisionTree {
			return nil
		}

		if currentNode.play == forValue {
			return currentNode
		}

		currentNode = currentNode.beats
	}
	return nil
}

func MyGameResult(opponentPlay string, myPlay string) GameResult {

	myScoreVal := scoreMap[myPlay]
	oppScoreVal := scoreMap[opponentPlay]

	// Check for draw
	if myScoreVal == oppScoreVal {
		return DRAW
	}

	// check the tree
	oppNode := GetDecisionNode(oppScoreVal)
	if oppNode == nil {
		panic("corrupted input or decision tree")
	}

	if oppNode.beats.play == myScoreVal {
		return LOSE
	}

	return WIN
}

func adjustMyPlay(opponentPlay string, myPlay string) (string, string) {

	// Draw, change my play to the same as opponent
	if myPlay == "Y" {
		switch opponentPlay {
		case "A": // Rock
			return opponentPlay, "X"
		case "B": // Paper
			return opponentPlay, "Y"
		case "C": // Scissors
			return opponentPlay, "Z"

		default:
			panic("bad input")
		}
	}

	if myPlay == "X" {
		// I need to lose
		switch opponentPlay {
		case "A": // Rock
			return opponentPlay, "Z"
		case "B": // Paper
			return opponentPlay, "X"
		case "C": // Scissors
			return opponentPlay, "Y"

		default:
			panic("bad input")
		}

	} else if myPlay == "Z" {
		// I need to win
		switch opponentPlay {
		case "A": // Rock
			return opponentPlay, "Y"
		case "B": // Paper
			return opponentPlay, "Z"
		case "C": // Scissors
			return opponentPlay, "X"

		default:
			panic("bad input")
		}

	} else {
		panic("did not finish")
	}

}

func calculateMyScore2(opponentPlay string, myPlay string) int {

	// automatically get score for playing
	myScore := scoreMap[myPlay]

	myScore += int(MyGameResult(opponentPlay, myPlay))
	return myScore
}

func Day_2() {
	fileIn, err := GetFileHandleByDay("2")

	if err != nil {
		panic("Input file expected")
	}
	defer fileIn.Close()

	InitDecisionEngine()
	myScorePart1 := 0
	myScorePart2 := 0

	// Read line by line
	s := bufio.NewScanner(fileIn)
	for s.Scan() {

		roundPlay := strings.Fields(s.Text())
		opponentPlay := roundPlay[0]
		myPlay := roundPlay[1]
		myScorePart1 += calculateMyScore2(opponentPlay, myPlay)

		opponentPlay, myPlay = adjustMyPlay(opponentPlay, myPlay)
		myScorePart2 += calculateMyScore2(opponentPlay, myPlay)
	}

	fmt.Println("Total Score: ", myScorePart1, myScorePart2)

}
