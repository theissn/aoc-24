package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 5")

	invalidLines := make([][]int, 0)
	input, inputErr := os.ReadFile("input")

	if inputErr != nil {
		panic(inputErr)
	}

	inputParts := strings.Split(string(input), "\n\n")
	total := 0

	fmt.Println(inputParts)
	fmt.Println()

	parsedOrder := make([][]int, 0)

	for _, set := range strings.Split(inputParts[0], "\n") {
		itemSet := make([]int, 0)

		for _, item := range strings.Split(set, "|") {
			value := stringToInt(item)

			itemSet = append(itemSet, value)
		}

		parsedOrder = append(parsedOrder, itemSet)
	}

	fmt.Println(parsedOrder)

	for _, line := range strings.Split(inputParts[1], "\n") {
		if line == "" {
			continue
		}

		itemsInLine := strings.Split(line, ",")
		newLine, lineErr := isLineValid(itemsInLine, parsedOrder)

		if lineErr != nil {
			fmt.Println("Line not valid: ", line)
			invalidLines = append(invalidLines, parseStringItemsToInt(itemsInLine))
			continue
		}

		fmt.Println("Line valid: ", line)

		total += newLine[len(newLine)/2]
	}

	fmt.Println("Part 1 ", total)

	totalInvalid := 0

	println("Part 2")

	for _, line := range invalidLines {
		newLine := fixLineOrder(line, parsedOrder)

		totalInvalid += newLine[len(newLine)/2]
	}

	fmt.Println("Part 2 ", totalInvalid)
}

func parseStringItemsToInt(line []string) []int {
	items := make([]int, 0)

	for _, i := range line {
		items = append(items, stringToInt(i))
	}

	return items
}

func isLineValid(itemsInLine []string, parsedOrder [][]int) ([]int, error) {
	items := make([]int, 0)

	for i, input := range itemsInLine {
		item := stringToInt(input)
		items = append(items, item)

		if i == len(itemsInLine)-1 {
			break
		}

		next := stringToInt(itemsInLine[i+1])

		if numberShouldGoBefore(item, next, parsedOrder) == false {
			return items, errors.New("Items in wrong order")
		}

	}

	return items, nil
}

func fixLineOrder(line []int, parsedOrder [][]int) []int {
	newLine := make([]int, 0)

	for i, input := range line {
		if i == len(line)-1 {
			newLine = append(newLine, input)
			break
		}

		next := line[i+1]

		if numberShouldGoBefore(input, next, parsedOrder) {
			newLine = append(newLine, input)
			continue
		}

		if numberShouldGoBefore(next, input, parsedOrder) {
			newLine = append(newLine, next)

			newNewLine := make([]int, 0)
			newNewLine = append(newNewLine, newLine...)
			newNewLine = append(newNewLine, input)
			newNewLine = append(newNewLine, line[i+2:]...)

			return fixLineOrder(newNewLine, parsedOrder)
		}
	}

	return newLine
}

func numberShouldGoBefore(before int, number int, parsedOrder [][]int) bool {
	for _, itemSet := range parsedOrder {

		if itemSet[0] == number && itemSet[1] == before {
			return false
		}
	}

	return true
}

func stringToInt(input string) int {
	value, valueErr := strconv.Atoi(input)

	if valueErr != nil {
		panic(valueErr)
	}

	return value
}
