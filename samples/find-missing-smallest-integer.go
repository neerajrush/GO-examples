package main
//package solution

import (
	"fmt"
	"sort"
)

/*******************
ask description

This is a demo task.

Write a function:

    func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Assume that:

        N is an integer within the range [1..100,000];
        each element of array A is an integer within the range [−1,000,000..1,000,000].

Complexity:

        expected worst-case time complexity is O(N);
        expected worst-case space complexity is O(N) (not counting the storage required for input arguments).

*/

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func qDivide(A []int, begin, end int) int {
	pivot := A[end]
	i := begin - 1
	for j := begin; j < end; j++ {
		if A[j] <= pivot {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[end] = A[end], A[i+1]
	return i + 1
}

func qSort(A []int, begin, end int) {
	if len(A) <= 1 {
		return
	}

	if begin >= end {
		return
	}

	pivot := qDivide(A, begin, end)
	qSort(A, begin, pivot-1)
	qSort(A, pivot+1, end)
}

func SolutionX(A []int) int {
	// write your code in Go 1.4
	qSort(A, 0, len(A)-1)
	result := 1
	for _, v := range A {
		if v < 0 {
			continue
		}
		if v == result {
			result++
		}
	}
	return result
}

func Solution(A []int) int {
	result := 1
	for _, v := range A {
		if v < 0 {
			continue
		}
		if v == result {
			result++
		}
	}
	return result
}

func SolutionY(A []int) int {
	l := len(A)
	lx := make([]int, l)
	for _,v := range A {
		if v < 1 || v > l {
			continue
		}
		lx[v-1]++
	}
	for i, v := range lx {
		if v == 0 {
			return i+1
		}
		if i == l-1 {
			return i+2
		}
	}
	return 1
}

func SolutionZ(nums []int) int {
    n := len(nums)
    lx := make([]int, n+1)
    for _,v := range nums {
        if v < 1 || v > n {
            continue
        }
        lx[v]++
    }
    for i, v := range lx {
	if i == 0 {  // ignore first element (0: index)
		continue
	}
        if v == 0 {
            return i
        }
        if i == n {
            return i+1
        }
    }
    return 1
}

func main() {
	A := []int { 1, 3, 6, 4, 1, 2 }
	B := []int { 1, 2, 3 }
	C := []int { -1, -3 }

	fmt.Println("Solution 1")
	fmt.Println(A, SolutionY(A))
	fmt.Println(B, SolutionY(B))
	fmt.Println(C, SolutionY(C))

	fmt.Println("Solution 2")
	fmt.Println(A, SolutionZ(A))
	fmt.Println(B, SolutionZ(B))
	fmt.Println(C, SolutionZ(C))

	return
	fmt.Println(SolutionX(A))

	sort.Slice(A, func (i, j int) bool {
		return A[i] < A[j]
	})

	fmt.Println(A)
}
