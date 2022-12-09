package utils

import (
	"fmt"
	"sort"
	"strconv"
)

func Atoi(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Cannot convert int", err)
	}
	return val
}

func Min(values []int) int {
	m := -1

	for i, v := range values {
		if i == 0 || v < m {
			m = v
		}
	}

	return m
}

func Max(values []int) int {
	m := -1

	for i, v := range values {
		if i == 0 || v > m {
			m = v
		}
	}

	return m
}

func ReverseSort(values []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
}

func MapInt(data []int, f func(int) int) []int {

	mapped := make([]int, len(data))

	for i, e := range data {
		mapped[i] = f(e)
	}

	return mapped
}
