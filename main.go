package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rivo/tview"
)

func usage(progName string) {
	fmt.Println(progName, " <day>")
	fmt.Println("Example: ", progName, " 1")
}

// List of functions that solve a particular 'Day' problem
var funcMap = map[int]func(){
	1: Day_1,
	2: Day_2,
}

func generateMenulist(app *tview.Application, menuList *tview.List) {
	var r rune = 'a'
	var entryStr string

	for key, funcVal := range funcMap {
		entryStr = "Day " + strconv.Itoa(key)
		menuList.AddItem(entryStr, "", r, funcVal)
		r++
	}

	menuList.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})
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

		if err != nil {
			usage(os.Args[0])
			os.Exit(1)
		}

		runSingleDay(dayi)
		os.Exit(1)
	}

	app := tview.NewApplication()
	menuList := tview.NewList()

	// Create the menu of challenge days that have solutions
	generateMenulist(app, menuList)

	if err := app.SetRoot(menuList, true).SetFocus(menuList).Run(); err != nil {
		panic(err)
	}
}
