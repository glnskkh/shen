package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPower(t *testing.T) {
	assert.Equal(t, 4, SlowPower(2, 2))
	assert.Equal(t, 144, SlowPower(12, 2))
	assert.Equal(t, 1, SlowPower(4, 0))
}

func Power(a, n int) int {
	k := n
	b := 1
	c := a

	/* Inv. a^n = b * c ^ k */
	for k > 0 {
		if k%2 == 0 {
			c *= c
			k /= 2
		} else {
			b *= c
			k--
		}
	}

	return b
}
