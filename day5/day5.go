package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	fmt.Print("Starting day 5")

	answer := Part2("input/data")
	fmt.Printf("Answer = %d", answer)

}

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type Point struct {
	x int
	y int
}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	counter := make(map[Point]int)
	numSpotsToAvoid := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		// Split input into x1, y1, x2, y2
		line := splitLine(scanner.Text())
		pointsCrossed := line.getPoints()

		for _, point := range pointsCrossed {
			fmt.Printf("\nadding to count of point %d. It was = %d\n", point, counter[point])
			counter[point] += 1

			if counter[point] == 2 {
				fmt.Printf("got here")
				numSpotsToAvoid += 1
			}
		}
		fmt.Printf("line = %d", line)
	}
	return numSpotsToAvoid

}

func Part2(filePath string) int {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	counter := make(map[Point]int)
	numSpotsToAvoid := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		// Split input into x1, y1, x2, y2
		line := splitLine(scanner.Text())
		pointsCrossed := line.getPointsNoIgnore()

		for _, point := range pointsCrossed {
			fmt.Printf("\nadding to count of point %d. It was = %d\n", point, counter[point])
			counter[point] += 1

			if counter[point] == 2 {
				fmt.Printf("got here")
				numSpotsToAvoid += 1
			}
		}
		fmt.Printf("line = %d", line)
	}
	return numSpotsToAvoid
}

// Splits input 'x1, y1 -> x2, y2' into a Line{x1, y1, x2, y2}
func splitLine(line string) Line {
	// Split on " -> " first
	initialSplit := strings.Split(line, " -> ")

	// Then Split on "->"
	firstPoint := strings.Split(initialSplit[0], ",")
	secondPoint := strings.Split(initialSplit[1], ",")

	x1, err := strconv.Atoi(firstPoint[0])
	check(err)

	y1, err := strconv.Atoi(firstPoint[1])
	check(err)

	x2, err := strconv.Atoi(secondPoint[0])
	check(err)

	y2, err := strconv.Atoi(secondPoint[1])
	check(err)

	return Line{x1, y1, x2, y2}
}

// Gets all points the line goes through
func (line Line) getPoints() []Point {

	// Only consider lines where x1 = x2 OR y1 = y2
	points := make([]Point, 0)

	// If x1 = x2, iterate from max(y1,y2) -> min(y1, y2)
	if line.x1 == line.x2 {

		for yStart := min(line.y1, line.y2); yStart <= max(line.y1, line.y2); yStart++ {
			point := Point{x: line.x1, y: yStart}
			points = append(points, point)
		}
	} else if line.y1 == line.y2 {
		// If y1 = y2, iterate from x1 -> x2
		for xStart := min(line.x1, line.x2); xStart <= max(line.x1, line.x2); xStart++ {
			point := Point{x: xStart, y: line.y1}
			points = append(points, point)
		}

	} else {
		fmt.Printf("\nLine %d does not have x1=x2 OR y1=y2. Skipping. \n", line)
	}

	return points
}

// Gets all points the line goes through
func (line Line) getPointsNoIgnore() []Point {
	points := make([]Point, 0)
	// If x1 = x2, iterate from max(y1,y2) -> min(y1, y2)
	if line.x1 == line.x2 {

		for yStart := min(line.y1, line.y2); yStart <= max(line.y1, line.y2); yStart++ {
			point := Point{x: line.x1, y: yStart}
			points = append(points, point)
		}
	} else if line.y1 == line.y2 {
		// If y1 = y2, iterate from x1 -> x2
		for xStart := min(line.x1, line.x2); xStart <= max(line.x1, line.x2); xStart++ {
			point := Point{x: xStart, y: line.y1}
			points = append(points, point)
		}

	} else {
		// Get the diagonals
		fmt.Printf("\nLine %d does not have x1=x2 OR y1=y2. getting diagonals. \n", line)

		points = append(points, Point{line.x1, line.y1})
		points = append(points, Point{line.x2, line.y2})
		deltaX := 0
		deltaY := 0
		if line.x1 > line.x2 {
			deltaX = -1
		} else {
			deltaX = 1
		}

		if line.y1 > line.y2 {
			deltaY = -1
		} else {
			deltaY = 1
		}

		// Start from x1, y1. Go till x2, y2 using deltas
		xStart := line.x1 + deltaX
		yStart := line.y1 + deltaY
		for {
			if xStart == line.x2 || yStart == line.y2 {
				break
			}

			point := Point{xStart, yStart}
			points = append(points, point)
			fmt.Printf("Added point %d\n", point)

			xStart += deltaX
			yStart += deltaY

		}

	}

	return points
}

func max(val1 int, val2 int) int {
	if val1 >= val2 {
		return val1
	}
	return val2
}

func min(val1 int, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}
