package main

import (
	"cmp"
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/sets"
	"sknoslo/ebc2024/vec2"
	"slices"
	"strings"
)

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

func partone(notes string) any {
	allstars := sets.New[vec2.Vec2](4)
	constellation := sets.New[vec2.Vec2](4)

	for y, line := range strings.Split(notes, "\n") {
		for x, c := range line {
			if c == '*' {
				if constellation.Count() == 0 {
					constellation.Insert(vec2.New(x, y))
				} else {
					allstars.Insert(vec2.New(x, y))
				}
			}
		}
	}

	total := 0

	for allstars.Count() > 0 {
		mindist := 1 << 32
		var minstar vec2.Vec2

		for a := range constellation.Items() {
			for b := range allstars.Items() {
				if vec2.Distance(a, b) < mindist {
					mindist = vec2.Distance(a, b)
					minstar = b
				}
			}
		}

		allstars.Remove(minstar)
		constellation.Insert(minstar)
		total += mindist
	}

	total += constellation.Count()

	return total
}

func parttwo(notes string) any {
	allstars := sets.New[vec2.Vec2](4)
	constellation := sets.New[vec2.Vec2](4)

	for y, line := range strings.Split(notes, "\n") {
		for x, c := range line {
			if c == '*' {
				if constellation.Count() == 0 {
					constellation.Insert(vec2.New(x, y))
				} else {
					allstars.Insert(vec2.New(x, y))
				}
			}
		}
	}

	total := 0

	for allstars.Count() > 0 {
		mindist := 1 << 32
		var minstar vec2.Vec2

		for a := range constellation.Items() {
			for b := range allstars.Items() {
				if vec2.Distance(a, b) < mindist {
					mindist = vec2.Distance(a, b)
					minstar = b
				}
			}
		}

		allstars.Remove(minstar)
		constellation.Insert(minstar)
		total += mindist
	}

	total += constellation.Count()

	return total
}

func partthree(notes string) any {
	allstars := sets.New[vec2.Vec2](4)

	for y, line := range strings.Split(notes, "\n") {
		for x, c := range line {
			if c == '*' {
				allstars.Insert(vec2.New(x, y))
			}
		}
	}

	brilliant := make([]int, 0, 16)

	for allstars.Count() > 0 {
		constellation := sets.New[vec2.Vec2](4)
		total := 0

		// weird use of an iterator... maybe add a way to "pop" a random item from a set?
		for starter := range allstars.Items() {
			constellation.Insert(starter)
			allstars.Remove(starter)
			break
		}

		for {
			mindist := 1 << 32
			var minstar vec2.Vec2

			for a := range constellation.Items() {
				for b := range allstars.Items() {
					if vec2.Distance(a, b) < mindist {
						mindist = vec2.Distance(a, b)
						minstar = b
					}
				}
			}

			if mindist >= 6 {
				break
			}

			total += mindist
			constellation.Insert(minstar)
			allstars.Remove(minstar)
		}

		total += constellation.Count()

		brilliant = append(brilliant, total)
	}

	slices.SortFunc(brilliant, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	return brilliant[0] * brilliant[1] * brilliant[2]
}
