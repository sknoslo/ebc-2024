package main

import (
	"fmt"
	"sknoslo/ebc2024/deques"
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
	// runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

func partone(notes string) any {
	type step struct {
		prev     vec2.Vec2
		pos      vec2.Vec2
		cost     int
		altitude int
		time     int
	}

	type cachekey struct {
		pos  vec2.Vec2
		time int
	}

	grid := grids.FromRunes(notes)
	start := grid.Find('S')
	seen := sets.New[cachekey](64)

	const startingAltitude = 1_000
	const targetTime = 100

	q := pqueues.New[step](64)
	q.Push(step{start, start, 0, startingAltitude, 0}, 0)

	for !q.Empty() {
		s := q.Pop()

		key := cachekey{s.pos, s.time}
		cell := grid.CellAt(s.pos)
		if cell == '#' || seen.Has(key) {
			continue
		}
		if s.time == targetTime {
			return s.altitude
		}

		seen.Insert(key)

		for _, dir := range vec2.CardinalDirs {
			npos := s.pos.Add(dir)
			if npos != s.prev && grid.InGrid(npos) {
				nalt := s.altitude
				ncost := s.cost

				switch grid.CellAt(npos) {
				case '-':
					nalt -= 2
					ncost += 4
				case '+':
					nalt += 1
					ncost += 1
				default:
					nalt -= 1
					ncost += 3
				}
				q.Push(step{s.pos, npos, ncost, nalt, s.time + 1}, ncost)
			}
		}
	}

	return "incomplete"
}

// slooooooooooooowwwww. Considering way too large of a search space because it's not clear
// to me what the optimal height ought to be at any particular checkpoint. The fastest path
// from S to A is not the right choice, need to account for the altitude as well. Some way
// of prioritizing both altitude and time would likely solve it faster.
func parttwo(notes string) any {
	type step struct {
		prev     vec2.Vec2
		pos      vec2.Vec2
		target   int
		altitude int
		time     int
	}

	type cachekey struct {
		pos      vec2.Vec2
		target   int
		altitude int
	}

	grid := grids.FromRunes(notes)
	start := grid.Find('S')
	seen := sets.New[cachekey](64)

	checkpoints := [4]rune{'A', 'B', 'C', 'S'}

	const startingAltitude = 10_000

	q := deques.New[step](64)
	q.PushFront(step{start, start, 0, startingAltitude, 0})

	for !q.Empty() {
		s := q.PopBack()

		target := checkpoints[s.target]
		key := cachekey{s.pos, s.target, s.altitude}
		cell := grid.CellAt(s.pos)
		ntarget := s.target
		if cell == '#' || seen.Has(key) {
			continue
		}

		seen.Insert(key)

		if cell == target {
			if target == 'S' && s.altitude >= startingAltitude {
				return s.time
			} else if target != 'S' {
				ntarget++
			}
		}

		for _, dir := range vec2.CardinalDirs {
			npos := s.pos.Add(dir)
			if npos != s.prev && grid.InGrid(npos) {
				nalt := s.altitude

				switch grid.CellAt(npos) {
				case '-':
					nalt -= 2
				case '+':
					nalt += 1
				default:
					nalt -= 1
				}
				q.PushFront(step{s.pos, npos, ntarget, nalt, s.time + 1})
			}
		}
	}

	return "incomplete"
}

func partthree(notes string) any {
	// The input has 3 columns that contain _only_ + signs, so the best path will always be to just go straight for the nearest
	// plus sign and then go straight down until the end.... so this is _not_ a general solution

	grid := grids.FromRunes(notes)
	start := grid.Find('S')

	firstPlusDist := 1 << 32
	firstPlus := vec2.Vec2{}

	height := grid.Size().Y

	for pos := range grid.FindCells('+') {
		dist := vec2.Distance(pos, start)
		if dist < firstPlusDist {
			firstPlus = pos
			firstPlusDist = dist
		}
	}

	altitude := 384400 - math.AbsDiff(firstPlus.X, start.X)
	altitudeLostPerSection := height

	for y := range height {
		switch grid.CellAtXY(firstPlus.X, y) {
		case '+':
			altitudeLostPerSection++
		case '.':
			altitudeLostPerSection--
		default:
			panic("This cheaty algorithm won't work on this input")
		}
	}

	dist := 0
	for altitude-altitudeLostPerSection > 0 {
		altitude -= altitudeLostPerSection
		dist += height
	}

	y := 0
	for altitude > 0 {
		switch grid.CellAtXY(firstPlus.X, y) {
		case '+':
			altitude++
		case '.':
			altitude--
		default:
			panic("how did you get here?")
		}
		dist++
		y++
	}

	return dist
}
