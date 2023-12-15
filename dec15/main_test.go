package main

import (
	"testing"
)

func TestQ1_test(t *testing.T) {
	l := parse("data/q1_test.txt")
	have := q1(l)
	want := 1320

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestQ2_test(t *testing.T) {
	l := parse("data/q1_test.txt")
	m := parseOps(l)
	have := q2(m)
	want := 145

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}
