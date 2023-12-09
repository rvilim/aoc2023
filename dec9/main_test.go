package main

import (
	"testing"
)

func TestQ1_test(t *testing.T) {
	seqs := parse("data/q1_test.txt")
	have := q1(seqs)
	want := 114

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestQ2_test(t *testing.T) {
	seqs := parse("data/q2_test.txt")
	have := q2(seqs)
	want := 2

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}
