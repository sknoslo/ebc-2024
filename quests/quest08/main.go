package main

import (
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/strutils"
)

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

func partone(notes string) any {
	available := strutils.MustAtoi(notes)
	blocks := 0

	layer := -1
	for blocks < available {
		layer += 2
		blocks += layer
	}

	return (blocks - available) * layer
}

func parttwo(notes string) any {
	available := 20240000
	priests := strutils.MustAtoi(notes)
	acolytes := 1111

	thickness := 1
	blocks := 0
	layerwidth := 1
	nextlayer := 1

	for blocks+nextlayer < available {
		blocks += nextlayer

		thickness = (thickness * priests) % acolytes
		layerwidth += 2
		nextlayer = layerwidth * thickness
	}
	blocks += nextlayer

	return (blocks - available) * layerwidth
}

func partthree(notes string) any {
	available := 202400000
	priests := strutils.MustAtoi(notes)
	acolytes := 10
	columns := make([]int, 0, 4096)

	thickness := 1
	blocks := 0
	layerwidth := 1
	nextlayer := 1
	toremove := 0

	columns = append(columns, 1)

	for blocks+nextlayer-toremove < available {
		blocks += nextlayer

		thickness = (thickness*priests)%acolytes + acolytes
		layerwidth += 2
		nextlayer = layerwidth * thickness
		columns = append(columns, 0)
		for i := range len(columns) {
			columns[i] += thickness
		}
		toremove = 0
		for i := range len(columns) - 1 {
			r := (priests * layerwidth * columns[i]) % acolytes
			if i == 0 {
				toremove += r
			} else {
				toremove += r * 2
			}
		}
	}
	blocks += nextlayer - toremove

	return blocks - available
}
