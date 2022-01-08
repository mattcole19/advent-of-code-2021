package main

import (
	"testing"
)

func TestDay5Part1(t *testing.T) {
	want := 5
	actual := Part1("input/example")
	if want != actual {
		t.Fatalf("Wrong answer. Got %d but expected %d", actual, want)
	}
}

func TestDay5Part2(t *testing.T) {
	want := 12
	actual := Part2("input/example")
	if want != actual {
		t.Fatalf("Wrong answer for part2. Got %d but expected %d", actual, want)
	}
}
