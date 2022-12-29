package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	result := doStuff(bufio.NewScanner(file))
	fmt.Println("RES: ", result)
}

func doStuff(scanner *bufio.Scanner) int {
	var cycle int = 0
	var step int = 40
	var signal int = 1
	var pendingSignal int = 0
	var exitNextCycle bool = false

	line := ""

	for true {
		nextOp := false

		if cycle > 0 {
			if int(math.Abs(float64(cycle-1-signal))) <= 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		if step == cycle {
			fmt.Println()
			cycle = 0
		}

		if pendingSignal != 0 {
			signal += pendingSignal
			pendingSignal = 0
			nextOp = true
		} else if strings.HasPrefix(line, "addx") {
			value, _ := strconv.Atoi(strings.Split(line, " ")[1])
			pendingSignal = value
		} else {
			nextOp = true
		}

		if exitNextCycle {
			return 0
		}

		if nextOp {
			if !scanner.Scan() {
				fmt.Println("EoF")
				exitNextCycle = true
			} else {
				line = scanner.Text()
			}
		}

		cycle += 1

	}

	return 0
}
