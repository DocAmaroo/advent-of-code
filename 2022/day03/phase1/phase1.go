package main

import (
	"aoc/common/files"
	"fmt"
	"strings"
)

func main() {
	ans := 0
	lines := files.Read("input.txt")

	for _, line := range lines {
		half := (len(line) / 2)
		slA := line[0:half]
		slB := line[half:]

		for _, r := range slA {
			if strings.ContainsRune(slB, r) {
				if r > 96 {
					ans += int(r-'a') + 1
				} else {
					ans += int(r-'A') + 27
				}

				break
			}
		}
	}

	fmt.Println(ans)
}
