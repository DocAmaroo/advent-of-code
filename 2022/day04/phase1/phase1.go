package main

import (
	"aoc/common/files"
	op "aoc/common/operators"
	"fmt"
	"strings"
)

func split(l string) (a int, b int, c int, d int) {
	spl := strings.Split(strings.ReplaceAll(l, ",", "-"), "-")
	return op.Atoi(spl[0]), op.Atoi(spl[1]), op.Atoi(spl[2]), op.Atoi(spl[3])
}

func main() {
	lines := files.Read("input.txt")

	ans := 0
	for _, l := range lines {
		a, b, x, y := split(l)
		aInb := a <= x && b >= y
		bIna := x <= a && y >= b
		ans += op.BoolToInt(aInb || bIna)
	}

	fmt.Println(ans)
}
