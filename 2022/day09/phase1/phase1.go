package main

import (
	"aoc/common/files"
	op "aoc/common/operators"
	"fmt"
	"image"
)

var dirs = map[rune]image.Point{
	'U': {0, -1},
	'R': {1, 0},
	'D': {0, 1},
	'L': {-1, 0},
}

func sign(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

func main() {
	input := files.Read("input.txt")
	path := make(map[image.Point]int)
	H := image.Point{0, 0}
	T := image.Point{0, 0}
	for _, in := range input {
		var d rune
		var s int
		fmt.Sscanf(in, "%c %d", &d, &s)

		for i := 0; i < s; i++ {

			H = H.Add(dirs[d])
			if d := H.Sub(T); op.Abs(d.X) > 1 || op.Abs(d.Y) > 1 {
				T = T.Add(image.Point{sign(d.X), sign(d.Y)})
			}

			path[T]++
		}

	}

	ans := 0
	for _, v := range path {
		if v > 1 {
			ans++
		}
	}

	fmt.Println(len(path))
}
