package main

import (
	"sknoslo/ebc2024/math"
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
	minimum := 100000000
	nails := strutils.SplitInts(notes, "\n")

	for _, nail := range nails {
		if nail < minimum {
			minimum = nail
		}
	}

	strikes := 0
	for _, nail := range nails {
		strikes += nail - minimum	
	}

	return strikes
}

func parttwo(notes string) any {
	minimum := 100000000
	nails := strutils.SplitInts(notes, "\n")

	for _, nail := range nails {
		if nail < minimum {
			minimum = nail
		}
	}

	strikes := 0
	for _, nail := range nails {
		strikes += nail - minimum	
	}

	return strikes
}

func partthree(notes string) any {
	nails := strutils.SplitInts(notes, "\n")
	slices.Sort(nails)

	median := nails[len(nails)/2]

	strikes := 0
	for _, nail := range nails {
		strikes += math.AbsDiff(nail, median)	
	}

	return strikes
}
