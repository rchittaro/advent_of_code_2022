package main

import (
	"strconv"

	"github.com/rivo/tview"
)

var app *tview.Application
var menuList *tview.List

func appInit() *tview.Application {
	app = tview.NewApplication()
	menuList = tview.NewList()

	if app != nil && menuList != nil {
		return app
	}

	panic(0)
}

func generateMenulist() *tview.List {
	var r rune = 'a'
	var entryStr string

	// Setup the menu list
	menuList.SetBorder(true)
	menuList.SetTitle(" Advent of Code 2022 - Ron Chittaro ")
	//menuList.SetBorderColor(tcell.ColorGreen)

	for key, funcVal := range funcMap {
		entryStr = "Day " + strconv.Itoa(key)
		menuList.AddItem(entryStr, "", r, funcVal)
		r++
	}

	menuList.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})

	return menuList
}
