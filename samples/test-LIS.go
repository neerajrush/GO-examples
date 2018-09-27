package main

import "fmt"
import "math"

func lis(A string) int {
	L := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		L[i] = 1
	}
	max := 0
	for i := 1; i < len(A); i++ {
		for j := 0; j <= i; j++ {
			if A[j] < A[i] {
				L[i] = int(math.Max(float64(L[i]), float64(L[j]+1)))
			}
		}
	}
	fmt.Println(L)
	for _,v := range L {
		if max < v {
			max = v
		}
	}
        return max
}

func main() {
	A := "ABCABCBADEF"
	fmt.Println(A, lis(A))
}
