package Day06

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

func filterLine[T ~string | ~int](line string, number bool) []T {
	splitted := strings.Fields(line)
	var result []T

	for _, str := range splitted {
		var v T
		if number {
			num := pkg.ParseToInt(str)
			v = any(num).(T)
		} else {
			v = any(str).(T)
		}

		result = append(result, v)
	}

	return result
}

func recalc(results []int, line []int, operations []string) {
	for i, _ := range results {
		if operations[i] == "+" {
			results[i] += line[i]
		} else if operations[i] == "*" {
			results[i] *= line[i]
		}
	}
}

func part1(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 1)
	}()

	operations := filterLine[string](data[len(data)-1], false)

	results := make([]int, len(data)-1)

	for i, line := range data {
		if i == len(data)-1 {
			break
		}
		filtered := filterLine[int](line, true)
		if i == 0 {
			results = filtered
		} else {
			recalc(results, filtered, operations)
		}
	}

	sum := 0
	for res := range results {
		sum += results[res]
	}

	return sum
}

func part2(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 2)
	}()
	operations := filterLine[string](data[len(data)-1], false)
	results := []int{}
	currArr := []int{}
	for i := len(data[0]) - 1; i >= 0; i-- {
		digits := []string{}
		for j := 0; j < len(data)-1; j++ {
			curr := string(data[j][i])
			if curr != "" && curr != " " {
				digits = append(digits, curr)
			}
		}

		if len(digits) == 0 || i == 0 {
			if i == 0 {
				lol := strings.Join(strings.Fields(strings.Join(digits, "")), "")
				number := pkg.ParseToInt(lol)
				currArr = append(currArr, number)
			}
			currOperation := operations[len(operations)-1]
			res := currArr[0]
			for _, num := range currArr[1:] {
				if currOperation == "+" {
					res += num
				} else if currOperation == "*" {
					res *= num
				}
			}

			results = append(results, res)
			operations = operations[:len(operations)-1]
			currArr = []int{}
		} else {
			lol := strings.Join(strings.Fields(strings.Join(digits, "")), "")
			number := pkg.ParseToInt(lol)
			currArr = append(currArr, number)
		}
	}

	sum := 0
	for _, num := range results {
		sum += num
	}
	return sum
}
