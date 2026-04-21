package main

import "testing"

func TestSquare(t *testing.T) {
	for i := 1; i < 100000; i++ {
		result := square(i)
		if result/i != i {
			t.Errorf("expected: %d, result: %d\n", i*i, result)
		}
	}
}

func TestDayOfWeek(t *testing.T) {
	type test struct {
		Input          int
		ExpectedResult string
	}

	var testCases = []test{
		{Input: 1, ExpectedResult: "شنبه"},
		{Input: 2, ExpectedResult: "یکشنبه"},
		{Input: 3, ExpectedResult: "دوشنبه"},
		{Input: 4, ExpectedResult: "سه شنبه"},
		{Input: 5, ExpectedResult: "چهارشنبه"},
		{Input: 6, ExpectedResult: "پنجشنبه"},
		{Input: 7, ExpectedResult: "جمعه"},
		{Input: 8, ExpectedResult: ""},
		{Input: 0, ExpectedResult: ""},
		{Input: -1, ExpectedResult: ""},
	}

	for _, c := range testCases {
		result := dayOfWeek(c.Input)
		if result != c.ExpectedResult {
			t.Errorf("expected: %s, result: %s\n", c.ExpectedResult, result)
		}
	}
}
