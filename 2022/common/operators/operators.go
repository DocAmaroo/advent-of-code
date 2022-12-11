package operators

import (
	"fmt"
	"strconv"
)

func Atoi(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Cannot convert int", err)
	}
	return val
}

// A Python modulo representation.
// Returns the rest of the division as a Python modulo do.
func PyMod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

// Transform a boolean to an integer
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
