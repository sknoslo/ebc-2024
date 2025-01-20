package main

import (
	"fmt"
	"maps"
	"sknoslo/ebc2024/math"
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/strutils"
	"strings"
)

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

func partone(notes string) any {
	const numWheels int = 4
	const faceSize int = 3
	const offset int = faceSize + 1
	ruleStr, wheelStr, _ := strings.Cut(notes, "\n\n")

	rules := strutils.SplitInts(ruleStr, ",")
	wheels := [numWheels][]string{}

	for _, line := range strings.Split(wheelStr, "\n") {
		for i := 0; i < numWheels && i*offset < len(line); i++ {
			face := line[i*offset : i*offset+faceSize]
			if face != "   " {
				wheels[i] = append(wheels[i], face)
			}
		}
	}

	var res [numWheels]string
	for i := range numWheels {
		wheel := wheels[i]
		idx := (rules[i] * 100) % len(wheel)
		res[i] = wheel[idx]
	}

	return strings.Join(res[:], " ")
}

func parttwo(notes string) any {
	const spins int = 202_420_242_024
	const faceSize int = 3
	const offset int = faceSize + 1
	ruleStr, wheelStr, _ := strings.Cut(notes, "\n\n")

	rules := strutils.SplitInts(ruleStr, ",")
	numWheels := len(rules)
	wheels := make([][]string, numWheels)

	for _, line := range strings.Split(wheelStr, "\n") {
		for i := 0; i < numWheels && i*offset < len(line); i++ {
			face := line[i*offset : i*offset+faceSize]
			if face != "   " {
				wheels[i] = append(wheels[i], face)
			}
		}
	}

	lcm := len(wheels[0])
	for _, wheel := range wheels[1:] {
		lcm = math.Lcm(lcm, len(wheel))
	}

	doSpins := func (spins int) int {
		coins := 0
		counter := make(map[byte]int, numWheels*2)
		for spin := 1; spin <= spins; spin++ {
			for i, wheel := range wheels {
				idx := (rules[i] * spin) % len(wheel)
				face := wheel[idx]
				counter[face[0]]++
				counter[face[2]]++
			}
			for v := range maps.Values(counter) {
				if v > 2 {
					coins += v - 2
				}
			}
			clear(counter)
		}

		return coins
	}

	return doSpins(lcm) * (spins/lcm) + doSpins(spins%lcm)
}

type Key struct {
	spins, offset int
}

type MaxMin struct {
	max, min int
}

func partthree(notes string) any {
	const targetSpins int = 256
	const faceSize int = 3
	const offset int = faceSize + 1
	ruleStr, wheelStr, _ := strings.Cut(notes, "\n\n")

	rules := strutils.SplitInts(ruleStr, ",")
	numWheels := len(rules)
	wheels := make([][]string, numWheels)

	for _, line := range strings.Split(wheelStr, "\n") {
		for i := 0; i < numWheels && i*offset < len(line); i++ {
			face := line[i*offset : i*offset+faceSize]
			if face != "   " {
				wheels[i] = append(wheels[i], face)
			}
		}
	}

	lcm := len(wheels[0])
	for _, wheel := range wheels[1:] {
		lcm = math.Lcm(lcm, len(wheel))
	}

	cache := make(map[Key]MaxMin, 256)
	var doSpin func(spins, offset int) MaxMin
	doSpin = func (spins, offset int) MaxMin {
		key := Key{spins, offset}
		if v, ok := cache[key]; ok {
			return v
		}

		coins := 0
		if spins > 0 {
			counter := make(map[byte]int, numWheels*2)
			res := make([]string, numWheels)
			for i, wheel := range wheels {
				idx := ((rules[i] * spins) + offset) % len(wheel)
				face := wheel[idx]
				counter[face[0]]++
				counter[face[2]]++
				res[i] = face
			}
			for v := range maps.Values(counter) {
				if v > 2 {
					coins += v - 2
				}
			}
		}

		if spins == targetSpins {
			return MaxMin{coins, coins}
		}

		v := MaxMin{0, 1000000000}
		for o := -1; o < 2; o++ {
			n := doSpin(spins + 1, offset + o)
			v.min = min(n.min+coins, v.min)
			v.max = max(n.max+coins, v.max)
		}

		cache[key] = v

		return v
	}

	mm := doSpin(0, 0)
	return fmt.Sprintf("%d %d", mm.max, mm.min)
}
