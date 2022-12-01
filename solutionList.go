package main

// List of functions that solve a particular 'Day' problem
var solutionMap = map[int]func(){
	1: Day_1,
	2: Day_2,
}

func getSolutionList() *map[int]func() {
	return &solutionMap
}
