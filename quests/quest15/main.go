package main

import (
	"sknoslo/ebc2024/deques"
	"sknoslo/ebc2024/grids"
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/sets"
	"sknoslo/ebc2024/vec2"
)

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

func partone(notes string) any {
	type Step struct {
		pos   vec2.Vec2
		steps int
	}

	grid := grids.FromRunes(notes)
	start := grid.Find('.')
	q := deques.New[Step](64)
	q.PushFront(Step{start, 0})

	for !q.Empty() {
		s := q.PopBack()

		if grid.CellAt(s.pos) == 'H' {
			return s.steps * 2
		}

		if grid.CellAt(s.pos) == '#' {
			continue
		}
		grid.SetCellAt(s.pos, '#')

		for _, dir := range vec2.CardinalDirs {
			npos := s.pos.Add(dir)
			if grid.InGrid(npos) {
				q.PushFront(Step{npos, s.steps + 1})
			}
		}
	}

	return "incomplete"
}

func parttwo(notes string) any {
	const allHerbs uint8 = 0b11111 // Assumes ABCDE, so not general

	type Key struct {
		pos   vec2.Vec2
		herbs uint8
	}

	has := func(herbs uint8, herb rune) bool {
		bit := uint8(1 << (herb - 'A'))
		return herbs&bit == 1
	}

	pickup := func(herbs uint8, herb rune) uint8 {
		bit := uint8(1 << (herb - 'A'))
		return herbs | bit
	}

	type Step struct {
		pos   vec2.Vec2
		full  bool
		herbs uint8
		steps int
	}

	grid := grids.FromRunes(notes)
	start := grid.Find('.')
	q := deques.New[Step](64)
	q.PushFront(Step{pos: start})
	seen := sets.New[Key](64)

	for !q.Empty() {
		s := q.PopBack()

		if s.full && s.pos == start {
			return s.steps
		}

		cell := grid.CellAt(s.pos)

		if cell == '#' || cell == '~' {
			continue
		}

		key := Key{s.pos, s.herbs}
		if seen.Has(key) {
			continue
		}
		seen.Insert(key)

		ns := s
		ns.steps++
		if !s.full && cell != '.' && !has(s.herbs, cell) {
			ns.herbs = pickup(s.herbs, cell)
			ns.full = ns.herbs == allHerbs
		}

		for _, dir := range vec2.CardinalDirs {
			ns.pos = s.pos.Add(dir)
			if grid.InGrid(ns.pos) {
				q.PushFront(ns)
			}
		}
	}

	return "incomplete"
}

// TODO: optimization
// This just does a BFS through the entire search space, it's very slow (about 2 minutes on my computer) and gobbles
// memory. Possible optimization routes:
//  1. Compress the graph to just the "interesting" nodes and use Dijkstra's.
//  2. Take advantage of the input structure. Part 2 was plenty fast and Part 3 is structured as 3 columns that could each
//     be solved individually using Part 2. The tricky bit is finding the sub problems without hard-coding. Should complete
//     just 100ms or so though. (hint: treat K as an exit to solve either side and you have to pick up K twice)
func partthree(notes string) any {
	type Key struct {
		pos   vec2.Vec2
		herbs uint32
	}

	has := func(herbs uint32, herb rune) bool {
		bit := uint32(1 << (herb - 'A'))
		return herbs&bit == 1
	}

	pickup := func(herbs uint32, herb rune) uint32 {
		bit := uint32(1 << (herb - 'A'))
		return herbs | bit
	}

	type Step struct {
		pos   vec2.Vec2
		full  bool
		herbs uint32
		steps int
	}

	grid := grids.FromRunes(notes)

	var allHerbs uint32 = 0

	for _, cell := range grid.Cells() {
		switch cell {
		case '.', '#', '~':
			continue
		default:
			allHerbs = pickup(allHerbs, cell)
		}
	}

	start := grid.Find('.')
	q := deques.New[Step](64)
	q.PushFront(Step{pos: start})
	seen := sets.New[Key](64)

	for !q.Empty() {
		s := q.PopBack()

		if s.full && s.pos == start {
			return s.steps
		}

		cell := grid.CellAt(s.pos)

		if cell == '#' || cell == '~' {
			continue
		}

		key := Key{s.pos, s.herbs}
		if seen.Has(key) {
			continue
		}
		seen.Insert(key)

		ns := s
		ns.steps++
		if !s.full && cell != '.' && !has(s.herbs, cell) {
			ns.herbs = pickup(s.herbs, cell)
			ns.full = ns.herbs == allHerbs
		}

		for _, dir := range vec2.CardinalDirs {
			ns.pos = s.pos.Add(dir)
			if grid.InGrid(ns.pos) {
				q.PushFront(ns)
			}
		}
	}

	return "incomplete"
}
