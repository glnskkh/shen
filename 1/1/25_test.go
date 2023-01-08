package shen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrimeGaussian(t *testing.T) {
	assert.False(t, IsPrimeGaussian(2, 0))
	assert.False(t, IsPrimeGaussian(5, 0))

	assert.True(t, IsPrimeGaussian(1, 1))
	assert.True(t, IsPrimeGaussian(1, -1))
	assert.True(t, IsPrimeGaussian(0, 1))
	assert.True(t, IsPrimeGaussian(1, 0))
	assert.True(t, IsPrimeGaussian(3, 0))
	assert.True(t, IsPrimeGaussian(7, 0))
}

func IsPrimeGaussian(a, b int) bool {
	m := a
	if b > m {
		m = b
	}

	for c := 1; c < m; c++ {
		for d := 1; d < m; d++ {
			s := c*c + d*d

			if a == c && b == d {
				break
			}

			if (a*c+b*d)%s == 0 && (b*c-a*d)%s == 0 {
				return false
			}
		}
	}

	return true
}

func ExampleFactorizePrimesGaussian() {
	FactorizePrimesGaussian(2, 0)

	FactorizePrimesGaussian(5, 0)

	FactorizePrimesGaussian(0, 1)
}

func FactorizePrimesGaussian(a, b int) {
	m := a
	if m < b {
		m = b
	}
	// so m is maximum of a and b

	for c := 1; c < m; c++ {
		for d := 1; d < m; d++ {
			if !IsPrimeGaussian(c, d) {
				continue
			}

			k := a*c + b*d
			l := b*c - a*d
			p := c*c + d*d
			if k%p == 0 && l%p == 0 {
				fmt.Printf("%d + %di\n", c, d)
				a = k / p
				b = l / p
			}
		}
	}

	fmt.Printf("%d + %di\n", a, b)
}
