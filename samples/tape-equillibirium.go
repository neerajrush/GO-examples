package main
import (
	"fmt"
	"math"
)

/************
A non-empty array A consisting of N integers is given. Array A represents numbers on a tape.

Any integer P, such that 0 < P < N, splits this tape into two non-empty parts: A[0], A[1], ..., A[P − 1] and A[P], A[P + 1], ..., A[N − 1].

The difference between the two parts is the value of: |(A[0] + A[1] + ... + A[P − 1]) − (A[P] + A[P + 1] + ... + A[N − 1])|

In other words, it is the absolute difference between the sum of the first part and the sum of the second part.

For example, consider array A such that:

  A[0] = 3
  A[1] = 1
  A[2] = 2
  A[3] = 4
  A[4] = 3
We can split this tape in four places:

P = 1, difference = |3 − 10| = 7 
P = 2, difference = |4 − 9| = 5 
P = 3, difference = |6 − 7| = 1 
P = 4, difference = |10 − 3| = 7 
Write a function:

int solution(int A[], int N);
that, given a non-empty array A of N integers, returns the minimal difference that can be achieved.

For example, given:

  A[0] = 3
  A[1] = 1
  A[2] = 2
  A[3] = 4
  A[4] = 3
the function should return 1, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [2..100,000];
each element of array A is an integer within the range [−1,000..1,000].
************/

func Solution(A []int) int {
	n := len(A)
	sum := 0
	for _, v := range A {
		sum += v
	}
	eql := math.MaxInt32
	sumL := 0
	sumR := sum
	for i := 0; i < n; i++ {
		if i > 0 {
			sumL += A[i-1]
			eql = int(math.Min(float64(eql), float64(math.Abs(float64(sumL) - float64(sumR)))))
		}
		sumR -= A[i]
	}

	return eql
}

func main() {
	A := []int{3, 1, 2, 4, 3}
	fmt.Println(A, Solution(A))

	B  := []int{3, 1, 2, 4, 3, 2, 1, 5}
	fmt.Println(B , Solution(B))

	C := []int{-1000, 1000}
	fmt.Println(C , Solution(C))

	D := []int{-10, -5, -3, -4, -5}
	fmt.Println(D , Solution(D))
}
