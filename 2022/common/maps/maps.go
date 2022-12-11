package maps

func GetValues(m map[string]int) []int {
	res := make([]int, 0)
	for _, v := range m {
		res = append(res, v)
	}

	return res
}
