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
	pairs, ranges := checkRange(bufio.NewScanner(file))
	fmt.Println("PART 1: ", pairs)
	fmt.Println("PART 2: ", ranges)
}

func checkRange(scanner *bufio.Scanner) (int, int) {
	var overlapedPairs int
	var overlapedRangeCount int

	for scanner.Scan() {
		splitterFunc := func(c rune) bool { return c == ',' || c == '-' }
		strData := strings.FieldsFunc(scanner.Text(), splitterFunc)

		rng := make([]int, len(strData))
		for i, v := range strData {
			rng[i], _ = strconv.Atoi(v)
		}

		if (rng[0] >= rng[2] && rng[0] <= rng[3]) ||
			(rng[1] >= rng[2] && rng[0] <= rng[3]) {
			overlapedPairs += 1

			if()
		}
	}
	return overlapedPairs, overlapedRangeCount
}
