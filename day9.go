package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

var headPos Coordinate
var tailPos Coordinate
var tailVisits []Coordinate

func Day_9() {
	fileIn, err := GetFileHandleByDay("9")

	if err != nil {
		panic("Input file expected")
	}
	defer fileIn.Close()

	s := bufio.NewScanner(fileIn)
	for s.Scan() {
		moveMent := strings.Fields(s.Text())
		fmt.Println(moveMent)
	}

}
