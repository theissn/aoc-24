package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//	sampleInput := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

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

	//	re2 := regexp.MustCompile("(?<!don't\(\))mul\([0-9]{1,3},[0-9]{1,3}\)")
	regex2 := re.FindAllString(string(input), -1)
	total2 := 0

	//	doRe := regexp.MustCompile(`do\\(\\)`)
	//	dontRe := regexp.MustCompile(`don\\'t`)

	for _, mul := range regex2 {
		//		print(mul)
		//index := strings.Index(sampleInput, mul)
		index := strings.Index(string(input), mul)
		textBefore := string(input)[:index]

		dosBefore := strings.LastIndex(textBefore, `do()`)
		dontsBefore := strings.LastIndex(textBefore, `don't()`)

		if dosBefore < dontsBefore {
			continue
		}

		//fmt.Println(mul, index, textBefore, dosBefore, dontsBefore)

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
