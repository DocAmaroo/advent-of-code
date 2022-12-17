package main

import (
	"aoc/common/files"
	"fmt"
)

var STEPS = []int{20, 60, 100, 140, 180, 220}

func step(cycles map[int]int, i *int, X int, stepIdx *int) {
	*i++
	if *i == 20+(40**stepIdx) {
		cycles[*i] = X + 1
		*stepIdx++
	}
}

func main() {
	input := files.Read("input.txt")

	cycles := map[int]int{20: 0, 60: 0, 100: 0, 140: 0, 180: 0, 220: 0}
	X := 0
	i := 0
	stepIdx := 0

	for _, in := range input {
		var sig string
		var val int
		fmt.Sscanf(in, "%s %d", &sig, &val)

		if sig == "addx" {
			step(cycles, &i, X, &stepIdx)
		}

		step(cycles, &i, X, &stepIdx)

		X += val
	}

	ans := 0
	for _, s := range STEPS {
		ans += cycles[s] * s
	}
	fmt.Println(ans)
}
