package main

import (
	"aoc/common/arrays"
	"aoc/common/files"
	"fmt"
	"strconv"
)

func main() {
	lines := files.Read("day01/input.txt")

	var calories []int
	total := 0
	for _, s := range lines {
		if s == "" {
			calories = append(calories, total)
			total = 0
		} else {
			val, _ := strconv.Atoi(s)
			total += val
		}
	}

	fmt.Println("p1", arrays.Max(calories))
	arrays.ReverseSort(calories)
	fmt.Println("p2", calories[0]+calories[1]+calories[2])
}
