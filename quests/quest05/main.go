package main

import (
	"container/list"
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/strutils"
	"strings"
)

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

func partone(notes string) any {
	const cols = 4
	var columns [cols]list.List

	for _, line := range strings.Split(notes, "\n") {
		for col, num := range strings.Fields(line) {
			columns[col].PushBack(strutils.MustAtoi(num))
		}
	}

	for round := range 10 {
		col := round % cols

		leader := columns[col].Remove(columns[col].Front()).(int)
		col = (col + 1) % cols
		onleft := true

		rem := leader

		pos := columns[col].Front()

		for rem > 1 {
			rem--
			var next *list.Element
			if onleft {
				next = pos.Next()
			} else {
				next = pos.Prev()
			}

			if next == nil {
				onleft = !onleft
			} else {
				pos = next
			}
		}
		if onleft {
			columns[col].InsertBefore(leader, pos)
		} else {
			columns[col].InsertAfter(leader, pos)
		}

		bignumber := 0
		for _, column := range columns {
			bignumber = bignumber*10 + column.Front().Value.(int)
		}
	}

	bignumber := 0
	for _, column := range columns {
		bignumber = bignumber*10 + column.Front().Value.(int)
	}

	return bignumber
}

func parttwo(notes string) any {
	const cols = 4
	var columns [cols]list.List
	seen := make(map[int]int)

	for _, line := range strings.Split(notes, "\n") {
		for col, num := range strings.Fields(line) {
			columns[col].PushBack(strutils.MustAtoi(num))
		}
	}

	for round := 1; round < 100_000_000; round++ {
		col := (round - 1) % cols

		for columns[col].Len() == 1 {
			col = (col + 1) % cols
		}

		leader := columns[col].Remove(columns[col].Front()).(int)
		col = (col + 1) % cols
		onleft := true

		rem := leader

		pos := columns[col].Front()

		for rem > 1 {
			rem--
			var next *list.Element
			if onleft {
				next = pos.Next()
			} else {
				next = pos.Prev()
			}

			if next == nil {
				onleft = !onleft
			} else {
				pos = next
			}
		}
		if onleft {
			columns[col].InsertBefore(leader, pos)
		} else {
			columns[col].InsertAfter(leader, pos)
		}

		bignumber := 0
		for _, column := range columns {
			bignumber *= 100
			front := column.Front()
			if front != nil {
				bignumber += front.Value.(int)
			}
		}
		seen[bignumber]++
		if seen[bignumber] == 2024 {
			return bignumber * round
		}
	}

	return "no answer"
}

func partthree(notes string) any {
	const cols = 4
	var columns [cols]list.List
	seen := make(map[int]int)
	largest := 0

	for _, line := range strings.Split(notes, "\n") {
		for col, num := range strings.Fields(line) {
			columns[col].PushBack(strutils.MustAtoi(num))
		}
	}

	for round := 1; round < 100_000_000; round++ {
		col := (round - 1) % cols

		for columns[col].Len() == 1 {
			col = (col + 1) % cols
		}

		leader := columns[col].Remove(columns[col].Front()).(int)
		col = (col + 1) % cols
		onleft := true

		rem := leader

		pos := columns[col].Front()

		rem = rem % (columns[col].Len() * 2)

		for rem > 1 {
			rem--
			var next *list.Element
			if onleft {
				next = pos.Next()
			} else {
				next = pos.Prev()
			}

			if next == nil {
				onleft = !onleft
			} else {
				pos = next
			}
		}
		if onleft {
			columns[col].InsertBefore(leader, pos)
		} else {
			columns[col].InsertAfter(leader, pos)
		}

		bignumber := 0
		for _, column := range columns {
			bignumber *= 10000
			front := column.Front()
			if front != nil {
				bignumber += front.Value.(int)
			}
		}
		if largest < bignumber {
			largest = bignumber
		}
		seen[bignumber]++
		// for a guaranteed answer, would need to do proper cycle detection
		if seen[bignumber] == 20 {
			return largest
		}
	}

	return "no answer"
}
