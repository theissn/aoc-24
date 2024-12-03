package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, inputErr := os.ReadFile("input")

	if inputErr != nil {
		panic(inputErr)
	}

	fmt.Println("Day 3")

	re := regexp.MustCompile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	sampleOut := re.FindAllString(string(input), -1)
	total := 0

	for _, mul := range sampleOut {
		toMultiply := strings.Replace(strings.Replace(string(mul), "mul(", "", 1), ")", "", -1)

		parts := strings.Split(toMultiply, ",")

		left, leftErr := strconv.Atoi(parts[0])
		right, rightErr := strconv.Atoi(parts[1])

		if leftErr != nil || rightErr != nil {
			panic("Error")
		}

		total += (left * right)
	}

	fmt.Println("Part 1")
	fmt.Println(total)

	regex2 := re.FindAllString(string(input), -1)
	total2 := 0

	for _, mul := range regex2 {
		index := strings.Index(string(input), mul)
		textBefore := string(input)[:index]

		dosBefore := strings.LastIndex(textBefore, `do()`)
		dontsBefore := strings.LastIndex(textBefore, `don't()`)

		if dosBefore < dontsBefore {
			continue
		}

		toMultiply := strings.Replace(strings.Replace(string(mul), "mul(", "", 1), ")", "", -1)

		parts := strings.Split(toMultiply, ",")

		left, leftErr := strconv.Atoi(parts[0])
		right, rightErr := strconv.Atoi(parts[1])

		if leftErr != nil || rightErr != nil {
			panic("Error")
		}

		total2 += (left * right)
	}

	fmt.Println("Part 2")
	fmt.Println(total2)
}
