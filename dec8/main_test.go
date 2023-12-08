package main

import (
	"testing"
)

func TestQ1_test(t *testing.T) {
	steps, desertMap := parse("data/q1_test.txt")
	have := q1(steps, desertMap)
	want := 6

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestQ2_test(t *testing.T) {
	steps, desertMap := parse("data/q2_test.txt")
	have := q2(steps, desertMap)
	want := 6

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestLCM_test(t *testing.T) {
	have := lcm([]int{10, 12})
	if have != 60 {
		t.Fatalf("Expected %d got %d", 60, have)
	}

	have = lcm([]int{10, 12, 60})
	if have != 60 {
		t.Fatalf("Expected %d got %d", 60, have)
	}

	have = lcm([]int{10, 12, 60, 120})
	if have != 120 {
		t.Fatalf("Expected %d got %d", 120, have)
	}

}
