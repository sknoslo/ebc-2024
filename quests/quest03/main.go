package main

import (
	"sknoslo/ebc2024/grids"
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/vec2"
)

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

func partone(notes string) any {
	grid := grids.FromRunes(notes)

	prev, blocks := -1, 0
	for prev != blocks {
		prev = blocks
		nextgrid := grids.Clone(grid)
	gridloop:
		for pos, ch := range grid.Cells() {
			if ch == '.' {
				continue
			}
			if ch != '#' {
				for _, dir := range vec2.CardinalDirs {
					npos := pos.Add(dir)
					if !grid.InGrid(npos) || grid.CellAt(npos) != ch {
						continue gridloop
					}
				}
			}

			nextgrid.SetCellAt(pos, ch+1)
			blocks++
		}
		grid = nextgrid
	}
	return blocks
}

func parttwo(notes string) any {
	grid := grids.FromRunes(notes)

	prev, blocks := -1, 0
	for prev != blocks {
		prev = blocks
		nextgrid := grids.Clone(grid)
	gridloop:
		for pos, ch := range grid.Cells() {
			if ch == '.' {
				continue
			}
			if ch != '#' {
				for _, dir := range vec2.CardinalDirs {
					npos := pos.Add(dir)
					if !grid.InGrid(npos) || grid.CellAt(npos) != ch {
						continue gridloop
					}
				}
			}

			nextgrid.SetCellAt(pos, ch+1)
			blocks++
		}
		grid = nextgrid
	}
	return blocks
}

func partthree(notes string) any {
	grid := grids.FromRunes(notes)

	prev, blocks := -1, 0
	for prev != blocks {
		prev = blocks
		nextgrid := grids.Clone(grid)
	gridloop:
		for pos, ch := range grid.Cells() {
			if ch == '.' {
				continue
			}
			if ch != '#' {
				for _, dir := range vec2.AllDirs {
					npos := pos.Add(dir)
					if !grid.InGrid(npos) || grid.CellAt(npos) != ch {
						continue gridloop
					}
				}
			}

			nextgrid.SetCellAt(pos, ch+1)
			blocks++
		}
		grid = nextgrid
	}
	return blocks
}
