package shen

import "fmt"

func ExampleFasterFactorizePrimes() {
	FactorizePrimes(12)
	// Output: 2 2 3
}

func FasterFactorizePrimes(n int) {
	l := 2

	for n != 1 {
		if n%l == 0 {
			// l|n и нет других делителей, меньших l => l просто
			fmt.Println(l)
			n /= l
		} else {
			if l*l > n {
				l = n // IDK WHY
			} else {
				l++
			}
		}
	}
}
