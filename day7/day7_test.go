package main

import (
	"testing"
)

func TestDay7Part1(t *testing.T) {
	want := 37
	actual := Part1("input/example")
	if want != actual {
		t.Fatalf("Wrong answer. Got %d but expected %d", actual, want)
	}
}

func TestDay1Part2(t *testing.T) {
	want := 168
	actual := Part2("input/example")
	if want != actual {
		t.Fatalf("Wrong answer for part2. Got %d but expected %d", actual, want)
	}
}
