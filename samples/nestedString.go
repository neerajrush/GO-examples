package main

import "fmt"

/****************
A string S consisting of N characters is called properly nested if:

S is empty;
S has the form "(U)" where U is a properly nested string;
S has the form "VW" where V and W are properly nested strings.
For example, string "(()(())())" is properly nested but string "())" isn't.

Write a function:

int solution(char *S);
that, given a string S consisting of N characters, returns 1 if string S is properly nested and 0 otherwise.

For example, given S = "(()(())())", the function should return 1 and given S = "())", the function should return 0, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..1,000,000];
string S consists only of the characters "(" and/or ")".
****************/

type Stack struct {
	S []byte
	P int
        o bool
        u bool
}

func (s *Stack) push(c byte) bool {
	if s.o  {
		return false
	}
	s.S[s.P] = c
	s.P += 1
	if s.P >= 100000 {
		s.o = true
	}
	if s.P >= 0 {
		s.u = false
	}
	if s.P < 100000 {
		s.o = false
	}
	return true
}

func (s *Stack) pop() (byte, bool) {
	if s.u  {
		return '0', true
	}
	s.P -= 1
	if s.P < 0 {
		s.u = true
		return '0', true
	}
	if s.P >= 0 {
		s.u = false
	}
	if s.P < 100000 {
		s.o = false
	}
	c := s.S[s.P]
	return c, false
}

func (s *Stack) empty() bool {
	if s.P == 0 {
		return true
	}
	return false
}

func (s *Stack) printStack() {
	for i := 0; i < s.P; i++ {
		fmt.Printf("%c %d\n", s.S[i], s.P)
	}
}

var St Stack

func Solution(s string) int {
	if s == "" {
		return 1
	}
	for _,v := range s {
		if v == '(' {
			if !St.push(byte(v)) {
				return 0
			}
		}
		if v == ')' {
			c, err := St.pop()
			if err || c != '(' {
				return 0
			}
		}
	}

	if St.empty() {
		return 1
	}

	return 0
}

func main() {
	St = Stack{S: make([]byte, 1000000), P: 0, o: false, u: false}
	s := "(()()(()))"
	fmt.Println(s, Solution(s))

	s = "(()))"
	fmt.Println(s, Solution(s))
}
