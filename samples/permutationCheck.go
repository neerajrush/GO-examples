package main

import (
	"fmt"
)

/**************
A non-empty array A consisting of N integers is given.

A permutation is a sequence containing each element from 1 to N once, and only once.

For example, array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2
is a permutation, but array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
is not a permutation, because value 2 is missing.

The goal is to check whether array A is a permutation.

Write a function:

int solution(int A[], int N);
that, given an array A, returns 1 if array A is a permutation and 0 if it is not.

For example, given array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2
the function should return 1.

Given array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
the function should return 0.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [1..1,000,000,000].
**************/

func Solution(A []int) int {
	n := len(A)
	lx := make([]int, n+2)
	for _, v := range A {
		if v > n+1 {
			return 0
		}
		lx[v] = 1
	}
	for i := 1; i < n+1; i++ {
		if lx[i] == 0 {
			return 0
		}
	}
	return 1
}

func main() {
	A := []int{4, 1, 3, 2}
	fmt.Println(A, Solution(A))

	B := []int{4, 1, 3}
	fmt.Println(B, Solution(B))

	C := []int{3, 1}
	fmt.Println(C, Solution(C))

	D := []int{2}
	fmt.Println(D, Solution(D))

	E := []int{1}
	fmt.Println(E, Solution(E))
}
