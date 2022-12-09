package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func split(l string) (a int, b int, c int, d int) {
	spl := strings.Split(strings.ReplaceAll(l, ",", "-"), "-")
	return utils.Atoi(spl[0]), utils.Atoi(spl[1]), utils.Atoi(spl[2]), utils.Atoi(spl[3])
}

func main() {
	lines := utils.Read("input.txt")

	ans := 0
	for _, l := range lines {
		a, b, x, y := split(l)
		aInb := a >= x && a <= y || b >= x && b <= y
		bIna := x <= a && x >= b || y >= a && y <= b
		ans += utils.BoolToint(aInb || bIna)
	}

	fmt.Println(ans)
}
