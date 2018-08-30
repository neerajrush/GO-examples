package main

import "fmt"

/********
Task description
A non-empty array A consisting of N integers is given. The consecutive elements of array A represent consecutive cars on a road.

Array A contains only 0s and/or 1s:

0 represents a car traveling east,
1 represents a car traveling west.
The goal is to count passing cars. We say that a pair of cars (P, Q), where 0 ≤ P < Q < N, is passing when P is traveling to the east and Q is traveling to the west.

For example, consider array A such that:

  A[0] = 0
  A[1] = 1
  A[2] = 0
  A[3] = 1
  A[4] = 1
We have five pairs of passing cars: (0, 1), (0, 3), (0, 4), (2, 3), (2, 4).

Write a function:

func Solution(A []int) int

that, given a non-empty array A of N integers, returns the number of pairs of passing cars.

The function should return −1 if the number of pairs of passing cars exceeds 1,000,000,000.

For example, given:

  A[0] = 0
  A[1] = 1
  A[2] = 0
  A[3] = 1
  A[4] = 1
the function should return 5, as explained above.

Assume that:

N is an integer within the range [1..100,000];
each element of array A is an integer that can have one of the following values: 0, 1.
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(1) (not counting the storage required for input arguments).
********/

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    if len(A) > 100000 {
   	return 0
    }
    total := 0
    for i,e := range A {
        if e == 1 {
            continue
        }
	for _, w := range A[i+1:] {
            if w == 0 {
                continue
            }
            total++
        }
    }
    if total > 1000000000 {
        return -1
    }
    return total
}

func SolutionX(A []int) int {
    // write your code in Go 1.4
    n := len(A)
    if n > 100000 {
   	return 0
    }
    total := 0
    east := 0
    west := 0
    for i := n-1; i >= 0; i-- {
        if A[i] == 1 {
            west++
        }
        if A[i] == 0 {
	    east++
	    total += east * west
	    east = 0
        }
    }
    if total > 1000000000 {
        return -1
    }
    return total
}

func main() {
	A := []int { 0, 1, 1, 0, 1, 0, 1, 1, 0, 1 }
	fmt.Println(A, Solution(A))

	fmt.Println(A, SolutionX(A))
        B := []int { 0, 1, 0, 1, 1 }
	fmt.Println(B, SolutionX(B))
}
