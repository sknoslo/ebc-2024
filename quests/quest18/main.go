package main

import (
	"sknoslo/ebc2024/deques"
	"sknoslo/ebc2024/grids"
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/vec2"
)

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
	size := original.Size()

	timegrids := make([]*grids.Grid[int], 0, 64)

	for start := range original.FindCells('P') {
		grid := grids.FromRunes(notes)
		times := grids.FromSize(size, 0)

		q := deques.New[step](16)
		q.PushFront(step{start, 0})

		for !q.Empty() {
			s := q.PopBack()

			cell := grid.CellAt(s.pos)
			if cell == '~' || cell == '#' {
				continue
			}
			if cell == '.' {
				times.SetCellAt(s.pos, s.time)
			}
			grid.SetCellAt(s.pos, '~')

			for _, dir := range vec2.CardinalDirs {
				npos := s.pos.Add(dir)

				if grid.InGrid(npos) && times.CellAt(npos) == 0 {
					q.PushFront(step{npos, s.time + 1})
				}
			}
		}

		timegrids = append(timegrids, times)
	}

	answer := 1 << 32
	for pos := range original.FindCells('.') {
		sumOfTimes := 0
		for _, times := range timegrids {
			sumOfTimes += times.CellAt(pos)
		}
		answer = min(sumOfTimes, answer)
	}

	return answer
}
