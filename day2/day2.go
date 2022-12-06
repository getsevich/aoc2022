package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		curValues := strings.Split(scanner.Text(), " ")
		myChoice := curValues[1]
		enemyChoice := curValues[0]

		//fmt.Println(myChoice, enemyChoice)

		// A, X —  rocks — 1
		// B, Y —  paper — 2
		// C, Z —  scissors — 3

		// X lose  0
		// Y draw 3
		// Z win 6

		if enemyChoice == "A" && myChoice == "X" {
			total = total + 0 + 3
		}
		if enemyChoice == "B" && myChoice == "X" {
			total = total + 0 + 1
		}
		if enemyChoice == "C" && myChoice == "X" {
			total = total + 0 + 2
		}

		if enemyChoice == "A" && myChoice == "Y" {
			total = total + 3 + 1
		}
		if enemyChoice == "B" && myChoice == "Y" {
			total = total + 3 + 2
		}
		if enemyChoice == "C" && myChoice == "Y" {
			total = total + 3 + 3
		}

		if enemyChoice == "A" && myChoice == "Z" {
			total = total + 6 + 2
		}
		if enemyChoice == "B" && myChoice == "Z" {
			total = total + 6 + 3
		}
		if enemyChoice == "C" && myChoice == "Z" {
			total = total + 6 + 1
		}

	}

	fmt.Println(total)
}
