package main

import (
	"fmt"
	"sknoslo/ebc2024/grids"
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/vec2"
	"slices"
	"strings"
)

var _ = fmt.Print // TODO: delete when done

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

func inRunicWord(grid *grids.Grid[rune], s vec2.Vec2, c rune) bool {
	for y := 2; y < 6; y++ {
		for x := 2; x < 6; x++ {
			if c == grid.CellAtXY(s.X+x, s.Y+y) {
				return true
			}
		}
	}

	return false
}

func buildRunicWord(grid *grids.Grid[rune], s vec2.Vec2) bool {
	complete := true
	for y := 2; y < 6; y++ {
		row := make([]rune, 0, 8)
		for x := 0; x < 8; x++ {
			c := grid.CellAtXY(s.X+x, s.Y+y)
			if c != '.' {
				row = append(row, c)
			}
		}
	xloop:
		for x := 2; x < 6; x++ {
			filled := 0
			for yy := 0; yy < 8; yy++ {
				c := grid.CellAtXY(s.X+x, s.Y+yy)
				if c != '.' && c != '?' && slices.Contains(row, c) && !inRunicWord(grid, s, c) {
					filled++
					grid.SetCellAtXY(s.X+x, s.Y+y, c)
					row = append(row, c)
					continue xloop
				}
			}
			if filled != 4 {
				complete = false
			}
		}
	}

	return complete
}

func getRunicWord(grid *grids.Grid[rune], s vec2.Vec2) string {
	b := strings.Builder{}
	for y := 2; y < 6; y++ {
		for x := 2; x < 6; x++ {
			c := grid.CellAtXY(s.X+x, s.Y+y)
			b.WriteRune(c)
		}
	}
	return b.String()
}

func partone(notes string) any {
	grid := grids.FromRunes(notes)
	s := vec2.New(0, 0)
	buildRunicWord(grid, s)
	return getRunicWord(grid, s)
}

func power(word string) int {
	p := 0
	for i, c := range word {
		p += (i + 1) * int(c+1-'A')
	}
	return p
}

func parttwo(notes string) any {
	grid := grids.FromRunes(notes)

	total := 0

	for y := 0; grid.InGrid(vec2.New(0, y)); y += 9 {
		for x := 0; grid.InGrid(vec2.New(x, 0)); x += 9 {
			s := vec2.New(x, y)
			buildRunicWord(grid, s)
			total += power(getRunicWord(grid, s))
		}
	}

	return total
}

// find items in a that are not in b
func difference(a, b []rune) []rune {
	res := make([]rune, 0, len(a))
mainloop:
	for _, aa := range a {
		for _, bb := range b {
			if aa == bb {
				continue mainloop
			}
		}
		res = append(res, aa)
	}
	return res
}

func colHasQues(grid *grids.Grid[rune], s vec2.Vec2) (vec2.Vec2, bool) {
	switch {
	case grid.CellAt(s) == '?':
		return s, true
	case grid.CellAtXY(s.X, s.Y+1) == '?':
		return s.Add(vec2.New(0, 1)), true
	case grid.CellAtXY(s.X, s.Y+6) == '?':
		return s.Add(vec2.New(0, 6)), true
	case grid.CellAtXY(s.X, s.Y+7) == '?':
		return s.Add(vec2.New(0, 7)), true
	default:
		return vec2.Vec2{}, false
	}
}

func rowHasQues(grid *grids.Grid[rune], s vec2.Vec2) (vec2.Vec2, bool) {
	switch {
	case grid.CellAt(s) == '?':
		return s, true
	case grid.CellAtXY(s.X+1, s.Y) == '?':
		return s.Add(vec2.New(1, 0)), true
	case grid.CellAtXY(s.X+6, s.Y) == '?':
		return s.Add(vec2.New(6, 0)), true
	case grid.CellAtXY(s.X+7, s.Y) == '?':
		return s.Add(vec2.New(7, 0)), true
	default:
		return vec2.Vec2{}, false
	}
}

func getColSplit(grid *grids.Grid[rune], s vec2.Vec2) ([]rune, []rune) {
	outer, inner := make([]rune, 0, 4), make([]rune, 0, 3)

	for y := 0; y < 8; y++ {
		c := grid.CellAtXY(s.X, s.Y+y)
		if y > 1 && y < 6 {
			inner = append(inner, c)
		} else {
			outer = append(outer, c)
		}
	}

	return outer, inner
}

func getRowSplit(grid *grids.Grid[rune], s vec2.Vec2) ([]rune, []rune) {
	outer, inner := make([]rune, 0, 4), make([]rune, 0, 3)

	for x := 0; x < 8; x++ {
		c := grid.CellAtXY(s.X+x, s.Y)
		if x > 1 && x < 6 {
			inner = append(inner, c)
		} else {
			outer = append(outer, c)
		}
	}

	return outer, inner
}

func fixRunicWord(grid *grids.Grid[rune], s vec2.Vec2) bool {
	fixed := true
	qs := make([]vec2.Vec2, 0, 2)
	for y := 2; y < 6; y++ {
		for x := 2; x < 6; x++ {
			xx, yy := s.X+x, s.Y+y
			c := grid.CellAtXY(xx, yy)
			if c == '.' {
				if ques, yes := rowHasQues(grid, vec2.New(s.X, s.Y+y)); yes {
					diff := difference(getColSplit(grid, vec2.New(s.X+x, s.Y)))
					if len(diff) == 1 && diff[0] != '?' && !inRunicWord(grid, s, diff[0]) {
						grid.SetCellAtXY(xx, yy, diff[0])
						grid.SetCellAt(ques, diff[0])
						qs = append(qs, ques)
						continue
					}
				} else if ques, yes := colHasQues(grid, vec2.New(s.X+x, s.Y)); yes {
					diff := difference(getRowSplit(grid, vec2.New(s.X, s.Y+y)))
					if len(diff) == 1 && diff[0] != '?' && !inRunicWord(grid, s, diff[0]) {
						grid.SetCellAtXY(xx, yy, diff[0])
						grid.SetCellAt(ques, diff[0])
						qs = append(qs, ques)
						continue
					}
				}
				fixed = false
			}
		}
	}

	if !fixed {
		// reset
		for _, q := range qs {
			grid.SetCellAt(q, '?')
		}
		for y := 2; y < 6; y++ {
			for x := 2; x < 6; x++ {
				xx, yy := s.X+x, s.Y+y
				grid.SetCellAtXY(xx, yy, '.')
			}
		}
	}

	return fixed
}

func partthree(notes string) any {
	grid := grids.FromRunes(notes)
	total := 0

	notsolved := make([]vec2.Vec2, 0, 100)

	for y := 0; grid.InGrid(vec2.New(0, y+7)); y += 6 {
		for x := 0; grid.InGrid(vec2.New(x+7, 0)); x += 6 {
			s := vec2.New(x, y)
			if buildRunicWord(grid, s) || fixRunicWord(grid, s) {
				total += power(getRunicWord(grid, s))
			} else {
				notsolved = append(notsolved, s)
			}
		}
	}

	changed := true
	for changed {
		changed = false
		nextnotsolved := make([]vec2.Vec2, 0, 100)

		for _, s := range notsolved {
			if buildRunicWord(grid, s) || fixRunicWord(grid, s) {
				changed = true
				total += power(getRunicWord(grid, s))
			} else {
				nextnotsolved = append(nextnotsolved, s)
			}
		}

		notsolved = nextnotsolved
	}

	return total
}
