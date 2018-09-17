package main

import "fmt"

var D map[int]string

func findCombinationsRecursive(result *[]string, input, resultStr string, pos int) {
	if pos >= len(input) {
		*result = append(*result, resultStr)
		return
	}
	if inp,ok := D[int(input[pos] - '0')]; ok {
		for _,s := range inp {
			findCombinationsRecursive(result, input, resultStr+string(s), pos+1)
		}
	}
}

func findCombinations(A string) []string {
	if A == "0" {
		result := make([]string, 1)
		result[0] = "0"
		return result
	}
	if A == "1" {
		result := make([]string, 1)
		result[0] = "1"
		return result
	}
        ref := [][]string{{"A", "B", "C"}, {"D", "E", "F"}, {"G", "H", "I"}, {"J", "K", "L"}, {"M", "N", "O"}, {"P", "Q", "R", "S"}, {"T", "U", "V"}, {"W", "X", "Y", "Z"}}
	result := make([]string, 1)

	for _,ip := range A {
		k := int(ip - '0')
		idx := k - 2
		T := make([]string, 0)
	        for r,_ := range result {
			for _, e := range ref[idx] {
				T = append(T, result[r]+e)
			}
		}
		result = make([]string, len(T))
		copy(result, T)
	}

	return result
}


func main() {
	D = map[int]string{ 0:"0", 1:"1", 2:"ABC", 3:"DEF", 4:"GHI", 5:"JKL", 6:"MNO", 7:"PQRS", 8:"TUV", 9:"WXYZ" }
	input := "23"
	result := make([]string, 1)
	if len(input) == 0 {
		return
	}
	if len(input) == 1 {
	        for _,ip := range input {
			if v,ok := D[int(ip-'0')]; ok {
				result[0] = v
			}
		}
		fmt.Println(result)
		return
	}
	for _,ip := range input {
		K := int(ip - '0')
		if V,ok := D[K]; ok {
		        T := make([]string, 0)
			for _,r := range result {
				for  _,v := range V {
					T = append(T, r + string(v))
				}
			}
			result = make([]string, len(T))
			copy(result, T)
		}
	}
	fmt.Println(result)
        fmt.Println(findCombinations(input))
	r_result := make([]string, 0)
	findCombinationsRecursive(&r_result, input, " ", 0)
	fmt.Println(r_result)
}

//A B C
//D E F
