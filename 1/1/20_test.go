package shen

import "fmt"

func ExampleDullSquares() {
	PrintDullSquares(4)
	// Output: 0 1 4 9 16
}

func PrintDullSquares(n int) {
	// Inv. 0, 1, 4, ..., k^2 are printed
	for k := 0; k <= n; k++ {
		fmt.Printf("%d ", k*k)
	}
}
