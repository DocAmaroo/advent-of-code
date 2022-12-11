package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	ans := 0
	lines := utils.RemoveBlankSpaces(utils.Read("input.txt"))

	pwd := make([]string, 0)
	sizes := make(map[string]int)
	for _, l := range lines {
		if l[0] == "dir" || l[1] == "ls" {
			continue
		} else if l[1] == "cd" && l[2] == "/" {
			pwd = nil // remove all elements
		} else if l[1] == "cd" && l[2] == ".." {
			pwd = pwd[:len(pwd)-1]
		} else if l[1] == "cd" {
			pwd = append(pwd, l[2])
		} else {
			resolve := ""
			for _, p := range pwd[:] {
				resolve += p
				sizes[resolve] += utils.Atoi(l[0])
			}
		}
	}

	for _, s := range sizes {
		if s <= 100000 {
			ans += s
		}
	}

	fmt.Println(ans)
}
