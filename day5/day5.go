package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	result := moveIt(bufio.NewScanner(file))
	fmt.Println("PART 1: ", result)
}

func moveIt(scanner *bufio.Scanner, keepOrder bool) string {
	var stacks [][]string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.IndexAny(line, "move") != -1 {
			commands := strings.Split(line, " ")
			mCount, _ := strconv.Atoi(commands[1])
			mFrom, _ := strconv.Atoi(commands[3])
			mTo, _ := strconv.Atoi(commands[5])
			mFrom -= 1
			mTo -= 1

			toMove := make([]string, len(stacks[mFrom][:mCount]))
			_ = copy(toMove, stacks[mFrom][:mCount])

			if keepOrder {
				stacks[mTo] = append(toMove, stacks[mTo]...)
			} else {
				for _, item := range toMove {
					stacks[mTo] = append([]string{item}, stacks[mTo]...)
				}
			}
			stacks[mFrom] = stacks[mFrom][mCount:]
		} else {
			for inputIndex, character := range line {
				if character == '[' {
					stackIdx := int(inputIndex / 4)
					for len(stacks) <= stackIdx {
						var emptyStack []string
						stacks = append(stacks, emptyStack)
					}
					newValue := string(line[inputIndex+1])
					stacks[stackIdx] = append(stacks[stackIdx], newValue)
				}
			}
		}
	}
	var result string
	for _, stack := range stacks {
		fmt.Println(stack)
		result = result + stack[0]
	}
	return result
}
