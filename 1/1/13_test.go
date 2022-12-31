package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCDEuclid(t *testing.T) {
	assert.Equal(t, 12, GCDEuclid(12, 36))
	assert.Equal(t, 12, GCDEuclid(132, 36))
	assert.Equal(t, 1, GCDEuclid(5, 3))
}

func GCDEuclid(a, b int) int {
	m := a
	n := b

	// Inv. GCD(a; b) = GCD(m; n); m, n >= 0
	for !(m == 0 || n == 0) {
		if m >= n {
			m -= n
		} else {
			n -= m
		}
	}

	return m + n
}
