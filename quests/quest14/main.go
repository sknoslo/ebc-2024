package main

import (
	"fmt"
	"sknoslo/ebc2024/deques"
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/sets"
	"sknoslo/ebc2024/strutils"
	"sknoslo/ebc2024/vec3"
	"strings"
)

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

func partone(notes string) any {
	maxheight := 0
	height := 0
	for _, step := range strings.Split(notes, ",") {
		segments := strutils.MustAtoi(step[1:])
		switch step[0] {
		case 'U':
			height += segments
		case 'D':
			height -= segments
		}
		maxheight = max(maxheight, height)
	}
	return maxheight
}

func parttwo(notes string) any {
	unique := sets.New[vec3.Vec3](128)
	for _, plan := range strings.Split(notes, "\n") {
		pos := vec3.New(0, 0, 0)
		for _, step := range strings.Split(plan, ",") {
			segments := strutils.MustAtoi(step[1:])
			dir := vec3.Vec3{}
			switch step[0] {
			case 'U':
				dir.Y = 1
			case 'D':
				dir.Y = -1
			case 'R':
				dir.X = 1
			case 'L':
				dir.X = -1
			case 'F':
				dir.Z = 1
			case 'B':
				dir.Z = -1
			}

			for range segments {
				pos = pos.Add(dir)
				unique.Insert(pos)
			}
		}
	}

	return unique.Count()
}

var directions = []vec3.Vec3{
	{X: 0, Y: 1, Z: 0},  // UP
	{X: 0, Y: -1, Z: 0}, // DOWN
	{X: 1, Y: 0, Z: 0},  // RIGHT
	{X: -1, Y: 0, Z: 0}, // LEFT
	{X: 0, Y: 0, Z: 1},  // FORWARD
	{X: 0, Y: 0, Z: -1}, // BACK
}

type Vec3Set struct {
	data    []bool
	w, h, d int
	offset  vec3.Vec3
}

func fromBounds(bmin, bmax vec3.Vec3) *Vec3Set {
	box := bmax.Sub(bmin).Add(vec3.New(1, 1, 1))
	size := box.X * box.Y * box.Z
	return &Vec3Set{
		data:   make([]bool, size),
		w:      box.X,
		h:      box.Y,
		d:      box.Z,
		offset: bmin,
	}
}

func fromSet(s *sets.Set[vec3.Vec3], bmin, bmax vec3.Vec3) *Vec3Set {
	set := fromBounds(bmin, bmax)

	for cell := range s.Items() {
		set.Insert(cell)
	}

	return set
}

func (c *Vec3Set) toIndex(pos vec3.Vec3) int {
	pos = pos.Sub(c.offset)
	return pos.Z*c.w*c.h + pos.Y*c.w + pos.X
}

func (c *Vec3Set) Has(pos vec3.Vec3) bool {
	idx := c.toIndex(pos)
	if idx >= len(c.data) || idx < 0 {
		return false
	}
	return c.data[idx]
}

func (c *Vec3Set) Insert(pos vec3.Vec3) {
	idx := c.toIndex(pos)
	if idx >= len(c.data) {
		fmt.Println(pos, "no space for")
	}
	c.data[c.toIndex(pos)] = true
}

type Step struct {
	pos  vec3.Vec3
	dist int
}

func dist(tree *Vec3Set, start, end, bmin, bmax vec3.Vec3) int {
	seen := fromBounds(bmin, bmax)
	var q = deques.New[Step](128)

	q.PushFront(Step{start, 0})

	for !q.Empty() {
		step := q.PopBack()

		if step.pos == end {
			return step.dist
		}

		if seen.Has(step.pos) {
			continue
		}

		seen.Insert(step.pos)

		for _, dir := range directions {
			npos := step.pos.Add(dir)
			if tree.Has(npos) {
				q.PushFront(Step{npos, step.dist + 1})
			}
		}
	}

	panic(fmt.Sprintf("never found path from %v to %v", start, end))
}

func partthree(notes string) any {
	lines := strings.Split(notes, "\n")
	leaves := sets.New[vec3.Vec3](len(lines))
	tree := sets.New[vec3.Vec3](1000)
	trunkHeight := 0
	bmin, bmax := vec3.New(1000, 1000, 1000), vec3.Vec3{}
	for _, plan := range lines {
		pos := vec3.New(0, 0, 0)
		for _, step := range strings.Split(plan, ",") {
			segments := strutils.MustAtoi(step[1:])
			dir := vec3.Vec3{}
			switch step[0] {
			case 'U':
				dir.Y = 1
			case 'D':
				dir.Y = -1
			case 'R':
				dir.X = 1
			case 'L':
				dir.X = -1
			case 'F':
				dir.Z = 1
			case 'B':
				dir.Z = -1
			}

			for range segments {
				pos = pos.Add(dir)
				if pos.X == 0 && pos.Z == 0 {
					trunkHeight = max(trunkHeight, pos.Y)
				}
				tree.Insert(pos)
				bmin.X = min(bmin.X, pos.X)
				bmin.Y = min(bmin.Y, pos.Y)
				bmin.Z = min(bmin.Z, pos.Z)
				bmax.X = max(bmax.X, pos.X)
				bmax.Y = max(bmax.Y, pos.Y)
				bmax.Z = max(bmax.Z, pos.Z)
			}
		}
		leaves.Insert(pos)
	}

	lowestMurkiness := 1 << 32
	stree := fromSet(tree, bmin, bmax)
	// With my input, you could speed this up with a binary search looking for a local minimum... but I can't prove
	// why that should be true (gaps it the trunk should make that not guaranteed?) so I'm not doing it.
	for y := 1; y <= trunkHeight; y++ {
		candidate := vec3.New(0, y, 0)

		murkiness := 0
		for leaf := range leaves.Items() {
			murkiness += dist(stree, candidate, leaf, bmin, bmax)
		}

		lowestMurkiness = min(lowestMurkiness, murkiness)
	}

	return lowestMurkiness
}
