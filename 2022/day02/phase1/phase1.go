package main

import (
	"aoc/common/files"
	"aoc/common/operators"
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
	lines := files.Read("input.txt")

	for _, line := range lines {
		runes := []rune(line)
		l := int(runes[0]) - int('A')
		r := int(runes[2]) - int('X')
		ans += r + 1

		if l == r {
			ans += 3
		} else if (operators.PyMod((r - l), 3)) == 1 {
			ans += 6
		}
	}

	fmt.Print(ans)
}
