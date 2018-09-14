package main

// you can also use imports, for example:
import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    countMap := make(map[int]int)
    for _,v := range A {
	countMap[v] = 1
    }
    return len(countMap)
}

func main() {
	A := []int{-1, -2, -3, -1, -2}
	fmt.Println(A, Solution(A))
}
