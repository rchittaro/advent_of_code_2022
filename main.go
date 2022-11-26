package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("Day 1", "Some explanatory text", 'a', nil).
		AddItem("Day 2", "Some explanatory text", 'b', nil).
		AddItem("Day 3", "Some explanatory text", 'c', nil).
		AddItem("Day 4", "Some explanatory text", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
