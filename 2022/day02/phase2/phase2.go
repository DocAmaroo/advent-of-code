package main

import (
	"aoc/common/files"
	"aoc/common/operators"
	"fmt"
)

func main() {
	ans := 0
	lines := files.Read("input.txt")

	for _, line := range lines {
		runes := []rune(line)
		l := int(runes[0]) - int('A')
		r := runes[2]

		if r == 'X' {
			ans += operators.PyMod((l-1), 3) + 1
		} else if r == 'Y' {
			ans += 3 + l + 1
		} else {
			ans += 6 + operators.PyMod((l+1), 3) + 1
		}
	}
	fmt.Print(ans)
}
