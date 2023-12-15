package main

import (
	"testing"
)

func TestQ1_test(t *testing.T) {
	l, rowBounds, colBounds := parse("dec14/data/q1_test.txt")

	have := q1(l, rowBounds, colBounds)
	want := 136

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestQ2_test(t *testing.T) {
	l, rowBounds, colBounds := parse("dec14/data/q2_test.txt")

	have := q2(l, rowBounds, colBounds)
	want := 64

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}
