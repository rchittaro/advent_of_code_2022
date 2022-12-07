package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	stack "github.com/ttdsuen/golang-stack"
)

var stackMap map[int]*stack.Stack[string]
var stackCt int

func allocStack() *stack.Stack[string] {
	return stack.NewStack[string]()
}

func getStack(x int) *stack.Stack[string] {
	var stk *stack.Stack[string] = stackMap[x]

	if stk == nil {
		panic("WTF")
	}

	return stk
}

func CheckStacks() {
	for i := 1; i <= stackCt; i++ {
		if stackMap[i] == nil {
			panic("not expected")
		}
	}
}
func InitStacks(n int) {

	stackMap[1] = allocStack()
	stackMap[2] = allocStack()
	stackMap[3] = allocStack()
	stackMap[4] = allocStack()
	stackMap[5] = allocStack()
	stackMap[6] = allocStack()
	stackMap[7] = allocStack()
	stackMap[8] = allocStack()
	stackMap[9] = allocStack()

}

func LoadStacks(s *bufio.Scanner) {
	// Assumption is we read until there is an empty line
	// Read all lines up to the empty line and push value onto a stack
	// we can then pop them off to store as the initial values representing the containers
	stk := stack.NewStack[string]()
	for s.Scan() {
		if len(s.Text()) == 0 {
			fmt.Println("Found empty line")
			break
		} else {
			stk.Push(s.Text())
		}
	}

	// Last line read will be all of the stacks. Assumption
	// is they are
	var val string
	var ok bool

	val, _ = stk.Pop()
	stackNames := strings.Fields(val)
	stackCt, _ = strconv.Atoi(stackNames[len(stackNames)-1])
	fmt.Println("number of stacks is: ", stackCt)

	InitStacks(stackCt)
	CheckStacks()
	fmt.Println(stackMap[1] == stackMap[9])

	for {
		val, ok = stk.Pop()

		if !ok {
			break
		}

		// Each input line uses spaces to separate each entry, and uses spaces to represent where there is 'no entry'
		// so standard string separation won't work. In order to make it work, let's replace every 3rd character with a comma
		var targetStk int = 0
		for i := 0; i < len(val); i += 4 {
			targetStk++
			entry := val[i : i+3]
			entry = strings.TrimSpace(entry)

			if len(entry) == 3 {
				stkEntry := getStack(targetStk)
				stkEntry.Push(entry)
			}
		}

	}

	PrintStacks()
}

func PrintStacks() {
	for i := 1; i <= stackCt; i++ {
		stk := getStack(i)

		fmt.Printf("stk: %v\n", stk)
	}
}

func MoveContainer(count int, fromStack int, toStack int, multiMove bool) {

	fmt.Println(count, fromStack, toStack)

	if count == 0 {
		return
	}

	fromStk := getStack(fromStack)
	toStk := getStack(toStack)

	if fromStk == toStk {
		panic("seriously")
	}

	if multiMove {
		tmpStk := stack.NewStack[string]()
		for i := 1; i <= count; i++ {
			entry, ok := fromStk.Pop()
			if ok {
				tmpStk.Push(entry)
			} else {
				break
			}
		}

		for {
			entry, ok := tmpStk.Pop()
			if ok {
				toStk.Push(entry)
			} else {
				break
			}
		}

	} else {
		for i := 1; i <= count; i++ {

			entry, ok := fromStk.Pop()
			if ok {
				toStk.Push(entry)
			} else {
				fmt.Println()
			}
		}
	}

}

func Day_5() {
	fileIn, err := GetFileHandleByDay("5")

	if err != nil {
		panic("Input file expected")
	}
	defer fileIn.Close()

	s := bufio.NewScanner(fileIn)

	stackMap = make(map[int]*stack.Stack[string])
	LoadStacks(s)

	for s.Scan() {
		//	var tmpm, tmpf, tmpt string
		var ct, fromStk, toStk int

		directive := strings.Fields(s.Text())

		ct, _ = strconv.Atoi(directive[1])
		fromStk, _ = strconv.Atoi(directive[3])
		toStk, _ = strconv.Atoi(directive[5])
		MoveContainer(ct, fromStk, toStk, true)
		PrintStacks()
	}

	fmt.Println("---------------")
	PrintStacks()

}
