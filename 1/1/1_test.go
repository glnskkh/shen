package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testCase(t *testing.T, valA, valB int) {
	a := valA
	b := valB

	Swap(&a, &b)

	assert.Equal(t, a, valB)
	assert.Equal(t, b, valA)
}

func TestSwap(t *testing.T) {
	testCase(t, 12, 22)
	testCase(t, 11, 11)
	testCase(t, 0, -1)
}

func Swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
