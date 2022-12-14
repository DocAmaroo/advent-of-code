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
	path := map[image.Point]struct{}{}
	rope := make([]image.Point, 10)

	for _, in := range input {
		var d rune
		var s int
		fmt.Sscanf(in, "%c %d", &d, &s)

		for i := 0; i < s; i++ {
			rope[0] = rope[0].Add(dirs[d])
			for i := 1; i < len(rope); i++ {
				if d := rope[i-1].Sub(rope[i]); op.Abs(d.X) > 1 || op.Abs(d.Y) > 1 {
					rope[i] = rope[i].Add(image.Point{sign(d.X), sign(d.Y)})
				}
			}
			path[rope[len(rope)-1]] = struct{}{}
		}

	}

	fmt.Println(len(path))
}
