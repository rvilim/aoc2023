package main

import (
	"testing"
)

func TestQ1_test(t *testing.T) {
	races := parse1("data/q1_test.txt")
	have := q(races)
	want := 288

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}
func TestQ2_test(t *testing.T) {
	races := parse2("data/q2_test.txt")
	have := q(races)
	want := 71503

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}
