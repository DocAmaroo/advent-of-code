package main

import (
	"aoc/common/files"
	"fmt"
	"strings"
)

func main() {
	ans := 0
	stream := files.Read("input.txt")[0]
	data := ""
	for i, r := range stream {
		idx := strings.IndexRune(data, r)
		if idx != -1 {
			data = data[idx+1:]
		}
		data += string(r)
		if len(data) == 4 {
			ans = i + 1
			break
		}
	}

	fmt.Println(ans)
}
