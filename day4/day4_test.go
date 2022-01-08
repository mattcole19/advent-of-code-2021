package main

import (
	"testing"
)

func TestDay4Part1(t *testing.T) {
	want := 4512
	actual := Part1("input/example")
	if want != actual {
		t.Fatalf("Wrong answer for part1. Got %d but expected %d", actual, want)
	}
}

func TestDay4Part2(t *testing.T) {
	want := 1924
	actual := Part2("input/example")
	if want != actual {
		t.Fatalf("Wrong answer for part2. Got %d but expected %d", actual, want)
	}
}
