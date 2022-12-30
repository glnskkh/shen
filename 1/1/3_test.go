package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlowPower(t *testing.T) {
	assert.Equal(t, 4, SlowPower(2, 2))
	assert.Equal(t, 144, SlowPower(12, 2))
	assert.Equal(t, 1, SlowPower(4, 0))
}

func SlowPower(a, n int) int {
	p := 1

	for k := 0; k < n; k++ {
		p *= a
	}

	return p
}
