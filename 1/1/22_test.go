package shen

import "fmt"

func ExampleFactorizePrimes() {
	FactorizePrimes(12)
	FactorizePrimes(1000)

	// Output: 2 2 3 2 2 2 5 5 5
}

func FactorizePrimes(n int) {
	k := 2

	// Inv. all primes up to k are printed
	for n > 1 {
		for n%k != 0 {
			k += 1
		}

		// k|n, k - наименьшее из тех, что делит n => k простое
		fmt.Printf("%d ", k)
		n /= k
	}
}
