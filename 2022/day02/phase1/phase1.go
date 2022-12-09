package main

import (
	"aoc/utils"
	"fmt"
)

/*
 * Rock     : A || X -> 1pt
 * Paper    : B || Y -> 2pt
 * Scissors : C || Z -> 3pt
 * -------------------
 * Lose -> 0pt
 * Draw -> 3pt
 * Win -> 6pt
 */

func main() {
	ans := 0
	file := utils.ReadFile("../input.txt")
	lines := utils.ParseByLines(file)

	for _, line := range lines {
		runes := []rune(line)
		l := int(runes[0]) - int('A')
		r := int(runes[2]) - int('X')
		ans += r + 1

		if l == r {
			ans += 3
		} else if (utils.PyMod((r - l), 3)) == 1 {
			ans += 6
		}
	}

	fmt.Print(ans)
}
