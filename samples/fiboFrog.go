package main

import "fmt"
import "math"

/***************
The Fibonacci sequence is defined using the following recursive formula:

    F(0) = 0
    F(1) = 1
    F(M) = F(M - 1) + F(M - 2) if M >= 2
A small frog wants to get to the other side of a river. The frog is initially located at one bank of the river (position −1) and wants to get to the other bank (position N). The frog can jump over any distance F(K), where F(K) is the K-th Fibonacci number. Luckily, there are many leaves on the river, and the frog can jump between the leaves, but only in the direction of the bank at position N.

The leaves on the river are represented in an array A consisting of N integers. Consecutive elements of array A represent consecutive positions from 0 to N − 1 on the river. Array A contains only 0s and/or 1s:

0 represents a position without a leaf;
1 represents a position containing a leaf.
The goal is to count the minimum number of jumps in which the frog can get to the other side of the river (from position −1 to position N). The frog can jump between positions −1 and N (the banks of the river) and every position containing a leaf.

For example, consider array A such that:

    A[0] = 0
    A[1] = 0
    A[2] = 0
    A[3] = 1
    A[4] = 1
    A[5] = 0
    A[6] = 1
    A[7] = 0
    A[8] = 0
    A[9] = 0
    A[10] = 0
The frog can make three jumps of length F(5) = 5, F(3) = 2 and F(5) = 5.

Write a function:

int solution(int A[], int N);
that, given an array A consisting of N integers, returns the minimum number of jumps by which the frog can get to the other side of the river. If the frog cannot reach the other side of the river, the function should return −1.

For example, given:

    A[0] = 0
    A[1] = 0
    A[2] = 0
    A[3] = 1
    A[4] = 1
    A[5] = 0
    A[6] = 1
    A[7] = 0
    A[8] = 0
    A[9] = 0
    A[10] = 0
the function should return 3, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer that can have one of the following values: 0, 1.
**************************************************************************************/

var fibMap map[int]bool
var P map[int]int

func init() {
	fibMap = make(map[int]bool)
	A := make([]int, 30)
	A[0] = 0
	A[1] = 1
	fibMap[A[0]] = true
	fibMap[A[1]] = true
	for i := 2; i < 30; i++ {
		A[i] = A[i-1] + A[i-2]
		fibMap[A[i]] = true
	}
	P = make(map[int]int)
}

func isFiboJump(n int) bool {
	if n < 4 {
		return true
	}
	if _,ok := fibMap[n]; ok {
		return true
	}
	return false
}

func Solution(A []int) int {
	n := len(A)
	if n == 0 {
		return 1
	}
	if isFiboJump(n+1) {
		return 1
	}

	minJumps := math.MaxInt32

	for i := 0; i < n; i++ {
		if A[i] == 0 || !isFiboJump(i+1) {
			continue
		}
		t := Solution(A[i+1:])
		if t == -1 {
			continue
		}
		minJumps = 1 + int(math.Min(float64(t), float64(minJumps)))
	}
	if minJumps == math.MaxInt32 {
		return -1
	}
	return minJumps
}

func SolutionDynamic(A []int) int {
	n := len(A)
	if isFiboJump(n+1) {
		return 1
	}

	if _,ok := P[n]; ok {
		return P[n]
	}

	minJumps := math.MaxInt32

        for i := 0; i < n; i++ {
	    if A[i] == 0 || !isFiboJump(i+1) {
		  continue
	    }
	    if isFiboJump(n-i) {
		return 2
	    }
	    t := SolutionDynamic(A[i+1:])
            if t == -1 {
		continue
	    }
	    minJumps = 1  + int(math.Min(float64(t), float64(minJumps)))
	}

	if minJumps == math.MaxInt32 {
	        P[n] = -1
		return -1
	}
	P[n] = minJumps
	return minJumps
}

func main() {
	A := []int{0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0 }
	fmt.Println(A, len(A), Solution(A))
	fmt.Println(A, len(A), SolutionDynamic(A))

	A = []int{1, 1, 1}
	fmt.Println(A, len(A), Solution(A))
	fmt.Println(A, len(A), SolutionDynamic(A))

	A = []int{0, 0, 0}
	fmt.Println(A, len(A), Solution(A))
	fmt.Println(A, len(A), SolutionDynamic(A))

	A = []int{0, 0, 0, 0}
	fmt.Println(A, len(A), Solution(A))
	fmt.Println(A, len(A), SolutionDynamic(A))

	A = []int{0, 0, 1, 0, 0, 0, 1, 1, 1, 1}
	fmt.Println(A, len(A), Solution(A))
	fmt.Println(A, len(A), SolutionDynamic(A))

	A = []int{1, 1, 0, 0, 0}
	fmt.Println(A, len(A), Solution(A))
	fmt.Println(A, len(A), SolutionDynamic(A))

	A = []int{0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1}
	fmt.Println(A, len(A), Solution(A))
	fmt.Println(A, len(A), SolutionDynamic(A))
}
