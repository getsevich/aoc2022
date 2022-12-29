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
	var sum int = 0
	var cycle int = 0
	var step int = 40
	var signal int = 1
	var pendingSignal int = 0
	var exitNextCycle bool = false

	line := ""

	for true {
		nextOp := false

		//fmt.Print("", cycle, " ")

		if cycle > 0 {
			if int(math.Abs(float64(cycle-1-signal))) <= 1 {
				fmt.Print("#")
				//fmt.Print("[", signal, "]")
			} else {
				//fmt.Print("(", signal, ")")
				fmt.Print(".")
			}
		}

		if step == cycle {
			sum += signal * step
			//fmt.Println(" >>> ", step, "*", signal, " = ", signal*step, " | sum ", sum)
			fmt.Println()
			cycle = 0
			//step += 40
		} else {
			//fmt.Println()
		}

		if pendingSignal != 0 {
			//fmt.Print("A2 ", signal, "+", pendingSignal, " =", (signal + pendingSignal))
			signal += pendingSignal
			pendingSignal = 0
			nextOp = true
		} else if strings.HasPrefix(line, "addx") {
			value, _ := strconv.Atoi(strings.Split(line, " ")[1])
			pendingSignal = value
			//fmt.Print("A1 *", pendingSignal)
		} else {
			//fmt.Print("NOOP")
			nextOp = true
		}

		if exitNextCycle {
			return sum
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

	return sum
}
