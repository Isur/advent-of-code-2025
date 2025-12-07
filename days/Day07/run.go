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
	beamsMap := map[int]bool{firstBeam: true}

	splits := 0

	for _, line := range data[1:] {
		for beam, _ := range beamsMap {
			if string(line[beam]) == "^" {
				delete(beamsMap, beam)
				splits++
				if beam > 0 {
					beamsMap[beam-1] = true
				}
				if beam < len(line) {
					beamsMap[beam+1] = true
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

	beams := map[int]map[int]int{
		0: {firstBeamPosition: 1},
	}

	lvl := 0

	for _, line := range data[1:] {
		if isEmpty(line) {
			continue
		}
		found := false
		for beam, _ := range beams[lvl] {

			_, ok := beams[lvl+1]
			if !ok {
				beams[lvl+1] = map[int]int{}
			}

			if string(line[beam]) == "^" {
				found = true
				if beam > 0 {
					_, ok := beams[lvl+1][beam-1]
					if ok {
						beams[lvl+1][beam-1] += beams[lvl][beam]
					} else {
						beams[lvl+1][beam-1] = beams[lvl][beam]
					}
				}
				if beam < len(line) {
					_, ok := beams[lvl+1][beam+1]
					if ok {
						beams[lvl+1][beam+1] += beams[lvl][beam]
					} else {
						beams[lvl+1][beam+1] = beams[lvl][beam]
					}
				}
			} else {
				_, ok := beams[lvl+1][beam]
				if ok {
					beams[lvl+1][beam] += beams[lvl][beam]
				} else {
					beams[lvl+1][beam] = beams[lvl][beam]
				}
			}
		}

		if found {
			lvl++
		}
	}

	sum := 0
	for _, v := range beams[len(beams)-1] {
		sum += v
	}

	return sum
}
