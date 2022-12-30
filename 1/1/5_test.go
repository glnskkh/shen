package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMult(t *testing.T) {
	assert.Equal(t, 4, Mult(2, 2))
	assert.Equal(t, 6, Mult(3, 2))
	assert.Equal(t, 6, Mult(2, 3))
	assert.Equal(t, 0, Mult(0, 2))
}

func Mult(a, b int) int {
	r := 0

	// Inv. r = a * k
	for k := 0; k < b; k++ {
		r += a
	}

	return r
}
