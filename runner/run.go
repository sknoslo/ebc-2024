package runner

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
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

func RunCpuPerf(notepath string, fn func(notes string) any) {
	notes, err := input.ReadNotes(notepath)
	if err != nil {
		fmt.Printf("Part %d: no notes, skipped\n", part)
	} else {
		f, err := os.Create(fmt.Sprintf("part%d.prof", part))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()

		res := fn(notes)
		fmt.Printf("Part %d: %v\n", part, res)
	}
	part++
}
