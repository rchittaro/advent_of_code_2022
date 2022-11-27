package main

import "os"

func GetFileHandleByDay(day string) (*os.File, error) {

	// expecting consistent input file location and names
	path := "data/day_" + day + ".txt"
	file, err := os.Open(path)

	return file, err
}
