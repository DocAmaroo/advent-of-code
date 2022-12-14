package main

import (
	"aoc/common/files"
	"aoc/common/maps"
	op "aoc/common/operators"
	"fmt"
	"sort"
)

func main() {
	lines := files.ReadAndSplit("input.txt", " ")
	ans := 0
	pwd := make([]string, 0)
	sizes := make(map[string]int)

	for _, l := range lines {
		if l[0] == "dir" || l[1] == "ls" {
			continue
		} else if l[1] == "cd" && l[2] == "/" {
			pwd = []string{"/"} // remove all elements
		} else if l[1] == "cd" && l[2] == ".." {
			pwd = pwd[:len(pwd)-1]
		} else if l[1] == "cd" {
			pwd = append(pwd, l[2])
		} else {
			resolve := ""
			for _, p := range pwd[:] {
				resolve += p
				sizes[resolve] += op.Atoi(l[0])
			}
		}
	}

	curr := sizes["/"]
	needed := 70000000 - 30000000
	values := maps.GetValues(sizes)
	sort.Ints(values)
	for _, v := range values {
		if v >= curr-needed {
			ans = v
			break
		}
	}

	fmt.Println(ans)
}
