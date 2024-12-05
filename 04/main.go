package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 4")
	input, inputErr := os.ReadFile("input")

	if inputErr != nil {
		panic(inputErr)
	}

	xmasCounter := 0
	inputString := string(input)
	xMap := make([][]int, 0)
	xmasMap := make([][][]int, 0)
	inputSlice := strings.Split(inputString, "\n")

	for lineIdx, line := range inputSlice {
		for charIndex, char := range line {
			stringChar := string(char)

			if stringChar == "X" {
				xMap = append(xMap, []int{lineIdx, charIndex})
			}
		}
	}

	for _, location := range xMap {
		for i := 1; i <= 8; i++ {
			found := make([][]int, 0)

			foundM, locationM, directionM := findLetterNextTo("M", location, inputSlice, i)
			foundA, locationA, _ := findLetterNextTo("A", locationM, inputSlice, directionM)
			foundS, locationS, _ := findLetterNextTo("S", locationA, inputSlice, directionM)

			if foundM && foundA && foundS {
				found = append(found, location)
				found = append(found, locationM)
				found = append(found, locationA)
				found = append(found, locationS)
			}

			if len(found) > 0 {
				xmasMap = append(xmasMap, found)
			}
		}

	}

	fmt.Println("Part 1: ", xmasCounter, len(xmasMap))

	aMap := make([][]int, 0)
	xMasMap := make([][]int, 0)

	for lineIdx, line := range inputSlice {
		if lineIdx == 0 || len(inputSlice)-1 == lineIdx {
			continue
		}

		for charIndex, char := range line {
			stringChar := string(char)

			if stringChar == "A" {
				aMap = append(aMap, []int{lineIdx, charIndex})
			}
		}
	}

	for _, location := range aMap {
		found := make([][]int, 0)

		for i := 2; i <= 8; i += 2 {
			foundS, locationS, directionS := findLetterNextTo("S", location, inputSlice, i)

			if foundS {
				directionOpposite, directionOppositeErr := getOppositeCorner(directionS)

				if directionOppositeErr != nil {
					panic(directionOppositeErr)
				}

				foundM, locationM, _ := findLetterNextTo("M", location, inputSlice, directionOpposite)

				if foundM {
					found = append(found, locationS)
					found = append(found, locationM)
				}
			}
		}

		if len(found) == 4 {
			xMasMap = append(xMasMap, location)
		}

	}

	fmt.Println("Part 2: ", len(xMasMap))
}

func findLetterNextTo(letter string, location []int, inputSlice []string, direction int) (bool, []int, int) {
	if location[0] < 0 || location[0] >= len(inputSlice) ||
		location[1] < 0 || location[1] >= len(inputSlice[0]) {
		return false, []int{0, 0}, -1
	}

	currentLine := inputSlice[location[0]]

	if currentLine == "" {
		return false, []int{0, 0}, 0
	}

	hasPrevLine := location[0]-1 >= 0
	hasNextLine := location[0]+1 < len(inputSlice)-1
	hasNextCol := location[1]+1 < len(currentLine)
	hasPrevCol := location[1]-1 >= 0

	if direction == 0 {
		return checkAllDirections(letter, location, inputSlice)
	}

	switch direction {
	case 1: // Right
		if hasNextCol && string(currentLine[location[1]+1]) == letter {
			return true, []int{location[0], location[1] + 1}, getDirection(direction, 1)
		}
	case 2: // Bottom Right
		if hasNextLine && hasNextCol && string(inputSlice[location[0]+1][location[1]+1]) == letter {
			return true, []int{location[0] + 1, location[1] + 1}, getDirection(direction, 2)
		}
	case 3: // Bottom
		if hasNextLine && string(inputSlice[location[0]+1][location[1]]) == letter {
			return true, []int{location[0] + 1, location[1]}, getDirection(direction, 3)
		}
	case 4: // Bottom Left
		if hasNextLine && hasPrevCol && string(inputSlice[location[0]+1][location[1]-1]) == letter {
			return true, []int{location[0] + 1, location[1] - 1}, getDirection(direction, 4)
		}
	case 5: // Left
		if hasPrevCol && string(currentLine[location[1]-1]) == letter {
			return true, []int{location[0], location[1] - 1}, getDirection(direction, 5)
		}
	case 6: // Top Left
		if hasPrevLine && hasPrevCol && string(inputSlice[location[0]-1][location[1]-1]) == letter {
			return true, []int{location[0] - 1, location[1] - 1}, getDirection(direction, 6)
		}
	case 7: // Top
		if hasPrevLine && string(inputSlice[location[0]-1][location[1]]) == letter {
			return true, []int{location[0] - 1, location[1]}, getDirection(direction, 7)
		}
	case 8: // Top Right
		if hasPrevLine && hasNextCol && string(inputSlice[location[0]-1][location[1]+1]) == letter {
			return true, []int{location[0] - 1, location[1] + 1}, getDirection(direction, 8)
		}
	}

	return false, []int{0, 0}, -1
}

func checkAllDirections(letter string, location []int, inputSlice []string) (bool, []int, int) {
	for dir := 1; dir <= 8; dir++ {
		found, newLoc, nextDir := findLetterNextTo(letter, location, inputSlice, dir)
		if found {
			return found, newLoc, nextDir
		}
	}
	return false, []int{0, 0}, -1
}

func getDirection(direction int, fallback int) int {
	if direction != -1 {
		return direction
	}

	return fallback
}

func getOppositeCorner(corner int) (int, error) {
	opposites := map[int]int{
		2: 6,
		4: 8,
		6: 2,
		8: 4,
	}

	if opposite, exists := opposites[corner]; exists {
		return opposite, nil
	}

	return 0, fmt.Errorf("corner %d does not exist", corner)
}
