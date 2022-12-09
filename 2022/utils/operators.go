package utils

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
func BoolToint(b bool) int {
	if b {
		return 1
	}
	return 0
}
