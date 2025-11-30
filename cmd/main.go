package main

import (
	"fmt"
	"os"

	"github.com/Isur/advent-of-code-2025/days/Day01"
	"github.com/Isur/advent-of-code-2025/days/Day02"
	"github.com/Isur/advent-of-code-2025/days/Day03"
	"github.com/Isur/advent-of-code-2025/days/Day04"
	"github.com/Isur/advent-of-code-2025/days/Day05"
	"github.com/Isur/advent-of-code-2025/days/Day06"
	"github.com/Isur/advent-of-code-2025/days/Day07"
	"github.com/Isur/advent-of-code-2025/days/Day08"
	"github.com/Isur/advent-of-code-2025/days/Day09"
	"github.com/Isur/advent-of-code-2025/days/Day10"
	"github.com/Isur/advent-of-code-2025/days/Day11"
	"github.com/Isur/advent-of-code-2025/days/Day12"
	"github.com/Isur/advent-of-code-2025/pkg"
)

type Settings struct {
	day  string
	mode bool
}

func setSettings() Settings {
	args := len(os.Args)
	if args < 2 {
		fmt.Println("Two args required, day and mode")
		os.Exit(1)
	}
	if args > 3 {
		fmt.Println("Too much args")
		fmt.Println("Two args required, day and mode")
		os.Exit(1)
	}

	mode := false
	if args == 3 {
		if os.Args[2] == "true" || os.Args[2] == "input" {
			mode = true
		} else if os.Args[2] == "false" || os.Args[2] == "example" {
			mode = false
		} else {
			fmt.Println("Wrong arg for mode, use one of [true, false, input, example]")
			os.Exit(1)
		}
	}

	settings := Settings{
		day:  os.Args[1],
		mode: mode,
	}

	return settings
}

func displaySettings(settings Settings) {
	fmt.Println("Day:  ", settings.day)
	fmt.Println("Mode: ", settings.mode)
	fmt.Println("==============")
}

func main() {
	settings := setSettings()
	displaySettings(settings)

	data := pkg.LoadByLine(settings.day, settings.mode)
	switch settings.day {
	case "01":
		Day01.Run(data)
	case "02":
		Day02.Run(data)
	case "03":
		Day03.Run(data)
	case "04":
		Day04.Run(data)
	case "05":
		Day05.Run(data)
	case "06":
		Day06.Run(data)
	case "07":
		Day07.Run(data)
	case "08":
		Day08.Run(data)
	case "09":
		Day09.Run(data)
	case "10":
		Day10.Run(data)
	case "11":
		Day11.Run(data)
	case "12":
		Day12.Run(data)
	default:
		fmt.Println("Day ", settings.day, " not exists here")
		os.Exit(1)
	}
}
