package main

import (
	"fmt"
	"os"
	"strings"
)

type field struct {
	x, y         int
	object, step bool
}

type guard struct {
	x, y, direction int
}

func main() {
	fmt.Println("Day 6")

	input, inputErr := os.ReadFile("input")

	if inputErr != nil {
		panic(inputErr)
	}

	mappedArea := make([][]field, 0)
	firstRun := true
	var startingPoint guard
	var newPosition guard
	var currentPosition guard

	for x, i := range strings.Split(string(input), "\n") {
		if i == "" {
			continue
		}

		items := make([]field, 0)

		for y, z := range strings.Split(i, "") {
			items = append(items, field{x, y, z == "#", false})

			if z == "^" {
				startingPoint = guard{x, y, 1}
			}
		}

		mappedArea = append(mappedArea, items)
	}

	totalMoves := 0

	mapCopy := deepCopyField(mappedArea)

	for {
		if firstRun {
			newPosition = getNewPosition(startingPoint)
			currentPosition = newPosition
			firstRun = false
		} else {
			newPosition = getNewPosition(newPosition)
		}

		if isPositionOutsideMap(mappedArea, newPosition) {
			break
		}

		if isObjectInTheWay(mappedArea, newPosition) {
			newPosition = getNewDirection(currentPosition)
			continue
		}

		currentPosition = newPosition
		mappedArea[newPosition.x][newPosition.y].step = true
	}

	for _, x := range mappedArea {
		for _, xx := range x {
			if xx.step {
				totalMoves++

			}
		}
	}

	fmt.Println("Part 1 ", totalMoves)

	println()

	fmt.Println("Part 2")

	totalObstacles := 0

	for y, line := range mappedArea {
		for x, item := range line {
			if item.y == startingPoint.y && item.x == startingPoint.x {
				continue
			}

			if item.object {
				continue
			}

			if item.step == false {
				continue
			}

			mapCopyNew := deepCopyField(mapCopy)
			mapCopyNew[y][x].object = true

			if tryMap(startingPoint, mapCopyNew) {
				fmt.Println("added object on ", y, x)

				totalObstacles++
			}
		}
	}

	fmt.Println("Part 2 ", totalObstacles)
}

func deepCopyField(original [][]field) [][]field {
	if len(original) == 0 {
		return nil
	}

	copy := make([][]field, len(original))
	for i := range original {
		copy[i] = make([]field, len(original[i]))
		for j := range original[i] {
			copy[i][j] = field{
				x:      original[i][j].x,
				y:      original[i][j].y,
				object: original[i][j].object,
				step:   original[i][j].step,
			}
		}
	}
	return copy
}

func tryMap(startingPoint guard, mappedArea [][]field) bool {
	var currentPosition guard
	var newPosition guard

	firstRun := true

	rounds := 0

	for {
		rounds++

		if rounds > 25000 {
			return true
		}

		if firstRun {
			newPosition = getNewPosition(startingPoint)
			firstRun = false
		} else {
			newPosition = getNewPosition(newPosition)
		}

		if isPositionOutsideMap(mappedArea, newPosition) {
			return false
		}

		if isObjectInTheWay(mappedArea, newPosition) {
			if currentPosition.direction == 0 {
				currentPosition = newPosition
			}

			newPosition = getNewDirection(currentPosition)
			currentPosition = newPosition
			continue
		}

		currentPosition = newPosition

		//	if mappedArea[newPosition.x][newPosition.y].step {
		//		return true
		//	}

		mappedArea[newPosition.x][newPosition.y].step = true

	}

	return false
}

func cleanSteps(ma [][]field) [][]field {
	for y, line := range ma {
		for x := range line {
			ma[y][x].step = false
		}
	}

	return ma
}

func getNewPosition(pos guard) guard {
	if pos.direction == 1 {
		return guard{pos.x - 1, pos.y, pos.direction}
	}

	if pos.direction == 2 {
		return guard{pos.x, pos.y + 1, pos.direction}
	}

	if pos.direction == 3 {
		return guard{pos.x + 1, pos.y, pos.direction}
	}

	if pos.direction == 4 {
		return guard{pos.x, pos.y - 1, pos.direction}
	}

	panic("Unknown Direction")
}

func isPositionOutsideMap(mappedArea [][]field, pos guard) bool {
	if pos.x < 0 || pos.y < 0 {
		return true
	}

	if len(mappedArea)-1 >= pos.x && len(mappedArea[pos.x])-1 >= pos.y {
		return false
	}

	return true
}

func isObjectInTheWay(mappedArea [][]field, pos guard) bool {
	line := mappedArea[pos.x]

	return line[pos.y].object
}

func getNewDirection(pos guard) guard {
	direction := pos.direction

	if direction == 4 {
		pos.direction = 1
		return pos
	}
	if direction == 3 {
		pos.direction = 4
		return pos
	}
	if direction == 2 {
		pos.direction = 3
		return pos
	}
	if direction == 1 {
		pos.direction = 2
		return pos
	}

	panic("No Direction")
}

func paintMap(ma [][]field) {
	println()
	for _, line := range ma {
		for _, item := range line {
			if item.step {
				fmt.Print("X")
				continue
			}

			if item.object {
				fmt.Print("#")
				continue
			}

			fmt.Print(".")
		}

		println()
	}
	println()
}
