package main

import (
	"sknoslo/ebc2024/grids"
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/vec2"
	"strings"
)

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

var cw = []vec2.Vec2{vec2.North, vec2.NorthEast, vec2.East, vec2.SouthEast, vec2.South, vec2.SouthWest, vec2.West, vec2.NorthWest}
var ccw = []vec2.Vec2{vec2.North, vec2.NorthWest, vec2.West, vec2.SouthWest, vec2.South, vec2.SouthEast, vec2.East, vec2.NorthEast}

func partone(notes string) any {
	key, message, _ := strings.Cut(notes, "\n\n")

	grid := grids.FromRunes(message)

	size := grid.Size()

	i := 0
	for y := 1; y < size.Y-1; y++ {
		for x := 1; x < size.X-1; x++ {
			point := vec2.New(x, y)
			switch key[i%len(key)] {
			case 'R':
				prevValue := grid.CellAt(point.Add(vec2.NorthWest))
				for _, dir := range cw {
					tmp := grid.CellAt(point.Add(dir))
					grid.SetCellAt(point.Add(dir), prevValue)
					prevValue = tmp
				}
			case 'L':
				prevValue := grid.CellAt(point.Add(vec2.NorthEast))
				for _, dir := range ccw {
					tmp := grid.CellAt(point.Add(dir))
					grid.SetCellAt(point.Add(dir), prevValue)
					prevValue = tmp
				}
			default:
				panic("Found unsupported instruction")
			}
			i++
		}
	}

	var builder strings.Builder

	s := grid.Find('>')
	e := grid.Find('<')

	for x := s.X + 1; x < e.X; x++ {
		builder.WriteRune(grid.CellAtXY(x, s.Y))
	}

	return builder.String()
}

func parttwo(notes string) any {
	key, message, _ := strings.Cut(notes, "\n\n")

	grid := grids.FromRunes(message)

	size := grid.Size()

	for range 100 {
		i := 0
		for y := 1; y < size.Y-1; y++ {
			for x := 1; x < size.X-1; x++ {
				point := vec2.New(x, y)
				switch key[i%len(key)] {
				case 'R':
					prevValue := grid.CellAt(point.Add(vec2.NorthWest))
					for _, dir := range cw {
						tmp := grid.CellAt(point.Add(dir))
						grid.SetCellAt(point.Add(dir), prevValue)
						prevValue = tmp
					}
				case 'L':
					prevValue := grid.CellAt(point.Add(vec2.NorthEast))
					for _, dir := range ccw {
						tmp := grid.CellAt(point.Add(dir))
						grid.SetCellAt(point.Add(dir), prevValue)
						prevValue = tmp
					}
				default:
					panic("Found unsupported instruction")
				}
				i++
			}
		}
	}

	var builder strings.Builder

	s := grid.Find('>')
	e := grid.Find('<')

	for x := s.X + 1; x < e.X; x++ {
		builder.WriteRune(grid.CellAtXY(x, s.Y))
	}

	return builder.String()
}

func partthree(notes string) any {
	key, message, _ := strings.Cut(notes, "\n\n")
	grid := grids.FromRunes(message)
	size := grid.Size()

	// The answer is going to be when all of the digits 1-9 are in a line, so we need to know how many to look for
	targetSize := 0
	for y := 1; y < size.Y-1; y++ {
		for x := 1; x < size.X-1; x++ {
			c := grid.CellAtXY(x, y)
			if c >= '1' && c <= '9' {
				targetSize++
			}
		}
	}

	doCycle := func() {
		i := 0
		for y := 1; y < size.Y-1; y++ {
			for x := 1; x < size.X-1; x++ {
				point := vec2.New(x, y)
				switch key[i%len(key)] {
				case 'R':
					prevValue := grid.CellAt(point.Add(vec2.NorthWest))
					for _, dir := range cw {
						tmp := grid.CellAt(point.Add(dir))
						grid.SetCellAt(point.Add(dir), prevValue)
						prevValue = tmp
					}
				case 'L':
					prevValue := grid.CellAt(point.Add(vec2.NorthEast))
					for _, dir := range ccw {
						tmp := grid.CellAt(point.Add(dir))
						grid.SetCellAt(point.Add(dir), prevValue)
						prevValue = tmp
					}
				default:
					panic("Found unsupported instruction")
				}
				i++
			}
		}
	}

	// find the first time that the > and < line up in the right place, and capture the correct digits and the gaps
	answer := make([]rune, targetSize, targetSize)
	var gaps []vec2.Vec2
	for range 1048576000 {
		doCycle()

		s := grid.Find('>')
		e := grid.Find('<')

		if s.Y == e.Y && e.X-s.X-1 == targetSize {
			for x := s.X + 1; x < e.X; x++ {
				pos := vec2.New(x, s.Y)
				c := grid.CellAt(pos)
				if c < '1' || c > '9' {
					gaps = append(gaps, pos)
				} else {
					answer[x-s.X-1] = c
				}
			}
			break
		}

	}

	// find the other digits that fill the gaps
	remainingNums := make([]rune, len(gaps), len(gaps))
outerloop:
	for range 1048576000 {
		doCycle()

		for i, g := range gaps {
			c := grid.CellAt(g)
			if c >= '1' && c <= '9' {
				remainingNums[i] = c
				
			} else {
				continue outerloop
			}
		}

		break
	}

	// merge the answer and the gaps
	i := 0
	for k, v := range answer {
		if v == 0 {
			answer[k] = remainingNums[i]
			i++
		}
	}

	return string(answer)
}
