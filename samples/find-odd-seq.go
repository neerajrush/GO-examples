package main

import "fmt"

func getOddNumberSeq(l, r int) {
	s := l
	e := r + 1
        if l%2 == 0 && r%2 == 0 {
		s = l + 1
		e = r
	}
        if l%2 == 0 {
		s = l + 1
	}
        if r%2 == 0 {
		e = r
	}

	fmt.Print("(", l, r, ") ==> [")

	for i := s; i < e+1; i+=2 {
		fmt.Print(i, " ")
	}
	fmt.Println("]")
}

func main() {
	getOddNumberSeq(2, 10)
	getOddNumberSeq(2, 11)
	getOddNumberSeq(3, 10)
	getOddNumberSeq(3, 11)
}
