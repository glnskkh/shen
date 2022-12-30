package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type comparable interface {
	int | float64
}

type TestCase[T comparable] struct {
	expected, input T
}

func TestCountSolutions(t *testing.T) {
	testCases := []TestCase[int]{
		{expected: 1, input: 1},
		{expected: 3, input: 2},
		{expected: 4, input: 4},
		{expected: 8, input: 6},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expected, CountSolutionsSlow(test.input))
	}
}
