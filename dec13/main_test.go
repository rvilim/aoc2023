package main

import (
	"testing"
)

func TestQ1_test(t *testing.T) {
	grids := parse("data/q1_test.txt")
	have := q1(grids)
	want := 405

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestQ2_test(t *testing.T) {
	grids := parse("data/q1_test.txt")
	have := q2(grids)
	want := 400

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}
