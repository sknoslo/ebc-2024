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

func reverse(str string) string {
	l := len(str)
	out := make([]rune, l)

	for i, c := range str {
		out[l-i-1] = c
	}

	return string(out)
}

func parseNotes(notes string) ([]string, string) {
	wordsin, inscription, found := strings.Cut(notes, "\n\n")
	if !found {
		panic("Could not cut notes on double new line")
	}

	wordsin, found = strings.CutPrefix(wordsin, "WORDS:")
	if !found {
		panic("Could not cut prefix 'WORDS:'")
	}
	words := strings.Split(wordsin, ",")

	return words, inscription
}

func partone(notes string) any {
	words, inscription := parseNotes(notes)

	wordcount := 0
	for _, word := range words {
		wordcount += strings.Count(inscription, word)
	}

	return wordcount
}

func parttwo(notes string) any {
	words, inscription := parseNotes(notes)

	for i := range len(words) {
		words = append(words, reverse(words[i]))
	}

	runecount := 0
	for _, sentence := range strings.Split(inscription, "\n") {
		seen := make([]bool, len(sentence))
		for _, word := range words {
			for i := range sentence {
				if i+len(word) <= len(sentence) && sentence[i:i+len(word)] == word {
					for j := i; j < i+len(word); j++ {
						if !seen[j] {
							runecount++
							seen[j] = true
						}
					}
				}
			}
		}
	}

	return runecount
}

func partthree(notes string) any {
	words, inscription := parseNotes(notes)

	scales := grids.FromRunes(inscription)
	size := scales.Size()
	seen := grids.FromSize(size, false)

	var countRunes func(pos, dir vec2.Vec2, word string) bool
	countRunes = func (pos, dir vec2.Vec2, word string) bool {
		if len(word) == 0 {
			return true
		}

		if !scales.InGrid(pos) || scales.CellAt(pos) != rune(word[0]) {
			return false
		}

		npos := pos.Add(dir)
		npos.X = (npos.X + size.X) % size.X
		if countRunes(npos, dir, word[1:]) {
			seen.SetCellAt(pos, true)
			return true
		}

		return false
	}

	for po := range scales.Cells() {
		for _, word := range words {
			for _, dir := range vec2.CardinalDirs {
				countRunes(po, dir, word)
			}
		}
	}

	runecount := 0
	for _, hasRune := range seen.Cells() {
		if hasRune {
			runecount++
		}
	}

	return runecount
}
