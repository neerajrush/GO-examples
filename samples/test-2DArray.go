package main

import "fmt"

func main() {
	A := make([][]int, 10)
	for i,_ := range A {
		A[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			A[i][j] = j+1
		}
	}

	for i,_ := range A {
		fmt.Print(i)
		for j := 0; j < 5; j++ {
			fmt.Print("==>", A[i][j])
		}
		fmt.Println("")
	}
}
