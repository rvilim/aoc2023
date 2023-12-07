package main

import (
	"testing"
)

func TestQ1_test(t *testing.T) {
	hands := parse("data/q1_test.txt")
	have := q1(hands)
	want := 6440

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestQ2_test(t *testing.T) {
	hands := parse("data/q2_test.txt")
	have := q2(hands)
	want := 5905

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}
