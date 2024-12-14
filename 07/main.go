package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	inputS := string(input)

	part1 := 0
	part2 := 0

	for _, line := range strings.Split(inputS, "\n") {
		if line == "" {
			continue
		}

		part1 += processLinePart1(line)
		part2 += processLinePart2(line)
	}

	println()
	fmt.Println("part 1", part1)
	fmt.Println("part 2", part2)
}

func processLinePart1(line string) int {
	lineParts := strings.Split(line, ": ")
	goal, _ := strconv.Atoi(lineParts[0])
	options := strings.Split(lineParts[1], " ")
	possibilities := math.Pow(2, float64(len(options)-1))

	for i := 0; i < int(possibilities); i++ {
		total, _ := strconv.Atoi(options[0])

		for j := 0; j < len(options)-1; j++ {
			operator := (i >> j) & 1
			next, _ := strconv.Atoi(options[j+1])

			switch operator {
			case 0:
				total += next
			case 1:
				total *= next
			}
		}

		if total == goal {
			return goal
		}
	}

	return 0
}

func processLinePart2(line string) int {
	lineParts := strings.Split(line, ": ")
	goal, _ := strconv.Atoi(lineParts[0])
	options := strings.Split(lineParts[1], " ")
	possibilities := math.Pow(3, float64(len(options)-1))

	// fmt.Println(goal)

	for i := 0; i < int(possibilities); i++ {
		total, _ := strconv.Atoi(options[0])

		for j := 0; j < len(options)-1; j++ {
			operator := (i / int(math.Pow(3, float64(j)))) % 3
			next, _ := strconv.Atoi(options[j+1])

			switch operator {
			case 0:
				total += next
			case 1:
				total *= next
			case 2:
				total, _ = strconv.Atoi(strconv.Itoa(total) + options[j+1])
			}
		}

		if total == goal {
			return goal
		}
	}

	return 0
}

func processLinePart22(options []int, goal int) int {
	possibilities := math.Pow(2, float64(len(options)-1))

	for i := 0; i < int(possibilities); i++ {
		total := options[0]

		for j := 0; j < len(options)-1; j++ {
			operator := (i >> j) & 1
			next := options[j+1]

			switch operator {
			case 0:
				total += next
			case 1:
				total *= next
			}
		}

		if total == goal {
			fmt.Println(goal, options)
			return goal
		}
	}

	fmt.Println(goal, possibilities, options)

	return 0
}
