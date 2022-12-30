package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFastFib(t *testing.T) {
	assert.Equal(t, 0, FastFib(0))
	assert.Equal(t, 1, FastFib(1))
	assert.Equal(t, 1, FastFib(2))
	assert.Equal(t, 2, FastFib(3))
	assert.Equal(t, 13, FastFib(7))
}

func TestMultMat22(t *testing.T) {
	assert.Equal(t, [2][2]int{{2, 1}, {1, 1}}, MultMat22([2][2]int{{1, 1}, {1, 0}}, [2][2]int{{1, 1}, {1, 0}}))
}

func MultMat22(A [2][2]int, B [2][2]int) [2][2]int {
	C := [2][2]int{}

	C[0][0] = A[0][0]*B[0][0] + A[0][1]*B[1][0]
	C[0][1] = A[0][0]*B[0][1] + A[0][1]*B[1][1]
	C[1][0] = A[1][0]*B[0][0] + A[1][1]*B[1][0]
	C[1][1] = A[1][0]*B[0][1] + A[1][1]*B[1][1]

	return C
}

func PowerMat22(A [2][2]int, n int) [2][2]int {
	k := n

	B := [2][2]int{{1, 0}, {0, 1}}
	C := [2][2]int{{A[0][0], A[0][1]}, {A[1][0], A[1][1]}}

	for k > 0 {
		if k%2 == 0 {
			C = MultMat22(C, C)
			k /= 2
		} else {
			B = MultMat22(B, C)
			k--
		}
	}

	return B
}

func FastFib(n int) int {
	A := [2][2]int{{1, 1}, {1, 0}}

	// O(logn) power expension
	A = PowerMat22(A, n)

	r := A[0][1]

	return r
}
