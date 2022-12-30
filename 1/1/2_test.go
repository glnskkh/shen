package shen

import "testing"

func TestArithmeticSwap(t *testing.T) {
	testSwapCase(t, 12, 22, ArithmeticSwap)
	testSwapCase(t, 11, 11, ArithmeticSwap)
	testSwapCase(t, 0, -1, ArithmeticSwap)
}

func ArithmeticSwap(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}
