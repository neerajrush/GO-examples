package main

import (
	"fmt"
	"math"
)

func find_equilibirium(A []int) int {
	totalSum := 0
	for _,v := range A {
		l := totalSum
		if v > 0 {
			if l > math.MaxInt32 - v {
				return -1  // totalSum + v will overflow
			}
		} else {
			if l < 0 && l < math.MaxInt32 - v {
				return -1 // totalSum + v will overflow if both are negativc.
			}
		}
		totalSum += v
	}
	// we have already covered overflow. So way it will reapper in the same sums.
	sumLeft := 0
	sumRight := totalSum
	for i, v := range A {
		sumRight -= v
		if sumLeft == sumRight {
			return i
		}
		sumLeft += v
	}
	return -1
}

func main() {
	A := []int{ -3, 2, 5, -3, 2 }
	fmt.Println(A, " eq: ", find_equilibirium(A))

	B := []int{-1, 3, -4, 5, 1, -6, 2, 1 }
	fmt.Println(B, " eq: ", find_equilibirium(B))
}
