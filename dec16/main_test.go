package main

import (
	"testing"
)

func TestQ1_test(t *testing.T) {
	l := parse("data/q1_test.txt")
	have := q1(l)
	want := 46

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestQ2_test(t *testing.T) {
	l := parse("data/q1_test.txt")
	have := q2(l)
	want := 51

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}
