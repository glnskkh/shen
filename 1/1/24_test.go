package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	assert.True(t, IsPrime(11))
	assert.True(t, IsPrime(7))
	assert.True(t, IsPrime(1021))

	assert.False(t, IsPrime(12))
	assert.False(t, IsPrime(1024))
}

func IsPrime(a int) bool {
	for k := 2; k*k <= a; k++ {
		if a%k == 0 {
			return false
		}
	}

	return true
}
