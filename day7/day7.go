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
	result := mapStorage(bufio.NewScanner(file))
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

func mapStorage(scanner *bufio.Scanner) int {
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

	updateDirSize(&root)
	var max int = root.size
	getTargetDir(&root, &max, 30000000-(70000000-root.size))
	return max
}

func getTargetDir(cur *Dir, maxCounter *int, limit int) {
	for _, child := range cur.children {
		if child.size >= limit && child.size < *maxCounter {
			*maxCounter = child.size
		}
		getTargetDir(child, maxCounter, limit)
	}
}

func updateDirSize(cur *Dir) int {
	curSize := 0
	for _, child := range cur.children {
		curSize += updateDirSize(child)
	}
	cur.size += curSize
	return cur.size
}
