package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	fmt.Println("PART ONE: ", checkWrongItem(bufio.NewScanner(file)))
	_, _ = file.Seek(0, io.SeekStart)
	fmt.Println("PART TWO: ", checkIdBadges(bufio.NewScanner(file)))
}

// (I) part one
func checkWrongItem(scanner *bufio.Scanner) int {
	var sum int

	for scanner.Scan() {
		line := scanner.Text()
		mid := len(line) / 2
		re := regexp.MustCompile("[" + line[:mid] + "]")
		wrongItem := re.Find([]byte(line[mid:]))[0]
		sum += IF(wrongItem > 96, int(wrongItem)-96, int(wrongItem)-38) // ascii => priority
	}
	return sum
}

// (II) part two
func checkIdBadges(scanner *bufio.Scanner) int {
	var sum int
	var group []string
	var wrongItem byte

	for scanner.Scan() {
		group = append(group, scanner.Text())

		if len(group) > 2 {
			for i := 0; i < len(group[0]); i++ {
				wrongItem = group[0][i]
				if strings.Contains(group[1], string(wrongItem)) &&
					strings.Contains(group[2], string(wrongItem)) {
					break
				}
			}
			sum += IF(wrongItem > 96, int(wrongItem)-96, int(wrongItem)-38) // ascii => priority
			group = nil
		}
	}
	return sum
}

// utils : ternary
func IF[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}
