package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivMod(t *testing.T) {
	var div, rem int

	div, rem = DivMod(12, 5)
	assert.Equal(t, div, 2)
	assert.Equal(t, rem, 2)

	div, rem = DivMod(33, 2)
	assert.Equal(t, div, 16)
	assert.Equal(t, rem, 1)
}

func DivMod(a, b int) (div, rem int) {
	rem = a
	div = 0

	// Inv. a = div * b + rem
	for !(rem < b) {
		rem -= b
		div++
	}

	return div, rem
}
