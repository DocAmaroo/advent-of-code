package main

import (
	"aoc/common/files"
	"fmt"
	"strings"
)

func checkCycle(pos *int, cpuIdx *int) {
	if *pos%40 == 0 && *pos != 0 {
		*cpuIdx++
		*pos = 0
	}
}

func step(cycle *[]string, pos *int, X *int, firstRun *bool) {
	if *pos >= *X-1 && *pos <= *X+1 {
		*cycle = append(*cycle, "#")
	} else {
		*cycle = append(*cycle, ".")
	}

	*pos++
	if *firstRun {
		*X++
		*firstRun = false
	}
}

func main() {
	input := files.Read("input.txt")

	firstRun := true
	cpu := make([][]string, 6)
	cpuIdx := 0
	X := 0
	pos := 0
	for _, in := range input {
		var sig string
		var val int
		fmt.Sscanf(in, "%s %d", &sig, &val)

		checkCycle(&pos, &cpuIdx)

		if sig == "addx" {
			step(&cpu[cpuIdx], &pos, &X, &firstRun)
		}

		checkCycle(&pos, &cpuIdx)
		step(&cpu[cpuIdx], &pos, &X, &firstRun)

		X += val
	}

	for _, c := range cpu {
		fmt.Println(strings.Join(c, ""))
	}
}
