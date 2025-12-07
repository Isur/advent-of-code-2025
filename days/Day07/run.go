package Day07

import (
	"fmt"
	"strings"
	"time"

	"github.com/Isur/advent-of-code-2025/pkg"
)

func Run(data []string) {
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

func findStart(line string) int {
	for i, pos := range line {
		if string(pos) == "S" {
			return i
		}
	}

	return -1
}

func part1(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 1)
	}()

	firstBeam := findStart(data[0])
	beamSet := pkg.CreateSet[int, bool](nil)
	beamSet.Add(firstBeam, true)

	splits := 0

	for _, line := range data[1:] {
		for beam, _ := range beamSet.GetMap() {
			if string(line[beam]) == "^" {
				beamSet.Delete(beam)
				splits++
				if beam > 0 {
					beamSet.Add(beam-1, true)
				}
				if beam < len(line) {
					beamSet.Add(beam+1, true)
				}
			}
		}
	}

	return splits
}

func isEmpty(line string) bool {
	if strings.Contains(line, "^") {
		return false
	}
	return true
}

func part2(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 2)
	}()
	firstBeamPosition := findStart(data[0])

	beamSet := pkg.CreateSet[int, *pkg.NumericalSet[int, int]](nil)
	beamSet.Add(0, pkg.CreateNumericalSet[int, int](map[int]int{firstBeamPosition: 1}))

	lvl := 0

	beams := beamSet.GetMap()
	for _, line := range data[1:] {
		if isEmpty(line) {
			continue
		}
		found := false
		for beam, _ := range beams[lvl].GetMap() {

			_, ok := beams[lvl+1]
			if !ok {
				beams[lvl+1] = pkg.CreateNumericalSet[int, int](nil)
			}

			c := beams[lvl].GetMap()[beam]
			if string(line[beam]) == "^" {
				found = true
				if beam > 0 {
					beams[lvl+1].AddOrSet(beam-1, c)
				}
				if beam < len(line) {
					beams[lvl+1].AddOrSet(beam+1, c)
				}
			} else {
				beams[lvl+1].AddOrSet(beam, c)
			}
		}

		if found {
			lvl++
		}
	}

	sum := 0
	for _, v := range beamSet.GetMap()[beamSet.GetSize()-1].GetMap() {
		sum += v
	}

	return sum
}
