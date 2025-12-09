package Day09

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

type Box struct {
	X int
	Y int
}

func CreateBoxFromLine(line string) Box {
	coords := strings.Split(line, ",")
	x := pkg.ParseToInt(coords[0])
	y := pkg.ParseToInt(coords[1])
	box := Box{
		X: x,
		Y: y,
	}

	return box
}

func CalcBoxSize(a, b Box) int {
	one := pkg.Abs(a.X, b.X) + 1
	two := pkg.Abs(a.Y, b.Y) + 1

	size := one * two

	return size
}

func part1(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 1)
	}()
	boxes := []Box{}
	biggestSize := -1

	for _, line := range data {
		box := CreateBoxFromLine(line)
		for _, b := range boxes {
			size := CalcBoxSize(box, b)
			if size > biggestSize {
				biggestSize = size
			}
		}
		boxes = append(boxes, box)
	}

	return biggestSize
}

func part2(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 2)
	}()

	return 0
}
