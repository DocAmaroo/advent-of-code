package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strings"
)

func readInstruction(s string) (m int, f int, t int) {
	ins := utils.SliceAtoi(regexp.MustCompile(`\d+`).FindAllString(s, -1))
	return ins[0], ins[1], ins[2]
}

func main() {
	ans := ""
	lines := utils.Read("input.txt")

	sepIdx := 0
	for lines[sepIdx] != "" {
		sepIdx++
	}
	// parse line: 1 2 3 ...
	n := regexp.MustCompile(`\d`).FindAllString(lines[sepIdx-1], -1)
	stacks := make([]string, len(n))

	// build stacks
	for _, s := range lines[0 : sepIdx-1] {
		rep := strings.ReplaceAll(s, "    ", ".")
		rep = regexp.MustCompile(`\s?\[?\]?`).ReplaceAllString(rep, "")
		for i, e := range rep {
			if e != '.' {
				stacks[i] += string(e)
			}
		}
	}

	// read and apply instructions
	for _, s := range lines[sepIdx+1:] {
		mv, fm, to := readInstruction(s)
		stacks[to-1] = stacks[fm-1][:mv] + stacks[to-1]
		stacks[fm-1] = stacks[fm-1][mv:]
	}

	// get first elements of each stack
	for _, s := range stacks {
		ans += string(s[0])
	}
	fmt.Println(ans)
}
