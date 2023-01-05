package shen

import (
	"fmt"
)

// func TestIsPrimeGaussian(t *testing.T) {
// 	assert.False(t, IsPrimeGaussian(2, 0))
// 	assert.False(t, IsPrimeGaussian(5, 0))

// 	assert.True(t, IsPrimeGaussian(0, 1))
// 	assert.True(t, IsPrimeGaussian(1, 0))
// }

func IsPrimeGaussian(a, b int) bool {
	m := a
	if b > m {
		m = b
	}

	for c := 0; c <= m; c++ {
		for d := 1; d <= c; d++ {
			s := c*c + d*d

			if (a*c+b*d)%s == 0 && (b*c-a*d)%s == 0 {
				return false
			}
		}
	}

	return true
}

// func ExampleFactorizePrimesGaussian() {
// 	FactorizePrimesGaussian(2, 0)

// 	FactorizePrimesGaussian(5, 0)

// 	FactorizePrimesGaussian(0, 1)

// 	// Unordered output: 1 + 1i
// 	// 1 - 1i
// 	// 2 - 1i
// 	// 2 + 1i
// 	// 0 + 1i
// }

func FactorizePrimesGaussian(a, b int) {
	m := a
	if b > m {
		m = b
	}

	for c := 0; c <= m && a*a+b*b > 0; c++ {
		for d := -m; d <= m && a*a+b*b > 0; d++ {
			s := c*c + d*d

			if !(s != 0 && IsPrimeGaussian(c, d)) {
				continue
			}

			if (a*c+b*d)%s == 0 && (b*c-a*d)%s == 0 {
				fmt.Printf("%d + %di\n", c, d)

				a = (a*c + b*d) / s
				b = (b*c - a*d) / s
			}
		}
	}
}
