package main

import (
	"bufio"
	"fmt"
	"strings"
)

func Day_7() {
	fileIn, err := GetFileHandleByDay("7")

	if err != nil {
		panic("Input file expected")
	}
	defer fileIn.Close()

	s := bufio.NewScanner(fileIn)
	for s.Scan() {
		directive := strings.Fields(s.Text())
		fmt.Println(directive)
	}
}
