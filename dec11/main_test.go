package main

import (
	"testing"
)

func TestQ1_test(t *testing.T) {
	universe := parse("data/q1_test.txt")
	have := q1(universe)
	want := 374

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestCountpath_test(t *testing.T) {
	type testCase struct {
		expansionFactor int
		want            int
	}

	testCases := []testCase{
		{10, 1030},
		{100, 8410},
	}

	universe := parse("data/q2_test.txt")
	for _, tc := range testCases {
		have := countPath(universe, tc.expansionFactor)
		if tc.want != have {
			t.Fatalf("Expected %d got %d", tc.want, have)
		}
	}
}
