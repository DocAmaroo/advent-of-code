package main

import (
	"aoc/common/files"
	"fmt"
	"strings"
)

func main() {
	ans := 0
	lines := files.Read("input.txt")
	group := make([]string, 3)
	for i, line := range lines {
		idx := (i + 1) % 3

		group[idx] = line
		if idx == 0 {
			for _, r := range group[0] {
				if strings.ContainsRune(group[1], r) && strings.ContainsRune(group[2], r) {
					if r > 96 {
						ans += int(r-'a') + 1
					} else {
						ans += int(r-'A') + 27
					}
					break
				}
			}
		}
	}

	fmt.Println(ans)
}
