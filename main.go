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

var funcMap = map[int]func(){
	1: Day_1,
	2: Day_2,
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

		dayFun, ok := funcMap[dayi]
		if !ok {
			fmt.Println("Challenge not completed for Day ", dayi)
			os.Exit(0)
		}

		dayFun()
		os.Exit(1)
	}

	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("Day 1", "", 'a', Day_1).
		AddItem("Day 2", "", 'b', nil).
		AddItem("Day 3", "", 'c', nil).
		AddItem("Day 4", "", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
