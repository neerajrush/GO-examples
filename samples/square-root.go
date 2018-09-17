package main

import "fmt"

func sqrt(x int) float64 {
	for z := 1; z < 10; z++ {
		z -=  (z * z  - x)/2 * z
	}
	return float64(v)
}

func main() {
	x := 9
	fmt.Println(sqrt(x))
}
