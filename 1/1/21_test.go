package shen

import "fmt"

func ExampleSquares() {
	PrintSquares(4)
	// Output: 0 1 4 9 16
}

func PrintSquares(n int) {
	s := 0

	// Inv. s = k^2
	for k := 0; k <= n; k++ {
		fmt.Printf("%d ", s)
		s += k + k + 1
	}
}
