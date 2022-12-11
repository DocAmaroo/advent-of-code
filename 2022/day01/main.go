package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	file := utils.ReadFile("input.txt")
	lines := utils.GetLines(file)

	var calories []int
	total := 0
	for _, s := range lines {
		if s == "" {
			calories = append(calories, total)
			total = 0
		} else {
			val := utils.Atoi(s)
			total += val
		}
	}

	fmt.Println("Maximum calories:", utils.Max(calories))
	utils.ReverseSort(calories)
	fmt.Println("Total top 3:", calories[0]+calories[1]+calories[2])
}
