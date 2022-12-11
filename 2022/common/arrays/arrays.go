package arrays

import (
	op "aoc/common/operators"
	"sort"
)

func ApplyAtoi(arr []string) []int {
	res := make([]int, 0, len(arr))
	for _, a := range arr {
		i := op.Atoi(a)
		res = append(res, i)
	}
	return res
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

func SumAll(data []int) int {
	sum := 0
	for _, v := range data {
		sum += v
	}
	return sum
}

func FilterInt(data []int, f func(int) bool) []int {
	res := make([]int, 0)
	for _, v := range data {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
