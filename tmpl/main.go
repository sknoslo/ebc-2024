package main

import (
	"sknoslo/ebc2024/runner"
)


func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

func partone(notes string) any {
	return "incomplete"
}

func parttwo(notes string) any {
	return "incomplete"
}

func partthree(notes string) any {
	return "incomplete"
}
