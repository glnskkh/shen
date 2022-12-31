package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCDModEuclid(t *testing.T) {
	assert.Equal(t, 12, GCDModEuclid(12, 36))
	assert.Equal(t, 12, GCDModEuclid(132, 36))
	assert.Equal(t, 1, GCDModEuclid(5, 3))
}

func GCDModEuclid(a, b int) int {
	m := a
	n := b

	// Inv. GCD(a; b) = GCD(m; n); m, n >= 0
	for !(m == 0 || n == 0) {
		if m >= n {
			m %= n
		} else {
			n %= m
		}
	}

	return m + n
}
