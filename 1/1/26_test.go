package shen

import "fmt"

func ExampleDecimal() {
	Decimal(12)
	Decimal(341)
	Decimal(100000)

	// Output: 12
	// 341
	// 100000
}

func Decimal(n int) {
	base := 1

	for base*10 < n {
		base *= 10
	}

	// Inv. digits above base are printed
	for base > 0 {
		fmt.Print(n / base)
		n = n % base
		base /= 10
	}

	fmt.Println()
}
