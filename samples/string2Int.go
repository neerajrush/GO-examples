package main

import (
	 "fmt"
	 "math"
	 "errors"
)

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

func main() {
	x1 := "123"
	x2 := "-1"
	x3 := "1a23"
	x4 := "--123"
	x5 := "0"
	x6 := "1234567890"
	x7 := "001"
	v1, err1 := String2Int(x1)
	if err1 == nil {
		fmt.Println(x1, v1)
	}
	v2, err2 := String2Int(x2)
	if err2 == nil {
		fmt.Println(x2, v2)
	}
	_, err3 := String2Int(x3)
	if err3 != nil {
		fmt.Println(x3, err3)
	}
	_, err4 := String2Int(x4)
	if err4 != nil {
		fmt.Println(x4, err4)
	}
	v5, err5 := String2Int(x5)
	if err5 == nil {
		fmt.Println(x5, v5)
	}
	v6, err6 := String2Int(x6)
	if err6 == nil {
		fmt.Println(x6, v6)
	}
	v7, err7 := String2Int(x7)
	if err7 == nil {
		fmt.Println(x7, v7)
	}
}
