package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFact(t *testing.T) {
	assert.Equal(t, 6, Fact(3))
	assert.Equal(t, 1, Fact(0))
	assert.Equal(t, 120, Fact(5))
}

func Fact(n int) int {
	p := 1

	for k := 2; k <= n; k++ {
		p *= k
	}

	return p
}
