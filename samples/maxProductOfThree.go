package main

import "fmt"

func find_min_max(A []int) (int, int) {
    MAX := 1001
    MIN := -1001
    max := MIN
    min := MAX
    m   := 0
    n   := 0
    for i, v := range A {
	if v == MAX || v == MIN {
		continue
	}
        if v > max {
            max = v
	    m = i
        }
        if v < min {
            min = v
	    n = i
        }
    }
    A[m] = MAX
    A[n] = MIN
    return  max, min
}

func max_product(A []int) int {
    N := len(A)
    maxProd := -1001
	for i,_ := range A {
		if i + 2 > N {
			break
		}
		for j := i+1; i < N; j++ {
			if j + 1 > N {
				break
			}
			for k := j+1; k < N; k++ {
				prod := A[i] * A[j] * A[k]
				if prod > maxProd {
					maxProd = prod
				}
			}
		}
	}
	return maxProd
}

func Solution(A []int) int {
    N := len(A)
    if N < 3 {
        return 0
    }

    if N > 5 {
       m1, n1 := find_min_max(A)
       m2, n2 := find_min_max(A)
       m3, n3 := find_min_max(A)
       if m1 < 0 || m2 < 0 || m3 < 0 ||
          n1 < 0 || n2 < 0 || n2 < 0 {
           B := []int{m1, m2, m3, n1, n2, n3}
           return max_product(B)
       }
       m := m1 * m2 * m3
       n := n1 * n2 * n3

       if m > n {
           return m
       }
       return n
    } else {
        return max_product(A)
    }
}

func main() {
	A := []int{-3, 1, 2, -2, 5, 6}
	fmt.Println(A, Solution(A))

	B := []int{4, 5, 1, 0}
	fmt.Println(B, Solution(B))

	C := []int{4, 7, 3, 2, 1, -3, -5}
	fmt.Println(C, Solution(C))

	D := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		if i < 50000 {
			D[i] = -1000
		} else {
			D[i] = 1
		}
	}
	fmt.Println("[-2,-2,...,1,1]", Solution(D))
}
