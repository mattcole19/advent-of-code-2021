package main

import (
	"testing"
)

func TestDay2Part19(t *testing.T) {
	want := 150
	actual := Part1("input/example")
	if want != actual {
		t.Fatalf("Wrong answer for part1. Got %d but expected %d", actual, want)
	}
}

func TestDay2Part2(t *testing.T) {

}
