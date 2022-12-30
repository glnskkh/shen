package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 4, Sum(2, 2))
	assert.Equal(t, 4, Sum(1, 3))
	assert.Equal(t, 0, Sum(0, 0))
	assert.Equal(t, 76, Sum(44, 32))
}

func Sum(a, b int) int {
	r := a

	// Inv. r = a + k
	for k := 0; k < b; k++ {
		r++
	}

	return r
}
