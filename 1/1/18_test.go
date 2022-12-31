package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCDTwoEuclid(t *testing.T) {
	assert.Equal(t, 12, GCDModEuclid(12, 36))
	assert.Equal(t, 12, GCDModEuclid(132, 36))
	assert.Equal(t, 1, GCDModEuclid(5, 3))
}

func GCDTwoEuclid(a, b int) int {
	m := a
	n := b

	d := 1

	// Inv. GCD(a;b) = d*GCD(m;n)
	for !(m == 0 || n == 0) {
		if m%2 == 0 && n%2 == 0 {
			d *= 2
			m /= 2
			n /= 2
		} else if m%2 == 0 && n%2 == 1 {
			m /= 2
		} else if m%2 == 1 && n%2 == 0 {
			n /= 2
		} else if m >= n {
			m -= n
		} else {
			n -= m
		}
	}

	if m != 0 {
		return d * m
	} else {
		return d * n
	}
}
