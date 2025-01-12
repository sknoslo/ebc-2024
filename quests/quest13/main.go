package main

import (
	"fmt"
	"sknoslo/ebc2024/grids"
	"sknoslo/ebc2024/math"
	"sknoslo/ebc2024/pqueues"
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/sets"
	"sknoslo/ebc2024/vec2"
)

var _ = fmt.Print // TODO: delete when done

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

type Step struct {
	plat vec2.Vec2
	time int
}

func level(plat rune) int {
	switch {
	case plat == 'S' || plat == 'E':
		return 0
	case plat >= '0' && plat <= '9':
		return int(plat - '0')
	default:
		panic("Can't get the level of a non-platform")
	}
}

func partone(notes string) any {
	grid := grids.FromRunes(notes)
	s := grid.Find('S')
	q := pqueues.New[Step](64)
	q.Push(Step{s, 0}, 0)
	seen := sets.New[vec2.Vec2](64)

	for !q.Empty() {
		step := q.Pop()
		if seen.Has(step.plat) {
			continue
		}
		seen.Insert(step.plat)

		o := grid.CellAt(step.plat)
		if o == 'E' {
			return step.time
		}

		ol := level(o)

		for _, dir := range vec2.CardinalDirs {
			nplat := step.plat.Add(dir)
			if grid.InGrid(nplat) && grid.CellAt(nplat) != '#' {
				nl := level(grid.CellAt(nplat))
				delta := min(math.AbsDiff(ol, nl), math.AbsDiff(ol+10, nl), math.AbsDiff(ol-10, nl))
				q.Push(Step{nplat, step.time + delta + 1}, step.time+delta+1)
			}
		}
	}

	return "no solution"
}

func parttwo(notes string) any {
	grid := grids.FromRunes(notes)
	s := grid.Find('S')
	q := pqueues.New[Step](64)
	q.Push(Step{s, 0}, 0)
	seen := sets.New[vec2.Vec2](64)

	for !q.Empty() {
		step := q.Pop()
		if seen.Has(step.plat) {
			continue
		}
		seen.Insert(step.plat)

		o := grid.CellAt(step.plat)
		if o == 'E' {
			return step.time
		}

		ol := level(o)

		for _, dir := range vec2.CardinalDirs {
			nplat := step.plat.Add(dir)
			if grid.InGrid(nplat) && grid.CellAt(nplat) != '#' {
				nl := level(grid.CellAt(nplat))
				delta := min(math.AbsDiff(ol, nl), math.AbsDiff(ol+10, nl), math.AbsDiff(ol-10, nl))
				q.Push(Step{nplat, step.time + delta + 1}, step.time+delta+1)
			}
		}
	}

	return "no solution"
}

func partthree(notes string) any {
	grid := grids.FromRunes(notes)
	q := pqueues.New[Step](64)
	for s, _ := range grid.FindCells('S') {
		q.Push(Step{s, 0}, 0)
	}
	seen := sets.New[vec2.Vec2](64)

	for !q.Empty() {
		step := q.Pop()
		if seen.Has(step.plat) {
			continue
		}
		seen.Insert(step.plat)

		o := grid.CellAt(step.plat)
		if o == 'E' {
			return step.time
		}

		ol := level(o)

		for _, dir := range vec2.CardinalDirs {
			nplat := step.plat.Add(dir)
			if !grid.InGrid(nplat) {
				continue
			}

			n := grid.CellAt(nplat)

			if n != '#' && n != 'S' {
				nl := level(grid.CellAt(nplat))
				delta := min(math.AbsDiff(ol, nl), math.AbsDiff(ol+10, nl), math.AbsDiff(ol-10, nl))
				q.Push(Step{nplat, step.time + delta + 1}, step.time+delta+1)
			}
		}
	}

	return "no solution"
}
