package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	fmt.Print("Starting day 0")

	fmt.Printf("\nAnswer for part1 = %d\n", Part1("input/data"))

}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	// Create some sort of map/counter
	oneCounter := make(map[int]int)
	zeroCounter := make(map[int]int)
	scanner := bufio.NewScanner(file)
	var lineLength int
	for scanner.Scan() {
		line := scanner.Text()
		lineLength = len(line)
		for i, bit := range line {
			if string(bit) == "1" {
				oneCounter[i] += 1
			} else {
				zeroCounter[i] += 1
			}
		}
	}

	// Create gamma string and epsilon string
	var gammaString strings.Builder
	var epsilonString strings.Builder
	for i := 0; i < lineLength; i++ {
		if oneCounter[i] > zeroCounter[i] {
			gammaString.WriteString("1")
			epsilonString.WriteString("0")
		} else {
			gammaString.WriteString("0")
			epsilonString.WriteString("1")
		}
	}

	// fmt.Print("\n")
	// fmt.Printf(gammaString.String())
	// fmt.Print("\n")
	// fmt.Printf(epsilonString.String())
	// fmt.Print("\n")
	// Convert gamma string and epsilon string to int
	gamma := binaryStringToInt(gammaString.String())
	epsilon := binaryStringToInt(epsilonString.String())

	return gamma * epsilon

}

func Part2(filePath string) int {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	// Create some sort of map/counter
	oneCounter := make(map[int]int)
	zeroCounter := make(map[int]int)
	oxyValues := make([]string, 0)
	co2Values := make([]string, 0)
	scanner := bufio.NewScanner(file)
	var lineLength int
	for scanner.Scan() {
		line := scanner.Text()
		oxyValues = append(oxyValues, line)
		co2Values = append(co2Values, line)
		lineLength = len(line)
		for i, bit := range line {
			if string(bit) == "1" {
				oneCounter[i] += 1
			} else {
				zeroCounter[i] += 1
			}
		}
	}
	for i := lineLength - 1; i >= 0; i-- {

		// If only one number remains, stop

		// Filter out at index i
		if oneCounter[i] >= zeroCounter[i] {
			oxyValues = filterSlice(oxyValues, i, "1")
			co2Values = filterSlice(co2Values, i, "0")
		} else {
			oxyValues = filterSlice(oxyValues, i, "0")
			co2Values = filterSlice(co2Values, i, "1")
		}
	}

	return binaryStringToInt(co2Values[0]) * binaryStringToInt(oxyValues[0])
}

func binaryStringToInt(binaryString string) int {
	binaryStringSlice := strings.Split(binaryString, "")

	counter := 0
	result := 0
	for i := len(binaryStringSlice) - 1; i >= 0; i-- {
		if binaryStringSlice[i] == "1" {
			result += int(1 * math.Pow(2, float64(counter)))
		}
		counter += 1
	}

	return result
}

// Takes in slice of strings and removes any values that don't match criteria
func filterSlice(values []string, index int, criteria string) []string {

	// If theres only one element, no need to filter
	if len(values) == 1 {
		return values
	}

	j := 0
	for i, value := range values {
		fmt.Printf("Index %d = %s", i, value)

		// Split value into characters
		valueSlice := strings.Split(value, "")

		if valueSlice[index] != criteria {
			// remove value from values
			values = removeElementFromSlice(values, j)
		} else {
			j += 1
		}

	}

	// Returns list with empty elements
	return values
}

func removeElementFromSlice(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
