package main

import (
	"testing"
)

func TestQ1_test(t *testing.T) {
	pipeMap, start := parse("data/q1_test.txt")
	have := q1(pipeMap, start)
	want := 8

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestQ2_test(t *testing.T) {
	// Define a struct for test cases
	type testCase struct {
		filename string
		want     int
	}

	// Create a slice of test cases
	testCases := []testCase{
		{"data/q2_test_1.txt", 4},
		{"data/q2_test_2.txt", 8},
		{"data/q2_test_3.txt", 10},
		// Add more test cases here
	}

	// Loop over the test cases
	for _, tc := range testCases {
		pipeMap, start := parse(tc.filename)
		have := q2(pipeMap, start)

		if tc.want != have {
			t.Fatalf("Test failed for file %s: Expected %d got %d", tc.filename, tc.want, have)
		}
	}
}
