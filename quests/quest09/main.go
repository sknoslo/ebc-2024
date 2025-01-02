package main

import (
	"math"
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/strutils"
	"slices"
)

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

func partone(notes string) any {
	brightnesses := strutils.SplitInts(notes, "\n")

	var beetles func(brightness int) int
	beetles = func(brightness int) int {
		switch {
		case brightness == 0:
			return 0
		case brightness < 3:
			return brightness
		case brightness < 5:
			return 1 + beetles(brightness-3)
		case brightness < 10:
			return 1 + beetles(brightness-5)
		default:
			return brightness/10 + beetles(brightness%10)
		}
	}

	total := 0
	for _, brightness := range brightnesses {
		total += beetles(brightness)

	}

	return total
}

func getFewest(stamps []int, rem int, cache map[int]int) int {
	if rem == 0 {
		return 0
	}

	if f, ok := cache[rem]; ok {
		return f
	}

	m := math.MaxInt

	for _, s := range stamps {
		if rem - s >= 0 {
			cost := 1 + getFewest(stamps, rem - s, cache)
			if cost < m {
				m = cost
			}
		}
	}

	cache[rem] = m
	return m
}

func parttwo(notes string) any {
	stamps := []int{1, 3, 5, 10, 15, 16, 20, 24, 25, 30}
	slices.Reverse(stamps)
	brightnesses := strutils.SplitInts(notes, "\n")
	total := 0
	cache := make(map[int]int, 5000)
	for _, brightness := range brightnesses {
		total += getFewest(stamps, brightness, cache)
	}
	return total
}

func partthree(notes string) any {
	stamps := []int{1, 3, 5, 10, 15, 16, 20, 24, 25, 30, 37, 38, 49, 50, 74, 75, 100, 101}
	slices.Reverse(stamps)
	brightnesses := strutils.SplitInts(notes, "\n")
	total := 0
	cache := make(map[int]int, 50000)
	for _, brightness := range brightnesses {
		var a, b int
		if brightness % 2 == 0 {
			a = brightness / 2
			b = a
		} else {
			a = brightness / 2
			b = a + 1
		}
		mincost := math.MaxInt
		for b - a <= 100 {
			cost := getFewest(stamps, a, cache) + getFewest(stamps, b, cache)	
			if cost < mincost {
				mincost = cost
			}
			b++
			a--
		}
		total += mincost
	}
	return total
}
