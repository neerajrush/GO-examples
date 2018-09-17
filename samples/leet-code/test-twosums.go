package main

import "fmt"

// Brute force O(n^2)
func Solution1(A []int, sum int) []int {
	N := len(A)
	result := make([]int, 0)
	for i := 0; i < N; i++ {
		for j := i+1; j < N; j++ {
			if A[i] + A[j] == sum {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return result
}

// Better solutuon O(2n)

func Solution2(A []int, sum int) []int {
	N := len(A)
	result := make([]int, 0)
	M := make(map[int]int)
	for i,v := range A {
		M[v] = i
	}
	for i := 0; i < N; i++ {
		diff := sum - A[i]
		if v,ok := M[diff]; ok && v != i {
			result = append(result, i)
			result = append(result, v)
			return result
		}
	}

	return result
}

// More Better solutuon O(n)

func Solution3(A []int, sum int) []int {
	N := len(A)
	result := make([]int, 0)
	M := make(map[int]int)
	for i := 0; i < N; i++ {
		diff := sum - A[i]
		if v,ok := M[diff]; ok && v != i {
			result = append(result, v)
			result = append(result, i)
			return result
		} else {
			M[A[i]] = i
		}
	}

	return result
}

func twoSum(nums []int, target int) []int {
    N := len(nums)
    result := make([]int, 0)
    M := make(map[int]int)
    for i := 0; i < N; i++ {
        diff := target - nums[i]
        if v, ok := M[diff]; ok && v != i {
            result = append(result, v)
            result = append(result, i)
            return result
        } else {
            M[nums[i]] = i
        }
    }
    return result
}

func main() {
	A := []int{4,7,9,2,6} // {2, 7, 11, 15, 17, 21, 23, 27}
	sum := 8 //40
	fmt.Println(sum, Solution1(A, sum))
	fmt.Println(sum, Solution2(A, sum))
	fmt.Println(sum, Solution3(A, sum))

}
