package main

import (
	"testing"
)

func TestQ1_test(t *testing.T) {
	cards := parse("data/q1_test.txt")
	have := q1(cards)
	want := 13

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestQ2_test(t *testing.T) {
	cards := parse("data/q2_test.txt")
	have := q2(cards)
	want := 30

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}
