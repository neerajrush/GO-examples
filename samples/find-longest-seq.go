package main

import (
	"fmt"
	"math"
)

// Longest sequence max

func SolutionLSM(A []int) int {
	n := len(A)
 	L := make([]int, n)
	for i,_ := range L {
		L[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if A[j] < A[i] {
				L[i] = int(math.Max(float64(L[i]), float64(L[j]+1)))
			}
		}
	}
	fmt.Println(L)
	max := 0
	for _,v := range L {
		if v > max {
			max = v
		}
	}
	return max
}

// sequence Max Sum

func SolutionLSS(A []int) int {
	n := len(A)
 	L := make([]int, n)
	copy(L, A)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if A[j] < A[i] {
				L[i] = int(math.Max(float64(L[i]), float64(L[j]+A[i])))
			}
		}
	}
	fmt.Println(L)
	max := 0
	for _,v := range L {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	A := []int{ 3, 10, 4, 20, 1, 2} 

	fmt.Println(A, SolutionLSM(A))
	fmt.Println(A, SolutionLSS(A))
}
