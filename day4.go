package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type SectionRange struct {
	start int
	end   int
}

func CreateRange(src string) *SectionRange {

	// need to split the string on the '-'
	startEnd := strings.Split(src, "-")
	rangeStruct := new(SectionRange)
	rangeStruct.start, _ = strconv.Atoi(startEnd[0])
	rangeStruct.end, _ = strconv.Atoi(startEnd[1])

	return rangeStruct
}

func RangesFullyContained(range1 *SectionRange, range2 *SectionRange) (bool, bool) {

	// Check for fully enclosed
	if range1.start >= range2.start && range1.end <= range2.end {
		// range1 is fully enclosed, also implies overlap
		return true, true
	} else if range2.start >= range1.start && range2.end <= range1.end {
		//range2 is fully enclosed, also implies overlap
		return true, true
	}

	// check for partial overlap
	if range1.start >= range2.start && range1.start <= range2.end {
		return false, true
	}

	if range2.start >= range1.start && range2.start <= range1.end {
		return false, true
	}

	return false, false
}

func Day_4() {
	fileIn, err := GetFileHandleByDay("4")

	if err != nil {
		panic("Input file expected")
	}
	defer fileIn.Close()

	var fullyEnclosedCt int = 0
	var partialEnclosedCt int = 0
	s := bufio.NewScanner(fileIn)
	for s.Scan() {
		sectRanges := strings.Split(s.Text(), ",")
		sect1Range := CreateRange(sectRanges[0])
		sect2Range := CreateRange(sectRanges[1])

		isEnclosed, partialEnclosed := RangesFullyContained(sect1Range, sect2Range)
		if isEnclosed {
			fullyEnclosedCt++
		}

		//partialEnclosed := RangesPartialContained(sect1Range, sect2Range)
		if partialEnclosed {
			partialEnclosedCt++
		}
	}

	fmt.Println("Total Enclosed: ", fullyEnclosedCt)
	fmt.Println("Partial Enclosed: ", partialEnclosedCt)
}
