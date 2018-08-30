package main

import "fmt"

/************
An array A consisting of N different integers is given. The array contains integers in the range [1..(N + 1)], which means that exactly one element is missing.

Your goal is to find that missing element.

Write a function:

int solution(int A[], int N);
that, given an array A, returns the value of the missing element.

For example, given array A such that:

  A[0] = 2
  A[1] = 3
  A[2] = 1
  A[3] = 5
the function should return 4, as it is the missing element.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
the elements of A are all distinct;
each element of array A is an integer within the range [1..(N + 1)].
************/

func Solution(A []int ) int {
	n := len(A)
	lx := make([]int, n+2)
	for _, v := range A {
		if v > n+1 {
			return 0
		}
		lx[v] = 1
	}
	for i := 1; i < n+2; i++ {
		if lx[i] == 0 {
			return i
		}
	}
	return 0
}

func main() {
	A := []int{ 2, 3, 1, 5 }
	fmt.Println(A, Solution(A))

	B := []int{ 7, 9, 5, 8, 1, 3, 4, 6 }
	fmt.Println(B, Solution(B))
}
