package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testSwapCase(t *testing.T, valA, valB int, f func(*int, *int)) {
	a := valA
	b := valB

	f(&a, &b)

	assert.Equal(t, a, valB)
	assert.Equal(t, b, valA)
}

func TestSwap(t *testing.T) {
	testSwapCase(t, 12, 22, Swap)
	testSwapCase(t, 11, 11, Swap)
	testSwapCase(t, 0, -1, Swap)
}

func Swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
