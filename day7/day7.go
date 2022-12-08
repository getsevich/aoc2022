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
	result := checkStorage(bufio.NewScanner(file))
	fmt.Println("RES: ", result)
}

type Dir struct {
	parent   *Dir
	name     string
	size     int
	children []*Dir
}

func getDir(cur *Dir, name string) *Dir {
	for _, child := range cur.children {
		if child.name == name {
			return child
		}
	}
	return nil
}

func appendDir(cur *Dir, name string) *Dir {
	if getDir(cur, name) == nil {
		newDir := Dir{parent: cur, name: name, size: 0}
		cur.children = append(cur.children, &newDir)
		return &newDir
	}
	return nil
}

func checkStorage(scanner *bufio.Scanner) int {
	root := Dir{parent: nil, name: "root", size: 0}
	cur := &root

	for scanner.Scan() {

		line := scanner.Text()
		if strings.HasPrefix(line, "$ cd ") { // traversal
			cmd := strings.Split(line, " ")[2]
			if cmd == "/" {
				cur = &root
			} else if cmd == ".." {
				if cur.parent != nil {
					cur = cur.parent
				}
			} else {
				cur = getDir(cur, cmd)
			}
		} else if strings.HasPrefix(line, "dir ") { // dir discoverer
			appendDir(cur, strings.Split(line, " ")[1])
		} else if !strings.HasPrefix(line, "$") { // file discovered
			fileSize, _ := strconv.Atoi(strings.Split(line, " ")[0])
			cur.size += fileSize
		}
	}

	calcSize(&root)
	var max int = root.size + 500000000
	printDriveStruct_(&root, "", &max, 30000000-(70000000-root.size))
	return max
}

func printDriveStruct(cur *Dir, prefix string, maxCounter *int) { // just for fun
	for _, child := range cur.children {
		fmt.Println(prefix, child.name, child.size)
		newPrefix := prefix + "—"

		//if child.size <= 100000 {
		*maxCounter += child.size
		//}
		printDriveStruct(child, newPrefix, maxCounter)
	}
}

func printDriveStruct_(cur *Dir, prefix string, maxCounter *int, limit int) { // just for fun
	for _, child := range cur.children {
		fmt.Println(prefix, child.name, child.size)
		newPrefix := prefix + "—"

		if child.size >= limit && child.size < *maxCounter {
			*maxCounter = child.size
		}
		printDriveStruct(child, newPrefix, maxCounter)
	}
}

func calcSize(cur *Dir) int {
	curSize := 0
	for _, child := range cur.children {
		curSize += calcSize(child)
	}
	cur.size += curSize
	return cur.size
}
