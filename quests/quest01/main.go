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
	potions := 0
	for _, c := range notes {
		switch c {
		case 'B':
			potions++
		case 'C':
			potions += 3
		}
	}
	return potions
}

func parttwo(notes string) any {
	potions := 0
	for i := 0; i < len(notes); i += 2 {
		if notes[i] != 'x' && notes[i+1] != 'x' {
			potions += 2
		}
		for j := range 2 {
			switch notes[i+j] {
			case 'B':
				potions++
			case 'C':
				potions += 3
			case 'D':
				potions += 5
			}
		}
	}
	return potions
}

func partthree(notes string) any {
	potions := 0
	for i := 0; i < len(notes); i += 3 {
		bonus := 2
		for j := range 3 {
			if notes[i+j] == 'x' {
				bonus--
			}
		}

		for j := range 3 {
			switch notes[i+j] {
			case 'A':
				potions += bonus
			case 'B':
				potions += 1 + bonus
			case 'C':
				potions += 3 + bonus
			case 'D':
				potions += 5 + bonus
			}
		}
	}
	return potions
}
