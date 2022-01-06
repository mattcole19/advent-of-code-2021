package main

import (
	"testing"
)

func TestDay3Part1(t *testing.T) {
	want := 198
	actual := Part1("input/example")
	if want != actual {
		t.Fatalf("Wrong answer for part1. Got %d but expected %d", actual, want)
	}
}

func TestDay3Part2(t *testing.T) {
	want := 230
	actual := Part2("input/example")
	if want != actual {
		t.Fatalf("Wrong answer for part2. Got %d but expected %d", actual, want)
	}
}

func TestDay3BinaryStringToInt(t *testing.T) {
	want := 22
	actual := binaryStringToInt("10110")
	if want != actual {
		t.Fatalf("Wrong answer for binaryStringToInt. Got %d but expected %d", actual, want)
	}
}

func TestDay3BinaryStringToInt2(t *testing.T) {
	want := 9
	actual := binaryStringToInt("01001")
	if want != actual {
		t.Fatalf("Wrong answer for binaryStringToInt. Got %d but expected %d", actual, want)
	}
}
