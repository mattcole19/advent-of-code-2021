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
	fmt.Print("Starting day 0\n")

	answer := Part1("input/data")

	fmt.Print(answer)

}

func Part1(filePath string) int {

	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	curDepth := 0
	curDistance := 0
	for scanner.Scan() {
		line := scanner.Text()
		dir, mag := splitInput(line)
		if dir == "forward" {
			curDistance += mag
		} else if dir == "down" {
			curDepth += mag
		} else {
			curDepth -= mag
		}
	}

	return curDepth * curDistance

}

func Part2(filePath string) {

}

func splitInput(line string) (string, int) {
	s := strings.Split(line, " ")
	direction := s[0]
	magnitude, err := strconv.Atoi(s[1])
	check(err)

	return direction, magnitude
}
