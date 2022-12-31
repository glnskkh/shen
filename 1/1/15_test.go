package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCDLinearEuclid(t *testing.T) {
	var d, x, y int

	d, x, y = GCDLinearEuclid(12, 36)
	assert.Equal(t, 12, d)
	assert.Equal(t, d, 12*x+36*y)

	d, x, y = GCDLinearEuclid(132, 36)
	assert.Equal(t, 12, d)
	assert.Equal(t, d, 132*x+36*y)

	d, x, y = GCDLinearEuclid(5, 3)
	assert.Equal(t, 1, d)
	assert.Equal(t, d, 5*x+3*y)
}

func GCDLinearEuclid(a, b int) (d, x, y int) {
	m := a
	n := b

	p := 1
	q := 0

	s := 0
	t := 1

	// Inv. GCD(a; b) = GCD(m; n); m, n >= 0
	// Inv. m = pa + qb; n = sa + tb
	for !(m == 0 || n == 0) {
		if m >= n {
			m -= n
			p -= s
			q -= t
		} else {
			n -= m
			s -= p
			t -= q
		}
	}

	if m != 0 {
		return m, p, q
	} else {
		return n, s, t
	}
}
