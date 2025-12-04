package Day02

import (
	"fmt"
	"strings"
	"time"

	"github.com/Isur/advent-of-code-2025/pkg"
)

func Run(data []string) {
	data = strings.Split(data[0], ",")
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

func validateId(number int) int {
	stringified := pkg.ParseIntToString(number)

	strLen := len(stringified)

	if strLen%2 != 0 {
		return 0
	}

	half := strLen / 2

	for i := 0; i < half; i++ {
		if stringified[i] != stringified[half+i] {
			return 0
		}
	}

	return number
}

func compare(slice []string) bool {
	firstItem := slice[0]
	for _, item := range slice[1:] {
		if firstItem != item {
			return false
		}
	}

	return true
}

func chunkify(slice string, len int, items int) []string {
	chunks := make([]string, items)
	for i := range items {
		index := i * len
		chunks[i] = slice[index : index+len]
	}
	return chunks
}

func validateSequence(number int) int {
	stringified := pkg.ParseIntToString(number)

	strLen := len(stringified)
	if strLen < 2 {
		return 0
	}

	half := strLen / 2

	possibleDivisors := []int{}

	for i := 2; i <= half; i++ {
		if strLen%i == 0 {
			possibleDivisors = append(possibleDivisors, i)
		}
	}
	possibleDivisors = append(possibleDivisors, strLen)

	for _, divisor := range possibleDivisors {
		chunks := chunkify(stringified, strLen/divisor, divisor)
		result := compare(chunks)
		if result {
			return number
		}
	}

	return 0
}

func validateRange(begin, end int, part2 bool) int {
	numbers := 0

	for i := begin; i <= end; i++ {
		if part2 {
			res := validateSequence(i)
			numbers += res
		} else {
			numbers += validateId(i)
		}
	}

	return numbers
}

func part1(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 1)
	}()
	invalidIds := 0
	for _, line := range data {
		splitted := strings.Split(line, "-")
		begin := pkg.ParseToInt(splitted[0])
		end := pkg.ParseToInt(splitted[1])
		invalidIds += validateRange(begin, end, false)
	}
	return invalidIds
}

func part2(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 2)
	}()
	invalidIds := 0
	for _, line := range data {
		splitted := strings.Split(line, "-")
		begin := pkg.ParseToInt(splitted[0])
		end := pkg.ParseToInt(splitted[1])
		r := validateRange(begin, end, true)
		invalidIds += r
	}
	return invalidIds
}
