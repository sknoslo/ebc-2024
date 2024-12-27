package runner

import (
	"fmt"
	"sknoslo/ebc2024/input"
	"time"
)

var part = 1

func Run(notepath string, fn func(notes string) any) {
	notes, err := input.ReadNotes(notepath)
	if err != nil {
		fmt.Printf("Part %d: no notes, skipped\n", part)
	} else {
		s := time.Now()
		res := fn(notes)
		e := time.Now().Sub(s)
		fmt.Printf("Part %d: %v (%v)\n", part, res, e)
	}
	part++
}
