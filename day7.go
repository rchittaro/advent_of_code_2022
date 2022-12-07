package main

import (
	"bufio"
	"fmt"
	"strings"
)

type FileSystemNode struct {
	name      string
	parent    string
	files     []string
	subdirs   []string
	totalSize int
}

// Use a map to track directory sizes

// Our current directory
var currDir string

func NavigateDir(toDir string) bool {

}

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
