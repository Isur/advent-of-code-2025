package Day04

import (
	"fmt"
	"time"

	"github.com/Isur/advent-of-code-2025/pkg"
)

func Run(data []string) {
	fmt.Println("Part 1:", part1(data)) // example = 13
	fmt.Println("Part 2:", part2(data))
}

func validatePosition(x, y int, data []string) bool {
	if string(data[y][x]) != "@" {
		return false
	}
	lastX := len(data[y]) - 1
	lastY := len(data) - 1
	rolls := 0

	if y > 0 && x > 0 && string(data[y-1][x-1]) == "@" {
		rolls++
	}
	if y > 0 && string(data[y-1][x]) == "@" {
		rolls++
	}
	if y > 0 && x < lastX && string(data[y-1][x+1]) == "@" {
		rolls++
	}
	if x < lastX && string(data[y][x+1]) == "@" {
		rolls++
	}
	if x > 0 && string(data[y][x-1]) == "@" {
		rolls++
	}
	if y < lastY && x < lastX && string(data[y+1][x+1]) == "@" {
		rolls++
	}
	if y < lastY && string(data[y+1][x]) == "@" {
		rolls++
	}
	if y < lastY && x > 0 && string(data[y+1][x-1]) == "@" {
		rolls++
	}

	return rolls < 4
}

func part1(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 1)
	}()
	total := 0
	for i := range len(data) {
		for j := range len(data[i]) {
			res := validatePosition(j, i, data)
			if res {
				total++
			}
		}
	}
	return total
}

func part2(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 2)
	}()
	total := 0
	curr := 0
	for {
		curr = total
		for i := range len(data) {
			for j := range len(data[i]) {
				res := validatePosition(j, i, data)
				if res {
					r := []rune(data[i])
					r[j] = '.'
					data[i] = string(r)
					total++
				}
			}
		}

		if total == curr {
			break
		}
	}
	return total
}
