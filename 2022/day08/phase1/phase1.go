package main

import (
	"aoc/common/files"
	"fmt"
)

func visible(m [][]int, x int, y int) bool {
	curr := m[x][y]
	r, l, t, b := true, true, true, true

	// left
	for _, v := range m[x][:y] {
		if v >= curr {
			r = false
			break
		}
	}

	// right
	for _, v := range m[x][y+1:] {
		if v >= curr {
			l = false
			break
		}
	}

	// top
	for i := 0; i < x; i++ {
		if m[i][y] >= curr {
			t = false
			break
		}
	}

	// bottom
	for i := len(m[0]) - 1; i > x; i-- {
		if m[i][y] >= curr {
			b = false
			break
		}
	}

	return r || l || t || b
}

func main() {
	ans := 0
	m := files.ReadAsMatrix("input.txt")
	n := len(m[0])

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				ans++
				continue
			} else {
				if visible(m, i, j) {
					ans++
				}
			}
		}
	}

	fmt.Println(ans)
}
