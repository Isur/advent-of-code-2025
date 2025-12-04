package Day01

import (
	"fmt"
	"time"

	"github.com/Isur/advent-of-code-2025/pkg"
)

func Run(data []string) {
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

func changePosition(position int, direction bool, change int) int {
	newPosition := position
	if direction {
		newPosition = newPosition + change
	} else {
		newPosition = newPosition - change
	}

	for newPosition >= 100 {
		newPosition = newPosition - 100
	}

	for newPosition < 0 {
		newPosition = newPosition + 100
	}

	return newPosition
}

func part1(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 1)
	}()
	position := 50
	zeros := 0
	for _, line := range data {
		if len(line) <= 1 {
			break
		}
		dir := string(line[0])
		change := pkg.ParseToInt(line[1:])
		position = changePosition(position, dir == "R", change)
		if position == 0 {
			zeros++
		}
	}

	return zeros
}

func changePositionPart2(position int, direction bool, change int) (int, int) {
	newPostition := position

	zeros := 0

	if change >= 100 {
		zeros += change / 100
	}

	if direction {
		newPostition += change % 100
	} else {
		newPostition -= change % 100
	}

	if newPostition == 0 {
		return zeros + 1, newPostition
	}

	if newPostition > 99 {
		newPostition -= 100
		if position != 0 {
			zeros++
		}
	}

	if newPostition < 0 {
		newPostition += 100
		if position != 0 {
			zeros++
		}
	}

	return zeros, newPostition
}

func part2(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 2)
	}()
	position := 50
	zeros := 0
	for _, line := range data {
		if len(line) <= 1 {
			break
		}
		dir := string(line[0])
		change := pkg.ParseToInt(line[1:])
		addZeros, newPostition := changePositionPart2(position, dir == "R", change)
		zeros += addZeros
		position = newPostition
	}

	return zeros
}
