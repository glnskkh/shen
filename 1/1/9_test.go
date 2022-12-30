package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	assert.Equal(t, 0, Fib(0))
	assert.Equal(t, 1, Fib(1))
	assert.Equal(t, 1, Fib(2))
	assert.Equal(t, 2, Fib(3))
	assert.Equal(t, 13, Fib(7))
}

func Fib(n int) int {
	a := 0
	b := 1

	// Inv. a = f_k, b = f_(k+1)
	for k := 0; k < n; k++ {
		b = a + b
		a = b - a
	}

	return a
}
