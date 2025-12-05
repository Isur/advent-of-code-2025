package Day05

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

type Range struct {
	begin int
	end   int
}

func getRanges(line string) Range {
	r := strings.Split(line, "-")
	begin := pkg.ParseToInt(r[0])
	end := pkg.ParseToInt(r[1])
	return Range{
		begin: begin,
		end:   end,
	}
}

func validateId(id int, ranges []Range) bool {
	for _, r := range ranges {
		if id >= r.begin && id <= r.end {
			return true
		}
	}
	return false
}

func part1(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 1)
	}()

	sum := 0
	ranges := []Range{}
	b := false

	for _, line := range data {
		if line == "" {
			b = true
			continue
		}
		if b == false {
			ranges = append(ranges, getRanges(line))
		} else {
			id := pkg.ParseToInt(line)

			ok := validateId(id, ranges)
			if ok {
				sum++
			}
		}
	}

	return sum
}

func mergeRanges(ranges []Range) []Range {
	newRanges := []Range{}

	for _, ra := range ranges {
		found := false
		for i, _ := range newRanges {
			if ra.begin >= newRanges[i].begin && ra.begin <= newRanges[i].end && ra.end > newRanges[i].end {
				newRanges[i].end = ra.end
				found = true
			}

			if ra.end <= newRanges[i].end && ra.end >= newRanges[i].begin && ra.begin < newRanges[i].begin {
				newRanges[i].begin = ra.begin
				found = true
			}

			if ra.begin <= newRanges[i].begin && ra.end >= newRanges[i].end {
				newRanges[i].begin = ra.begin
				newRanges[i].end = ra.end
				found = true
			}

			if ra.begin >= newRanges[i].begin && ra.end <= newRanges[i].end {
				found = true
			}

			if found {
				break
			}
		}

		if found == false {
			newRanges = append(newRanges, ra)
		}
	}

	return newRanges
}

func calcIds(ranges []Range) int {
	sum := 0
	for _, ra := range ranges {
		sum += ra.end - ra.begin + 1
	}

	return sum
}

func part2(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 2)
	}()

	ranges := []Range{}

	for _, line := range data {
		if line == "" {
			break
		}
		ra := getRanges(line)
		ranges = append(ranges, ra)
	}

	currLen := len(ranges)
	ranges = mergeRanges(ranges)
	for len(ranges) != currLen {
		currLen = len(ranges)
		ranges = mergeRanges(ranges)
	}

	return calcIds(ranges)
}
