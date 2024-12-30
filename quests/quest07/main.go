package main

import (
	"cmp"
	"sknoslo/ebc2024/combo"
	"sknoslo/ebc2024/grids"
	"sknoslo/ebc2024/runner"
	"sknoslo/ebc2024/vec2"
	"slices"
	"strings"
)

func main() {
	runner.Run("part1.notes", partone)
	runner.Run("part2.notes", parttwo)
	runner.Run("part3.notes", partthree)
}

type Result struct {
	label   string
	essence int
}

func partone(notes string) any {
	lines := strings.Split(notes, "\n")
	results := make([]Result, 0, len(lines))

	for _, line := range lines {
		label, actionstr, found := strings.Cut(line, ":")
		if !found {
			panic("Failed to cut on ':' for string '" + line + "'")
		}
		actions := strings.Split(actionstr, ",")
		essence, power := 0, 0
		for i := range 10 {
			switch actions[i%len(actions)] {
			case "+":
				power++
			case "-":
				power--
			}
			essence += power
		}
		results = append(results, Result{label, essence})
	}

	slices.SortFunc(results, func(a, b Result) int {
		return cmp.Compare(b.essence, a.essence)
	})

	var b strings.Builder
	for _, res := range results {
		b.WriteString(res.label)
	}

	return b.String()
}

func trackToString(trackstr string) string {
	var b strings.Builder
	track := grids.FromRunes(strings.TrimSpace(trackstr))
	pos, dir := vec2.New(1, 0), vec2.East

	for track.CellAt(pos) != 'S' {
		b.WriteRune(track.CellAt(pos))

		dirs := [3]vec2.Vec2{dir, dir.RotateCardinalCCW(), dir.RotateCardinalCW()}
		for _, ndir := range dirs {
			npos := pos.Add(ndir)
			if track.InGrid(npos) && track.CellAt(npos) != ' ' {
				pos = npos
				dir = ndir
				break
			}
		}
	}

	b.WriteRune(track.CellAt(pos))

	return b.String()
}

func parttwo(notes string) any {
	const trackstr = `
S-=++=-==++=++=-=+=-=+=+=--=-=++=-==++=-+=-=+=-=+=+=++=-+==++=++=-=-=--
-                                                                     -
=                                                                     =
+                                                                     +
=                                                                     +
+                                                                     =
=                                                                     =
-                                                                     -
--==++++==+=+++-=+=-=+=-+-=+-=+-=+=-=+=--=+++=++=+++==++==--=+=++==+++-`
	track := trackToString(trackstr)
	lines := strings.Split(notes, "\n")
	results := make([]Result, 0, len(lines))

	for _, line := range lines {
		label, actionstr, found := strings.Cut(line, ":")
		if !found {
			panic("Failed to cut on ':' for string '" + line + "'")
		}
		actions := strings.Split(actionstr, ",")
		essence, power := 0, 0
		for i := range 10 * len(track) {
			action := actions[i%len(actions)]
			segment := track[i%len(track)]
			switch {
			case segment == '+':
				power++
			case segment == '-':
				power--
			case action == "+":
				power++
			case action == "-":
				power--
			}
			essence += power
		}
		results = append(results, Result{label, essence})
	}

	slices.SortFunc(results, func(a, b Result) int {
		return cmp.Compare(b.essence, a.essence)
	})

	var b strings.Builder
	for _, res := range results {
		b.WriteString(res.label)
	}

	return b.String()
}

func partthree(notes string) any {
	const trackstr = `
S+= +=-== +=++=     =+=+=--=    =-= ++=     +=-  =+=++=-+==+ =++=-=-=--
- + +   + =   =     =      =   == = - -     - =  =         =-=        -
= + + +-- =-= ==-==-= --++ +  == == = +     - =  =    ==++=    =++=-=++
+ + + =     +         =  + + == == ++ =     = =  ==   =   = =++=       
= = + + +== +==     =++ == =+=  =  +  +==-=++ =   =++ --= + =          
+ ==- = + =   = =+= =   =       ++--          +     =   = = =--= ==++==
=     ==- ==+-- = = = ++= +=--      ==+ ==--= +--+=-= ==- ==   =+=    =
-               = = = =   +  +  ==+ = = +   =        ++    =          -
-               = + + =   +  -  = + = = +   =        +     =          -
--==++++==+=+++-= =-= =-+-=  =+-= =-= =--   +=++=+++==     -=+=++==+++-`
	track := trackToString(trackstr)
	_, actionstr, found := strings.Cut(notes, ":")
	if !found {
		panic("Failed to cut on ':' for string '" + notes + "'")
	}

	var simulate func(actions []string) int
	simulate = func(actions []string) int {
		essence, power := 0, 0
		for i := range 2024 * len(track) {
			action := actions[i%len(actions)]
			segment := track[i%len(track)]
			switch {
			case segment == '+':
				power++
			case segment == '-':
				power--
			case action == "+":
				power++
			case action == "-":
				power--
			}
			essence += power
		}

		return essence
	}

	scoreToBeat := simulate(strings.Split(actionstr, ","))

	better := 0

	for actions := range combo.UniquePermutations(strings.Split("---===+++++", "")) {
		if simulate(actions) > scoreToBeat {
			better++
		}
	}

	return better
}
