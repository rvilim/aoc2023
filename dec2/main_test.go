package main

import (
	"testing"
)

func TestSampleInputQ1_test(t *testing.T) {
	have := q1("data/q1_test.txt")
	want := 8

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestSampleInputQ2_test(t *testing.T) {
	have := q2("data/q2_test.txt")
	want := 2286

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}
