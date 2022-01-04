package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("Starting day1 of Advent of code\n")

	answer := Part2("input/data")
	fmt.Print(answer)

}

func Part1(filePath string) int {

	// Open a file and defer closing it
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	// Read file contents
	numIncreases := 0
	scanner := bufio.NewScanner(file)
	var prevDepth int
	firstScan := true

	for scanner.Scan() {

		// First scan doesn't have a comparison and should be skipped
		if firstScan {
			prevDepth, err = strconv.Atoi(scanner.Text())
			check(err)
			firstScan = false
			continue
		}

		curDepth, err := strconv.Atoi(scanner.Text())
		check(err)

		if curDepth > prevDepth {
			numIncreases += 1

			// fmt.Printf("Curdepth %d is greater than prevdepth %d\n", curDepth, prevDepth)
		}

		prevDepth = curDepth
	}

	return numIncreases
}

func Part2(filePath string) int {

	// Open a file and defer closing it
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	// Read file contents and create an array
	numIncreases := 0
	scanner := bufio.NewScanner(file)
	depths := make([]int, 0)
	for scanner.Scan() {
		curDepth, err := strconv.Atoi(scanner.Text())
		check(err)
		depths = append(depths, curDepth)
	}

	// a = ave(depths[0, 3])
	// b = ave(depths[1, 4])
	// c = ave(depths[2, 5])
	// ...
	// Z = ave(depths[])

	// Iterate over array
	start := 0
	end := 3
	avePrevDepth := 10000000.0 // Hard coded for now..
	for {
		// Break out if end == len(depths)
		if end-1 == len(depths) {
			// fmt.Printf("Last slice is from %d to %d", start-1, end-1)
			break
		}

		aveCurrDepth := average(depths[start:end])
		if aveCurrDepth > avePrevDepth {
			numIncreases += 1
		}

		start += 1
		end += 1
		avePrevDepth = aveCurrDepth
	}

	return numIncreases

}

// Method to take in slice of ints and return average
func average(nums []int) float64 {
	total := 0.0
	for _, value := range nums {
		total += float64(value)
	}

	return total / float64(len(nums))
}
