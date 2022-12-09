package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type FileEntry struct {
	name string
	sz   uint
}

type DirEntry struct {
	name      string
	parent    *DirEntry
	FileEntry []*FileEntry
	subdirs   []*DirEntry
	sz        uint
}

// Our current directory
var currDir *DirEntry
var startDir *DirEntry

func findDirEntry(startDir *DirEntry, name string) *DirEntry {

	if startDir == nil {
		return nil
	}

	for _, dir := range startDir.subdirs {
		if dir.name == name {
			return dir
		}
	}

	return nil
}

func GetCurrentPath(startDir *DirEntry) string {

	if startDir == nil {
		return string("")
	}

	return GetCurrentPath(startDir.parent) + "/" + startDir.name
}

func ChangeDirectory(destDir string) {

	// Where are we changing to?
	fmt.Println("Current Dir: ", GetCurrentPath(currDir))

	// Assumption is the input will start with a 'cd'
	if currDir == nil {
		currDir = new(DirEntry)
		currDir.name = destDir
		startDir = currDir
		fmt.Println("New Current Dir: ", currDir)
		return
	}

	if destDir == ".." {
		currDir = currDir.parent

		if currDir == nil {
			panic("Parent not supposed to be nil")
		}
		fmt.Println(GetCurrentPath(currDir))
		return
	}

	// Not a new directory
	currDir = findDirEntry(currDir, destDir)
	if currDir == nil {
		panic("something really borked here!")
	}

	fmt.Println("New Current Dir: ", GetCurrentPath(currDir))
}

func IsCommand(directive []string) bool {
	return directive[0] == "$"
}

func handleDirEntry(dirName string) {
	fmt.Println("found directory:", dirName)

	// Make sure it doesn't already exist
	if findDirEntry(currDir, dirName) != nil {
		panic("aha!!")
	}

	newDir := new(DirEntry)
	newDir.name = dirName
	newDir.parent = currDir
	newDir.sz = 0
	currDir.subdirs = append(currDir.subdirs, newDir)
}

func handleFileEntry(sz int, name string) {
	// add it to the current directory file list

	newFile := new(FileEntry)
	newFile.name = name
	newFile.sz = uint(sz)
	currDir.FileEntry = append(currDir.FileEntry, newFile)
	currDir.sz += uint(sz)
}

func ListDirectory(s *bufio.Scanner) []string {
	// we have to read the input until we find a new command
	// prompt that signals the end of the 'ls' command

	fmt.Println(("---------------"))
	var directive []string
	for {
		directive = GetNextCmd(s)
		fmt.Println(directive)
		if len(directive) == 0 || IsCommand(directive) {
			fmt.Println(("-----cmd----------"))
			return directive
		}

		// Not a command, so assume it is output from the 'ls'
		if directive[0] == "dir" {
			handleDirEntry(directive[1])
		} else {
			if val, err := strconv.Atoi(directive[0]); err == nil {
				handleFileEntry(val, directive[1])
			} else {
				panic("bad input is not expected")
			}
		}
	}
}

func ProcessCommand(directive []string, s *bufio.Scanner) []string {

	fmt.Print("Processing: ", directive, " -->")
	if directive[0] == "$" {
		switch directive[1] {
		case "cd":
			ChangeDirectory(directive[2])
		case "ls":
			// Assumption is there are no arguments to 'ls' in the input
			if len(directive) > 2 {
				panic("Bad assumption")
			}

			return (ListDirectory(s))
		default:
			panic("Uknown command" + directive[1])
		}
	}

	return []string{}
}

func GetNextCmd(s *bufio.Scanner) []string {

	if !s.Scan() {
		return []string{}
	}
	return strings.Fields(s.Text())
}

func CalcDirSize(startDir *DirEntry, accumByMax uint, accumTotal *uint) uint {
	// Local files already calculated. Need to add up size
	// of each subdirectory
	var sz uint = 0

	if startDir == nil {
		return 0
	}

	sz = startDir.sz // files already calculated
	for _, dir := range startDir.subdirs {
		if dir.name == "fhg" {
			fmt.Println("here")
		}
		sz += CalcDirSize(dir, accumByMax, accumTotal)
	}

	startDir.sz = sz
	if sz <= accumByMax {
		*accumTotal += startDir.sz
	}

	fmt.Println(startDir.name, startDir.sz)

	return startDir.sz
}

func FindMinDir(startDir *DirEntry, minSz uint, res *uint) uint {

	if startDir == nil {
		return *res
	}

	for _, dir := range startDir.subdirs {
		sz := FindMinDir(dir, minSz, res)

		if sz > minSz && sz < *res {
			*res = sz
		}
	}

	if startDir.sz > minSz && startDir.sz < *res {
		*res = startDir.sz
	}

	return *res

}

func Day_7() {
	fileIn, err := GetFileHandleByDay("7")

	if err != nil {
		panic("Input file expected")
	}
	defer fileIn.Close()

	s := bufio.NewScanner(fileIn)
	directive := GetNextCmd(s)

	for len(directive) != 0 {
		directive = ProcessCommand(directive, s)

		if len(directive) == 0 {
			directive = GetNextCmd(s)
		}
	}

	var accum uint = 0

	totalUsed := CalcDirSize(startDir, 100000, &accum)
	fmt.Println("total size of 100k dirs: ", accum, totalUsed)

	// We have the total number used, we need to figure out how much to delete
	freeSpace := 70000000 - totalUsed
	needToFree := 30000000 - freeSpace
	var res uint = 7000000

	fmt.Println("Smallest dir to free: ", needToFree,
		FindMinDir(startDir, needToFree, &res))
}
