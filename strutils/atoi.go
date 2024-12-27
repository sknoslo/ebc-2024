package strutils

import (
	"strconv"
	"strings"
)

func MustAtoi(s string) int {
	v, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return v
}

func SplitInts(s string, sep string) []int {
	nums := strings.Split(s, sep)
	out := make([]int, 0, len(nums))

	for _, num := range nums {
		out = append(out, MustAtoi(num))
	}

	return out
}
