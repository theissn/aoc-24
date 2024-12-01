package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 1")

	left := make([]int, 0)
	right := make([]int, 0)

	final := 0

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		string := strings.Split(s.Text(), "   ")

		leftInt, leftErr := strconv.Atoi(string[0])

		if leftErr != nil {
			panic(leftErr)
		}

		left = append(left, leftInt)

		rightInt, rightErr := strconv.Atoi(string[1])

		if rightErr != nil {
			panic(rightErr)
		}

		right = append(right, rightInt)
	}

	sort.Ints(left)
	sort.Ints(right)

	for idx, i := range left {
		fmt.Println(idx, i, right[idx])

		toAdd := i - right[idx]

		if toAdd < 0 {
			toAdd = -toAdd
		}

		final = final + toAdd
	}

	fmt.Println("Result Part 1:")
	fmt.Println(final)

	finalPart2 := 0

	for _, i := range left {
		toMultipy := 0

		for _, y := range right {
			if i == y {
				toMultipy++
			}
		}

		finalPart2 += i * toMultipy
	}

	fmt.Println("Result part 2")
	fmt.Println(finalPart2)
}
