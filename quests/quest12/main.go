package main

import (
	"fmt"
	"sknoslo/ebc2024/grids"
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/strutils"
	"sknoslo/ebc2024/vec2"
	"strings"
)

var _ = fmt.Print // TODO: delete when done

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

func partone(notes string) any {
	segments := [3]vec2.Vec2{}
	targets := []vec2.Vec2{}

	grid := grids.FromRunes(notes)

	for pos, c := range grid.Cells() {
		switch c {
		case '.', '=':
			continue
		case 'T':
			targets = append(targets, pos)
		default:
			segments[c-'A'] = pos
		}
	}

	rank := func(s, t vec2.Vec2, sn int) (int, bool) {
		dy := t.Y - s.Y
		nt := t.Sub(vec2.New(dy, dy))
		dx := nt.X - s.X

		if dx%3 != 0 {
			return -1, false
		}

		power := dx / 3

		return sn * power, true
	}

	total := 0
	for _, t := range targets {
		minRank := 1_000_000

		for i, s := range segments {
			if r, hit := rank(s, t, i+1); hit {
				minRank = min(minRank, r)
			}
		}

		total += minRank
	}

	return total
}

func parttwo(notes string) any {
	segments := [3]vec2.Vec2{}
	targets := []vec2.Vec2{}

	grid := grids.FromRunes(notes)

	for pos, c := range grid.Cells() {
		switch c {
		case '.', '=':
			continue
		case 'T', 'H':
			targets = append(targets, pos)
		default:
			segments[c-'A'] = pos
		}
	}

	rank := func(s, t vec2.Vec2, sn int) (int, bool) {
		dy := t.Y - s.Y
		nt := t.Sub(vec2.New(dy, dy))
		dx := nt.X - s.X

		if dx%3 != 0 {
			return -1, false
		}

		power := dx / 3

		return sn * power, true
	}

	total := 0
	for _, t := range targets {
		minRank := 1_000_000

		for i, s := range segments {
			if r, hit := rank(s, t, i+1); hit {
				minRank = min(minRank, r)
			}
		}

		mul := 1
		if grid.CellAt(t) == 'H' {
			mul = 2
		}
		total += minRank * mul
	}

	return total
}

func partthree(notes string) any {
	segments := [3]vec2.Vec2{vec2.New(0, 0), vec2.New(0, 1), vec2.New(0, 2)}

	rank := func(s, t vec2.Vec2, sn, dt int) (int, bool) {
		dx, dy := t.X-s.X, t.Y-s.Y

		// straight shot
		if dx == dy {
			return sn * dt, true
		}
		// at altitude
		if dx > dy && dx <= dy*2 {
			return sn * dy, true
		}
		// on down slope (there must be an exact solution like previous parts, but I )
		for power := (dx + dy) / 3; power*2 < t.X; power++ {
			dx, dy = t.X-power*2, s.Y+power-t.Y
			if dx == dy {
				return sn * power, true
			}
		}

		return -1, false
	}

	total := 0

	for _, line := range strings.Split(notes, "\n") {
		nums := strutils.SplitInts(line, " ")
		target := vec2.New(nums[0], nums[1])
		delay := target.X % 2
		time := target.X/2 + delay
		target = target.Sub(vec2.New(time, time))

		minRank := 100_000_000
		for i, s := range segments {
			if r, hit := rank(s, target, i+1, time-delay); hit {
				minRank = min(minRank, r)
			}
		}

		total += minRank
	}

	return total
}
