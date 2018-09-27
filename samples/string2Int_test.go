package main

import (
	"testing"
	"math"
	"errors"
	"fmt"
)

type TestTable struct {
	input string
	output int
	expected int
	err	error
}

var Tests []TestTable

func TestString2Int(t *testing.T) {
	Tests = []TestTable{
			TestTable{input: "123",          output: 0,   expected: 123,         err: nil},
			TestTable{input: "-1",           output: 0,   expected: -1,          err: nil},
			TestTable{input: "1a23",         output: 0,   expected: 0,           err: nil},
			TestTable{input: "--123",        output: 0,   expected: 0,           err: nil},
			TestTable{input: "0",            output: 0,   expected: 0,           err: nil},
			TestTable{input: "1234567890",   output: 0,   expected: 1234567890,  err: nil},
			TestTable{input: "001",          output: 0,   expected: 1,           err: nil},
       }
	for i,x := range Tests {
		x.output, x.err = String2Int(x.input)
		if x.output != x.expected {
			t.Errorf("String2Int: testcase %d input: %s output: %d expected: %d", i, x.input, x.output, x.expected)
		}
	}
}

func String2Int(str string) (int, error) {
	n := len(str)
	if n == 0 {
		return 0, errors.New("empty string")
	}
	start := 0
	multiplier := 1
	result := 0
	switch byte(str[0]) {
	case '-':
		multiplier = -1
		start = 1
	}
	for i := start; i < n; i++ {
		diff := byte(str[i]) - '0'
		if diff < 0 || diff > 9 {
			msg := fmt.Sprintf("non-digit char: %s", string(str[i]))
			return 0, errors.New(msg)
		}
		result += int(diff) * int(math.Pow(float64(10), float64(n-i-1)))
	}

	return multiplier * result, nil
}
