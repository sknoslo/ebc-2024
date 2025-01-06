package main

import (
	"fmt"
	"maps"
	"math"
	"sknoslo/ebc2024/runner"
	"strings"
)

var _ = fmt.Print // TODO: delete when done

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

func partone(notes string) any {
	lines := strings.Split(notes, "\n")
	rules := make(map[string][]string, len(lines))
	for _, line := range lines {
		category, conversions, _ := strings.Cut(line, ":")
		rules[category] = strings.Split(conversions, ",")
	}

	var count func(category string, remdays int) int
	count = func(category string, remdays int) int {
		if 0 == remdays {
			return 1
		}

		c := 0
		for _, conv := range rules[category] {
			c += count(conv, remdays - 1)
		}
		return c
	}
	return count("A", 4)
}

type CacheKey struct {
	remdays int
	category string
}

func parttwo(notes string) any {
	lines := strings.Split(notes, "\n")
	rules := make(map[string][]string, len(lines))
	for _, line := range lines {
		category, conversions, _ := strings.Cut(line, ":")
		rules[category] = strings.Split(conversions, ",")
	}

	cache := make(map[CacheKey]int, 0)
	var count func(category string, remdays int) int
	count = func(category string, remdays int) int {
		if 0 == remdays {
			return 1
		}

		key := CacheKey{remdays, category}
		if v, ok := cache[key]; ok {
			return v
		}

		c := 0
		for _, conv := range rules[category] {
			c += count(conv, remdays - 1)
		}
		cache[key] = c
		return c
	}
	return count("Z", 10)
}

func partthree(notes string) any {
	lines := strings.Split(notes, "\n")
	rules := make(map[string][]string, len(lines))
	for _, line := range lines {
		category, conversions, _ := strings.Cut(line, ":")
		rules[category] = strings.Split(conversions, ",")
	}

	cache := make(map[CacheKey]int, 0)
	var count func(category string, remdays int) int
	count = func(category string, remdays int) int {
		if 0 == remdays {
			return 1
		}

		key := CacheKey{remdays, category}
		if v, ok := cache[key]; ok {
			return v
		}

		c := 0
		for _, conv := range rules[category] {
			c += count(conv, remdays - 1)
		}
		cache[key] = c
		return c
	}


	smallest, largest := math.MaxInt, 0

	for k := range maps.Keys(rules) {
		total := count(k, 20)

		if total < smallest {
			smallest = total
		}
		if total > largest {
			largest = total
		}
	}
	return largest - smallest
}
