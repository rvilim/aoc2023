package main

import (
	"testing"
)

func TestQ1_test(t *testing.T) {
	seeds, mappings := parse("data/q1_test.txt")
	have := q1(seeds, mappings)
	want := int64(35)

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestQ2_test(t *testing.T) {
	seeds, mappings := parse("data/q2_test.txt")
	have := q2(seeds, mappings)
	want := int64(46)

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}
