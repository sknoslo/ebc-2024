package main

import (
	"fmt"
	"sknoslo/ebc2024/deques"
	"sknoslo/ebc2024/grids"
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/vec2"
)

var _ = fmt.Print // TODO: delete when done

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

type step struct {
	pos  vec2.Vec2
	time int
}

func partone(notes string) any {
	grid := grids.FromRunes(notes)
	start := vec2.New(0, 1)
	q := deques.New[step](16)
	q.PushFront(step{start, 0})

	minutesToLastPalm := 0

	for !q.Empty() {
		s := q.PopBack()

		cell := grid.CellAt(s.pos)
		if cell == '~' || cell == '#' {
			continue
		}
		if cell == 'P' {
			minutesToLastPalm = s.time
		}
		grid.SetCellAt(s.pos, '~')

		for _, dir := range vec2.CardinalDirs {
			npos := s.pos.Add(dir)

			if grid.InGrid(npos) {
				q.PushFront(step{npos, s.time + 1})
			}
		}
	}

	return minutesToLastPalm
}

func parttwo(notes string) any {
	grid := grids.FromRunes(notes)
	start := vec2.New(0, 1)
	end := vec2.New(200, 69)
	q := deques.New[step](16)
	q.PushFront(step{start, 0})
	q.PushFront(step{end, 0})

	minutesToLastPalm := 0

	for !q.Empty() {
		s := q.PopBack()

		cell := grid.CellAt(s.pos)
		if cell == '~' || cell == '#' {
			continue
		}
		if cell == 'P' {
			minutesToLastPalm = s.time
		}
		grid.SetCellAt(s.pos, '~')

		for _, dir := range vec2.CardinalDirs {
			npos := s.pos.Add(dir)

			if grid.InGrid(npos) {
				q.PushFront(step{npos, s.time + 1})
			}
		}
	}

	return minutesToLastPalm
}

func partthree(notes string) any {
	original := grids.FromRunes(notes)

	answer := 1 << 32

	// TODO: a faster way...
	// from each tree flood fill, maintaining a grid of distances to each empty segment
	// then iterate over the grid and sum the times from each segment and find a min
	for start, startcell := range original.Cells() {
		grid := grids.FromRunes(notes)
		if startcell != '.' {
			continue
		}

		q := deques.New[step](16)
		q.PushFront(step{start, 0})

		sumOfTimes := 0

		for !q.Empty() {
			s := q.PopBack()

			cell := grid.CellAt(s.pos)
			if cell == '~' || cell == '#' {
				continue
			}
			if cell == 'P' {
				sumOfTimes += s.time
			}
			grid.SetCellAt(s.pos, '~')

			for _, dir := range vec2.CardinalDirs {
				npos := s.pos.Add(dir)

				if grid.InGrid(npos) {
					q.PushFront(step{npos, s.time + 1})
				}
			}
		}

		answer = min(sumOfTimes, answer)
	}

	return answer
}
