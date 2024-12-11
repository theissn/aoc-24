package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 6")

	input, _ := os.ReadFile("input-sample")
	inputString := string(input)
	sum := 0
	sum2 := 0

	for _, line := range strings.Split(inputString, "\n") {
		if line == "" {
			continue
		}

		sum += processLine(line)
		sum2 += processLinePart2(line)
	}

	fmt.Println("Part 1", sum)
	fmt.Println("Part 2", sum2)
}

func processLine(line string) int {
	parts := strings.Split(line, ": ")
	goal, _ := strconv.Atoi(parts[0])
	items := strings.Split(parts[1], " ")
	possibilities := powInt(2, len(items)-1)

	for i := 0; i < possibilities; i++ {
		total := 0

		for j := 0; j < len(items)-1; j++ {
			operator := (i >> j) & 1

			if total == 0 {
				if operator == 0 {
					//					fmt.Println(items[j], "+", strToInt(items[j+1]))
					total += (strToInt(items[j]) + strToInt(items[j+1]))

					continue
				}

				if operator == 1 {
					//					fmt.Println(items[j], "*", strToInt(items[j+1]))
					total += (strToInt(items[j]) * strToInt(items[j+1]))
					continue
				}
				continue
			}

			if operator == 0 {
				//				fmt.Println(total, "+", strToInt(items[j+1]))
				total = (total + strToInt(items[j+1]))

				continue
			}

			if operator == 1 {
				//				fmt.Println(total, "*", strToInt(items[j+1]))
				total = (total * strToInt(items[j+1]))
				continue
			}
		}

		//		fmt.Println(items, total, goal)

		if total == goal {
			return total
		}
	}

	return 0
}

func strToInt(value string) int {
	val, _ := strconv.Atoi(value)
	return val
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func processLinePart2(line string) int {
	parts := strings.Split(line, ": ")
	goal, _ := strconv.Atoi(parts[0])
	items := strings.Split(parts[1], " ")
	possibilities := powInt(3, len(items)-1)

	fmt.Println(items, goal)

	for i := 0; i < possibilities; i++ {
		total := 0

		for j := 0; j < len(items)-1; j++ {
			//operator := i % 3
			operator := (i / powInt(3, j)) % 3

			fmt.Print(operator)
		}

		println()

		if total == goal {
			fmt.Println(items, goal)
			return total
		}
	}

	return 0
	// panic(1)
}
