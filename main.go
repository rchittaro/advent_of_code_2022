package main

import (
	"fmt"
	"os"
	"strconv"
)

// List of functions that solve a particular 'Day' problem
var funcMap = map[int]func(){
	1: Day_1,
	2: Day_2,
	3: Day_3,
	4: Day_4,
	5: Day_5,
	6: Day_6,
	7: Day_7,
	8: Day_8,
	9: Day_9,
}

func usage(progName string) {
	fmt.Println(progName, " <day>")
	fmt.Println("Example: ", progName, " 1")
}

func runSingleDay(dayi int) {

	dayFun, ok := funcMap[dayi]
	if !ok {
		fmt.Println("Challenge not ready for Day ", dayi)
		os.Exit(0)
	}

	dayFun()
}

// Main
func main() {

	// Check to see if the 'day' is provided on the command line. If so, we skip
	// the gui and and just run that one..and we will use 'reflection'
	if len(os.Args) == 2 {
		dayi, err := strconv.Atoi(os.Args[1])

		// Single arg needs to be an int
		if err != nil {
			usage(os.Args[0])
			os.Exit(1)
		}

		runSingleDay(dayi)
		os.Exit(1)
	}

	// Create the menu of challenge days that have solutions
	app := appInit()
	menuList := generateMenulist()

	if err := app.SetRoot(menuList, true).SetFocus(menuList).Run(); err != nil {
		panic(err)
	}
}
