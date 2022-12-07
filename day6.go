package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

var message string
var msgStart int = 0

func incStart() {
	sigStart++
}

func incMsgStart() {
	msgStart++
}

func ProcessMessag(c rune, sz int) bool {
	message += string(c)

	if len(message) < sz {
		return false
	}

	defer incMsgStart()

	// check if the current sz character slice is unique
	m := make(map[rune]bool)

	for i := msgStart; i < msgStart+sz; i++ {
		_, ok := m[rune(message[i])]

		if ok {
			return false
		}

		m[rune(message[i])] = true
	}

	return true
}

var signal string
var sigStart int = 0

func ProcessSignal(c rune, sz int) bool {
	signal += string(c)

	if len(signal) < sz {
		return false
	}

	defer incStart()

	// check if the current 4 character slice is unique
	m := make(map[rune]bool)

	for i := sigStart; i < sigStart+sz; i++ {
		_, ok := m[rune(signal[i])]

		if ok {
			return false
		}

		m[rune(signal[i])] = true
	}

	return true
}

func Day_6() {
	fileIn, err := GetFileHandleByDay("6")

	if err != nil {
		panic("Input file expected")
	}
	defer fileIn.Close()

	charsRead := 0
	signalFound := false
	r := bufio.NewReader(fileIn)
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {

			if !signalFound {
				charsRead++
				fmt.Printf("%q [%d]\n", string(c), charsRead)
				if ProcessSignal(c, 4) {
					fmt.Println("Found Signal: ", charsRead)
					signalFound = true
				}
			} else {
				// Looking for messages now
				charsRead++
				fmt.Printf("Signal %q [%d]\n", string(c), charsRead)
				if ProcessMessag(c, 14) {
					fmt.Println("Found Message: ", charsRead)
					break
				}

			}
		}
	}

}
