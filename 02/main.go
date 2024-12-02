package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 2")

	data := make([][]string, 0)

	final := 0
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		string := strings.Split(s.Text(), " ")
		data = append(data, string)
	}

	for _, a := range data {
		toAdd, _ := checkArray(a)
		final += toAdd
	}

	fmt.Println("Result Part 1:")
	fmt.Println(final)

	finalPart2 := 0

	for _, b := range data {
		var toAdd int
		toAdd, _ = checkArray(b)

		if toAdd == 0 {
			rounds := len(b)

			for i := 0; i < rounds; i++ {
				c := RemoveIndex(b, i)
				toAdd, _ = checkArray(c)

				if toAdd == 1 {
					finalPart2 += 1
					break
				}
			}

		} else {
			finalPart2 += 1
		}

	}

	fmt.Println("Result part 2")
	fmt.Println(finalPart2)
}

func RemoveIndex(s []string, index int) []string {
	bCopy := make([]string, len(s))
	copy(bCopy, s)
	return append(bCopy[:index], bCopy[index+1:]...)
}

func checkArray(d []string) (int, int) {
	var direction int

	for x, i := range d {
		//		fmt.Println(i)
		if x == 0 {
			continue
		}

		prev, prevErr := strconv.Atoi(d[x-1])
		current, currentErr := strconv.Atoi(i)

		if currentErr != nil || prevErr != nil {
			panic(fmt.Sprintf("currenterr: %v, nexterr: %v", currentErr, prevErr))
		}

		if x == 1 {
			if prev < current {
				direction = 1
			} else {
				direction = -1
			}
		}

		diff := current - prev

		if diff < 0 {
			diff = -diff
		}

		if diff > 3 || diff < 1 {
			return 0, x
		}

		if direction == 1 && current < prev {
			return 0, x
		} else if direction == -1 && current > prev {
			return 0, x
		}
	}

	return 1, -1
}
