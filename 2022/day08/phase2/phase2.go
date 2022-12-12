package main

import (
	"aoc/common/arrays"
	"aoc/common/files"
	"fmt"
)

func visible(m [][]int, x int, y int) int {
	curr := m[x][y]
	r, l, t, b := 1, 1, 1, 1

	// left
	for j := y - 1; j > 0; j-- {
		if m[x][j] >= curr {
			break
		}
		l++
	}

	// right
	for j := y + 1; j < len(m[0])-1; j++ {
		if m[x][j] >= curr {
			break
		}
		r++
	}

	// top
	for i := x - 1; i > 0; i-- {
		if m[i][y] >= curr {
			break
		}
		t++
	}

	// bottom
	for i := x + 1; i < len(m[0])-1; i++ {
		if m[i][y] >= curr {
			break
		}
		b++
	}

	return r * l * t * b
}

func main() {
	m := files.ReadAsMatrix("input.txt")
	n := len(m[0])

	values := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				continue
			} else {
				values = append(values, visible(m, i, j))
			}
		}
	}

	// fmt.Println(values)
	fmt.Println(arrays.Max(values))
}
