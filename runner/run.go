package runner

import (
	"fmt"
	"sknoslo/ebc2024/input"
)

var part = 1

func Run(notepath string, fn func(notes string) any) {
	notes, err := input.ReadNotes(notepath)
	if err != nil {
		fmt.Printf("Part %d: no notes, skipped\n", part)
	} else {
		fmt.Printf("Part %d: %v\n", part, fn(notes))
	}
	part++
}
