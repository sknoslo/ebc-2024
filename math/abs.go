package math

func Abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

func AbsDiff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
