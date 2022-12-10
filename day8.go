package main

import (
	"bufio"
	"fmt"
)

func treeHeight(x int, y int, treeGrid []string) byte {
	return treeGrid[x][y] - '0'
}

func isVisible(row int, col int, treeGrid []string) (bool, int) {

	// top and bottom are visible
	if row == 0 || row == len(treeGrid) {
		return true, 0
	}

	// left and right are visible
	if col == 0 || col == len(treeGrid[0]) {
		return true, 0
	}

	checkTreeHeight := treeHeight(row, col, treeGrid)
	var vleft, vright, vup, vdown bool = true, true, true, true
	var sleft, sright, sup, sdown int = 0, 0, 0, 0

	// check up
	for x := row - 1; x >= 0; x-- {
		sup++
		if treeHeight(x, col, treeGrid) >= checkTreeHeight {
			vup = false
			break
		}
	}

	// check down
	for x := row + 1; x < len(treeGrid); x++ {
		sdown++
		if treeHeight(x, col, treeGrid) >= checkTreeHeight {
			vdown = false
			break
		}
	}

	// check left
	for y := col - 1; y >= 0; y-- {
		sleft++
		if treeHeight(row, y, treeGrid) >= checkTreeHeight {
			vleft = false
			break
		}
	}

	// check right
	for y := col + 1; y < len(treeGrid[row]); y++ {
		sright++
		if treeHeight(row, y, treeGrid) >= checkTreeHeight {
			vright = false
			break
		}
	}

	return vleft || vright || vup || vdown, sup * sleft * sright * sdown
}

func countVisibleTrees(treeGrid []string) (int, int) {
	ct := 0
	var viewScore int = 0

	for x := 0; x < len(treeGrid[0]); x++ {
		for y := 0; y < len(treeGrid); y++ {
			visible, vs := isVisible(x, y, treeGrid)

			if visible {
				ct++
			}

			if vs > viewScore {
				viewScore = vs
			}
		}
	}
	return ct, viewScore
}

func Day_8() {
	fileIn, err := GetFileHandleByDay("8")

	if err != nil {
		panic("Input file expected")
	}
	defer fileIn.Close()

	var treeMatrix []string

	s := bufio.NewScanner(fileIn)
	i := 0
	for s.Scan() {
		treeMatrix = append(treeMatrix, s.Text())
		i += 1
	}

	visible, viewScore := countVisibleTrees(treeMatrix)
	fmt.Println("The number of visible trees is: ", visible, viewScore)
}
