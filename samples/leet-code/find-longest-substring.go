package main

import "fmt"
import "math"

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func Solution(A string, pos int) int {
	if pos >= len(A){
		return 0
	}
        M := make(map[byte]bool)
        xLen := 0
	for i := pos; i < len(A); i++ {
	    if _,ok := M[A[i]]; ok {
		break
	    } else {
	        M[A[i]] = true
		xLen++
	   }
	}
        return int(math.Max(float64(Solution(A, pos+1)), float64(xLen)))
}

func longestSubstr(A string) int {
	N := len(A)
	maxLen := 0
	for i := 0; i < N; i++ {
	        M := make(map[byte]bool)
		xLen := 0
		for j := i; j < N; j++ {
			if _,ok := M[A[j]]; ok {
				break
			} else {
				M[A[j]] = true
				xLen++
			}
		}
		maxLen = int(math.Max(float64(maxLen), float64(xLen)))
	}
	return maxLen
}

func SolutionC(A string) int {
	N := len(A)
	if N == 0 {
		return 0
	}
	M := make(map[byte]int)
	maxLen := 0
	j := 0
	for i := 0; i < N; i++ {
		if _,ok := M[A[i]]; ok {
			j = int(math.Max(float64(j), float64(M[A[i]]+1)))
		}
		M[A[i]] = i
		maxLen = int(math.Max(float64(maxLen), float64(i-j+1)))
	}
	return maxLen
}

func main() {
	A := "abcabcbb"
	fmt.Println(A, longestSubstr(A))
	fmt.Println(A, Solution(A, 0))
	fmt.Println(A, SolutionC(A))

	A = "bbbbb"
	fmt.Println(A, longestSubstr(A))
	fmt.Println(A, Solution(A, 0))
	fmt.Println(A, SolutionC(A))

	A = "pwwkew"
	fmt.Println(A, longestSubstr(A))
	fmt.Println(A, Solution(A, 0))
	fmt.Println(A, SolutionC(A))

	A = "abba"
	fmt.Println(A, longestSubstr(A))
	fmt.Println(A, Solution(A, 0))
	fmt.Println(A, SolutionC(A))
}
