package main

import (
	"sknoslo/ebc2024/runner"
	"strings"
)

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

func partone(notes string) any {
	lines := strings.Split(notes, "\n")
	tree := make(map[string][]string, len(lines))

	for _, line := range lines {
		node, branches, found := strings.Cut(line, ":")
		if !found {
			panic("expected a colon, but didn't find one in '" + line + "'")
		}

		tree[node] = strings.Split(branches, ",")
	}

	var getPaths func(node string) []string
	getPaths = func(node string) []string {
		if node == "@" {
			return []string{node}
		}
		branches := tree[node]
		out := make([]string, 0, len(branches)*2)
		for _, branch := range branches {
			for _, path := range getPaths(branch) {
				out = append(out, node+path)
			}
		}
		return out
	}

	paths := getPaths("RR")

	lenMap := make(map[int][]string, len(paths)/2+1)

	for _, path := range paths {
		l := len(path)
		lenMap[l] = append(lenMap[l], path)
	}

	for _, p := range lenMap {
		if len(p) == 1 {
			return p[0]
		}
	}

	return "no unique path found"
}

func parttwo(notes string) any {
	lines := strings.Split(notes, "\n")
	tree := make(map[string][]string, len(lines))

	for _, line := range lines {
		node, branches, found := strings.Cut(line, ":")
		if !found {
			panic("expected a colon, but didn't find one in '" + line + "'")
		}

		tree[node] = strings.Split(branches, ",")
	}

	var getPaths func(node string) []string
	getPaths = func(node string) []string {
		if node == "@" {
			return []string{node}
		}
		branches := tree[node]
		out := make([]string, 0, len(branches)*2)
		for _, branch := range branches {
			for _, path := range getPaths(branch) {
				out = append(out, node[0:1]+path)
			}
		}
		return out
	}

	paths := getPaths("RR")

	lenMap := make(map[int][]string, len(paths)/2+1)

	for _, path := range paths {
		l := len(path)
		lenMap[l] = append(lenMap[l], path)
	}

	for _, p := range lenMap {
		if len(p) == 1 {
			return p[0]
		}
	}

	return "no unique path found"
}

func partthree(notes string) any {
	lines := strings.Split(notes, "\n")
	tree := make(map[string][]string, len(lines))

	for _, line := range lines {
		node, branches, found := strings.Cut(line, ":")
		if node == "BUG" || node == "ANT" {
			continue
		}
		if !found {
			panic("expected a colon, but didn't find one in '" + line + "'")
		}

		tree[node] = strings.Split(branches, ",")
	}

	var getPaths func(node string) []string
	getPaths = func(node string) []string {
		if node == "@" {
			return []string{node}
		}
		branches := tree[node]
		out := make([]string, 0, len(branches)*2)
		for _, branch := range branches {
			for _, path := range getPaths(branch) {
				out = append(out, node[0:1]+path)
			}
		}
		return out
	}

	paths := getPaths("RR")

	lenMap := make(map[int][]string, len(paths)/2+1)

	for _, path := range paths {
		l := len(path)
		lenMap[l] = append(lenMap[l], path)
	}

	for _, p := range lenMap {
		if len(p) == 1 {
			return p[0]
		}
	}

	return "no unique path found"
}
