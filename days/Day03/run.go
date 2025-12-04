package Day03

import (
	"fmt"
	"time"

	"github.com/Isur/advent-of-code-2025/pkg"
)

func Run(data []string) {
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

func findPair(arr []int) int {
	a := 0

	for i := 1; i < len(arr)-1; i++ {
		if arr[a] < arr[i] {
			a = i
		}
	}

	b := a + 1

	for i := a + 1; i < len(arr); i++ {
		if arr[i] > arr[b] {
			b = i
		}
	}

	stringifiedSum := fmt.Sprintf("%d%d", arr[a], arr[b])

	return pkg.ParseToInt(stringifiedSum)
}

func split(line string) []int {
	arr := []int{}
	for i := range len(line) {
		num := pkg.ParseToInt(string(line[i]))
		arr = append(arr, num)
	}

	return arr
}

func find12(arr []int) int {
	result := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

	for i := range 12 {
		start := 0
		if i > 0 {
			start = result[i-1] + 1
		}
		for j := start; j < len(arr)-(12-i-1); j++ {
			if result[i] == -1 {
				result[i] = j
			} else if arr[result[i]] < arr[j] {
				result[i] = j
			}
		}

	}

	res := ""
	for _, item := range result {
		res = fmt.Sprintf("%v%v", res, arr[item])
	}

	return pkg.ParseToInt(res)
}

func part1(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 1)
	}()
	sum := 0
	for _, line := range data {
		arr := split(line)
		number := findPair(arr)
		sum += number
	}
	return sum
}

func part2(data []string) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		pkg.PrintTiming(elapsed, 2)
	}()
	sum := 0
	for _, line := range data {
		arr := split(line)
		number := find12(arr)
		sum += number
	}
	return sum
}
