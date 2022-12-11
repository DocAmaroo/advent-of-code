package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	ans := 0
	file := utils.ReadFile("../input.txt")
	lines := utils.GetLines(file)

	for _, line := range lines {
		runes := []rune(line)
		l := int(runes[0]) - int('A')
		r := runes[2]

		if r == 'X' {
			ans += utils.PyMod((l-1), 3) + 1
		} else if r == 'Y' {
			ans += 3 + l + 1
		} else {
			ans += 6 + utils.PyMod((l+1), 3) + 1
		}
	}
	fmt.Print(ans)
}
