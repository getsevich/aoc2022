package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maxValue := 0
	currentMaxValue := 0
	list := []int{}

	// 212836

	for scanner.Scan() {
		curLine := scanner.Text()
		if curLine == "" {
			list = append(list, currentMaxValue)
			if len(list) > 3 {
				sort.Sort(sort.Reverse(sort.IntSlice(list)))
				list = list[:3]
			}
			if currentMaxValue > maxValue {
				maxValue = currentMaxValue
			}
			currentMaxValue = 0
		} else {
			curValue, err := strconv.Atoi(curLine)
			if err == nil {
				currentMaxValue += curValue
			}
		}
	}

	result := 0
	for _, v := range list {
		result += v
	}

	fmt.Println(list)
	fmt.Println(result)
}
