package main

import (
	"bufio"
	"fmt"
	"math"
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
	fmt.Print("Starting day 7")
	answer := Part2("input/data")
	fmt.Printf("\n Answer for part2= %d\n", answer)

}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	// numCrabs := 0
	crabFuels := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		numStrings := strings.Split(line, ",")
		// numCrabs = len(numStrings)
		for _, numString := range numStrings {
			num, err := strconv.Atoi(numString)
			crabFuels = append(crabFuels, num)
			check(err)
			total += num
		}
	}

	// Major improvements could be made here!
	// Slow way: Iterate from min, max calulating how much fuel it would take
	// Get min and max
	minPosition := min(crabFuels)
	maxPosition := max(crabFuels)

	fuelCosts := make([]int, 0)
	for position := minPosition; position < maxPosition; position++ {
		fuelCosts = append(fuelCosts, calculateFuel(crabFuels, position))
	}

	minFuelCost := min(fuelCosts)

	return minFuelCost
}

// Given a slice of positions, calculate how much fuel it would cost for all of them to get to endPosition
func calculateFuel(positions []int, endPosition int) int {
	totalFuelCost := 0
	for _, position := range positions {
		totalFuelCost += int(math.Abs(float64(endPosition - position)))
	}
	return totalFuelCost
}

func min(s []int) int {
	minElement := s[0]
	for _, element := range s {
		if element < minElement {
			minElement = element
		}
	}
	return minElement
}

func max(s []int) int {
	maxElement := s[0]
	for _, element := range s {
		if element > maxElement {
			maxElement = element
		}
	}
	return maxElement
}

func Part2(filePath string) int {

	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	crabFuels := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		numStrings := strings.Split(line, ",")
		for _, numString := range numStrings {
			num, err := strconv.Atoi(numString)
			crabFuels = append(crabFuels, num)
			check(err)
			total += num
		}
	}

	// Major improvements could be made here!
	// Slow way: Iterate from min, max calulating how much fuel it would take
	// Get min and max
	minPosition := min(crabFuels)
	maxPosition := max(crabFuels)

	fuelCosts := make([]int, 0)

	// new fuel cost calulation
	for position := minPosition; position < maxPosition; position++ {
		fuelCosts = append(fuelCosts, newFuelCost(crabFuels, position))
	}

	minFuelCost := min(fuelCosts)

	return minFuelCost

}

func newFuelCost(s []int, endPosition int) int {
	totalFuelCost := 0
	for _, position := range s {
		fuelCost := calculateNonConstantFuelCost(position, endPosition)
		totalFuelCost += fuelCost
	}

	return totalFuelCost
}

//  first step costs 1, second costs 2, third costs 3, and so on...
func calculateNonConstantFuelCost(start int, end int) int {
	return additionFactorial(int(math.Abs(float64(end - start))))
}

// Formula (n^2 + n) / (2)
func additionFactorial(n int) int {
	return int((math.Pow(float64(n), 2) + float64(n)) / 2)
}
