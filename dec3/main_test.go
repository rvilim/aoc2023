package main

import (
	"testing"
)

func TestSampleInputQ1_test(t *testing.T) {
	schematic := parse("data/q1_test.txt")
	have := q1(schematic)
	want := 4361

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestSampleInputQ2_test(t *testing.T) {
	schematic := parse("data/q2_test.txt")
	have := q2(schematic)
	want := 467835

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestQ1(t *testing.T) {

	schematic := []string{"123*123"}
	have := q2(schematic)
	want := 123 * 123

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}
